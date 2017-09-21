package main

// Copyright 2017 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/Masterminds/semver"
)

const (
	defaultRemoteAzureRestAPISpecsPath = "https://github.com/Azure/azure-rest-api-specs.git"
	defaultAzureRESTAPIBranch          = "current"
)

// ExitCode gives a hint to the end user about why the program exited without relying on seeing stderr.
const (
	ExitCodeOkay int = iota
	ExitCodeMissingRequirements
	ExitCodeFailedToClone
	ExitCodeFinderFailure
)

var (
	remoteAzureRESTAPISpecsPath string
	localAzureRESTAPISpecsPath  string
	azureRESTAPIBranch          string
	outputLocation              string
	dryRun                      bool
	help                        bool
	anyMissing                  bool
	noClone                     bool
	verbose                     bool
	wait                        bool
	targetFile                  string
	debugLog                    *log.Logger
	version                     *semver.Version
)

func init() {
	var err error
	flag.StringVar(&azureRESTAPIBranch, "branch", defaultAzureRESTAPIBranch, "The branch, tag, or SHA1 identifier in the Azure Rest API Specs repository to use during API generation.")
	flag.StringVar(&remoteAzureRESTAPISpecsPath, "repo", defaultRemoteAzureRestAPISpecsPath, "The path to the location of the Azure REST API Specs repository that should be used for generation.")
	flag.StringVar(&outputLocation, "output", getDefaultOutputLocation(), "a directory in which all generated code should be placed.")
	flag.StringVar(&targetFile, "target", "", "If a target file is provided, generator will only run on this file instead of all swaggers it encounters in the repository.")
	flag.BoolVar(&dryRun, "preview", false, "Use this flag to print a list of packages that would be generated instead of actually generating the new sdk.")
	flag.BoolVar(&help, "help", false, "Provide this flag to print usage information instead of running the program.")
	flag.BoolVar(&noClone, "noClone", false, "Use this flag to prevent cloning a new copy of Azure REST API Specs. The existing enlistment should be used instead.")
	flag.BoolVar(&verbose, "verbose", false, "Print status messages as processing is done.")
	flag.BoolVar(&wait, "wait", false, "Use this program to halt execution before the cleanup phase is entered.")
	useDebug := flag.Bool("debug", false, "Include this flag to print debug messages as the program executes.")
	rawVersion := flag.String("version", "0.0.0", "The version that should be stamped on the generated code. Common usage will be to report user agent.")

	flag.Parse()

	if help {
		return
	}

	version, err = semver.NewVersion(*rawVersion)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read version \"%s\" because: %v", *rawVersion, err)
		os.Exit(1)
	}

	debugWriter := ioutil.Discard
	if *useDebug {
		fmt.Fprintln(os.Stdout, "Turning on debug logger.")
		debugWriter = os.Stderr
	}
	debugLog = log.New(debugWriter, "[DEBUG]", log.Ltime)

	optionalTools := []string{"gofmt", "golint"}
	requiredTools := []string{"autorest", "git"}

	for _, tool := range optionalTools {
		if _, err := exec.LookPath(tool); err != nil {
			log.Printf("WARNING: Could not find \"%s\" usage of this tool will be skipped.", tool)
		}
	}

	anyMissing = false
	for _, tool := range requiredTools {
		if _, err := exec.LookPath(tool); err != nil {
			log.Printf("ERROR: Could not find \"%s\" this tool will exit without executing.", tool)
			anyMissing = true
		}
	}

	if noClone {
		localAzureRESTAPISpecsPath = remoteAzureRESTAPISpecsPath
	} else {
		var err error
		localAzureRESTAPISpecsPath, err = ioutil.TempDir("", "")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(ExitCodeFailedToClone)
		}
	}
}

func main() {
	var repoLoc string
	exitStatus := ExitCodeOkay
	var err error
	defer func() { os.Exit(exitStatus) }()

	errLog := log.New(os.Stderr, "[ERROR] ", 0)

	if help {
		flag.Usage()
		return
	}

	if anyMissing {
		exitStatus = ExitCodeMissingRequirements
		return
	}

	if noClone {
		repoLoc = localAzureRESTAPISpecsPath
	} else if temp, err := cloneAPISpecs(remoteAzureRESTAPISpecsPath, localAzureRESTAPISpecsPath); err == nil {
		repoLoc = temp
		defer func() {
			if wait {
				fmt.Print("Press ENTER to continue...")
				fmt.Scanln()
			}

			if err := os.RemoveAll(repoLoc); err != nil {
				errLog.Print(err)
			}
		}()
	} else {
		errLog.Print(err)
		exitStatus = ExitCodeFailedToClone
		return
	}

	if err = checkoutAPISpecsVer(azureRESTAPIBranch, repoLoc); err != nil {
		errLog.Print(err)
		exitStatus = ExitCodeFailedToClone
		return
	}

	finderOutput := ioutil.Discard
	if verbose {
		finderOutput = os.Stdout
	}

	finder, err := NewSwaggerFinder(repoLoc, finderOutput)
	if err != nil {
		errLog.Print(err)
		exitStatus = ExitCodeFinderFailure
		return
	}

	if dryRun {
		fmt.Println("Executing Preview")
		namespaces := finder.Enumerate(nil).Select(func(x interface{}) interface{} {
			namespace, err := getNamespace(x.(SwaggerFile))
			if err != nil {
				return err
			}
			return namespace
		})

		for namespace := range namespaces {
			var output io.Writer
			switch namespace.(type) {
			case error:
				output = os.Stderr
			case string:
				output = os.Stdout
			}
			fmt.Fprintln(output, namespace)
		}
		return
	}

	var logFolder string
	useLogs := true

	logFolder, err = ioutil.TempDir("", "autorestLogs")
	if err != nil {
		useLogs = false
	}
	fmt.Println("Logs will be stored in: ", logFolder)

	found, generated, formatted, vetted := 0, 0, 0, 0

	processor := finder.Enumerate(nil).Select(func(x interface{}) interface{} {
		found++
		return x
	}).ParallelSelect(func(x interface{}) interface{} {
		cast := x.(SwaggerFile)
		var err error
		var logWriter io.Writer
		var outputFile *os.File
		if useLogs {
			outputFile, err = os.OpenFile(path.Join(logFolder, cast.Info.Title+"_"+cast.Info.Version+".txt"), os.O_WRONLY|os.O_CREATE, 0777)
			logWriter = outputFile
			if err == nil {
				defer outputFile.Close()
			} else {
				fmt.Fprintln(os.Stderr, "Cannot log results of: ", cast.Info.Title, cast.Info.Version, "because: ", err)
			}
		} else {
			logWriter = ioutil.Discard
		}
		var name string
		name, err = generate(cast, outputLocation, repoLoc, logWriter)
		if err != nil {
			return err
		}
		generated++
		return name
	}).Where(func(x interface{}) bool {
		switch x.(type) {
		case string:
			if verbose {
				fmt.Fprintln(os.Stdout, "Generated: ", x)
			}
			return true
		case error:
			fmt.Fprintln(os.Stderr, "Error: ", x)
		}
		return false
	}).Select(func(x interface{}) interface{} {
		subject := x.(string)
		subject = strings.TrimPrefix(subject, outputLocation)
		subject = path.Join("github.com", "Azure", "azure-sdk-for-go", subject)
		subject = path.Clean(subject)
		return subject
	}).Where(func(x interface{}) bool {
		return format(x.(string)) == nil
	}).Where(func(x interface{}) bool {
		formatted++
		return vet(x.(string)) == nil
	})

	for range processor {
		vetted++
	}

	fmt.Println("Summary:")
	fmt.Println("\tFound:    \t", found)
	fmt.Println("\tGenerated:\t", generated)
	fmt.Println("\tFormatted:\t", formatted)
	fmt.Println("\tVetted:   \t", vetted)
}

func cloneAPISpecs(origin, local string) (string, error) {
	_, cloneLoc := filepath.Split(local)
	if verbose {
		fmt.Println("Cloning Azure REST API Specs to: ", cloneLoc)
	}
	clone := exec.Command("git", "clone", origin, cloneLoc)
	clone.Stdout = os.Stdout
	clone.Stderr = os.Stderr
	return cloneLoc, clone.Run()
}

func checkoutAPISpecsVer(target, repoLocation string) error {
	checkout := exec.Command("git", "checkout", target)
	checkout.Stdout = os.Stdout
	checkout.Stderr = os.Stderr
	checkout.Dir = repoLocation
	return checkout.Run()
}

// getDefaultOutputLocation returns the path to the local enlistment of the Azure SDK for Go.
// If there is no local enlistment, it creates a temporary directory for the output.
func getDefaultOutputLocation() string {
	goPath := os.Getenv("GOPATH")

	if goPath != "" {
		sdkLocation := path.Join(goPath, "src", "github.com", "Azure", "azure-sdk-for-go")
		if isGitDir(sdkLocation) {
			return filepath.Clean(sdkLocation)
		}
	}

	if tempDir, err := ioutil.TempDir("", "azure-sdk-for-go"); err == nil {
		return filepath.Clean(tempDir)
	}
	return "./"
}

func isGitDir(dir string) bool {
	retval := false
	if children, err := ioutil.ReadDir(dir); err == nil {
		for _, child := range children {
			if child.IsDir() && child.Name() == ".git" {
				retval = true
				break
			}
		}
	}
	return retval
}

func generate(swag SwaggerFile, outputRootPath, specsRootPath string, output io.Writer) (outputDir string, err error) {
	if output == nil {
		output = ioutil.Discard
	}

	var namespace string
	namespace, err = getNamespace(swag)
	if err != nil {
		return
	}
	outputDir = path.Clean(filepath.Join(outputRootPath, namespace))

	autorest := exec.Command(
		"autorest",
		"-Input", swag.Path,
		"-CodeGenerator", "Go",
		"-Header", "MICROSOFT_APACHE",
		"-Namespace", namespace[strings.LastIndex(namespace, "/")+1:],
		"-OutputDirectory", outputDir,
		"-Modeler", "Swagger",
		"-pv", version.String(),
		"-SkipValidation")
	autorest.Stdout = output
	autorest.Stderr = output
	autorest.Dir = specsRootPath

	err = autorest.Run()
	return
}

func format(path string) (err error) {
	formatter := exec.Command(
		"go",
		"fmt",
		path,
	)
	return formatter.Run()
}

func vet(path string) (err error) {
	vetter := exec.Command(
		"go",
		"vet",
		path,
	)
	return vetter.Run()
}

// getNamespace takes a swagger and finds the appropriate namespace to be fed into Autorest.
var getNamespace = func() func(SwaggerFile) (string, error) {
	//Defining the Regexp strategies statically like this helps improve perf by ensuring they only get compiled once.
	armPattern := getAzureSpecsPattern("resource-manager")
	dataPattern := getAzureSpecsPattern("data-plane")

	return func(swag SwaggerFile) (result string, err error) {
		strategies := []struct {
			pattern *regexp.Regexp
			plane   string
		}{
			{armPattern, "management"},
			{dataPattern, ""},
		}

		debugLog.Println("Inspecting path:", swag.Path)

		for _, currentStrategy := range strategies {
			results := currentStrategy.pattern.FindAllStringSubmatch(swag.Path, -1)
			if len(results) == 0 {
				continue
			}

			pkg := results[0][1]
			version := results[0][3]
			filename := results[0][4]
			filename = strings.ToLower(filename[:1]) + filename[1:]

			namespace := []string{"service", pkg}
			if currentStrategy.plane != "" {
				namespace = append(namespace, currentStrategy.plane)
			}
			namespace = append(namespace, version, filename)
			result = strings.Replace(strings.Join(namespace, "/"), `\`, "/", -1)
			return
		}
		err = fmt.Errorf("%s is not in a recognized namespace format", swag.Path)
		return
	}
}()

func getAzureSpecsPattern(plane string) *regexp.Regexp {
	elements := []string{
		`specification`,
		`(?P<package>[\w\-\d\.]+)`,
		plane,
		`[Mm]icrosoft\.(?P<category>[\w\-]+)`,
		`(?P<apiVersion>v?\d{4}-\d{2}-\d{2}(?:[\-\.][\w\d\-\.]+)?)`,
		`(?P<group>[\w\-\d\.]+)\.json`,
	}
	return regexp.MustCompile(strings.Join(elements, `[/\\]`))
}
