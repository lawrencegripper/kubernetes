package main

// Copyright 2017 Microsoft Corporation
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/marstr/collection"
	"github.com/marstr/swagger"
)

type SwaggerFile struct {
	swagger.Swagger
	Path string
}

// SwaggerFinder will enumerate all Swaggers in a particular directory.
type SwaggerFinder struct {
	output *log.Logger
	Root   string
}

// NewSwaggerFinder creates a new instance of SwaggerFinder which will search the
// directory `targetDir`.
func NewSwaggerFinder(targetDir string, output io.Writer) (constructed *SwaggerFinder, err error) {
	constructed = &SwaggerFinder{
		Root:   targetDir,
		output: log.New(output, "[FINDER] ", log.Ltime),
	}
	return
}

// Enumerate lists out all instances of Swagger files in the SwaggerFinder's `Root`.
func (f *SwaggerFinder) Enumerate(cancel <-chan struct{}) collection.Enumerator {
	retval := make(chan interface{})

	go func() {
		defer close(retval)

		seen := map[string][]string{}

		seenContains := func(needle SwaggerFile) bool {
			if previouslySeen, ok := seen[needle.Info.Title]; ok {
				for _, version := range previouslySeen {
					if version == needle.Info.Version {
						return true
					}
				}
			}
			return false
		}

		filepath.Walk(f.Root, func(path string, info os.FileInfo, err error) (result error) {
			if err != nil {
				return
			}

			if strings.ToLower(filepath.Ext(path)) == ".json" {
				var contents []byte
				if temp, err := ioutil.ReadFile(path); err == nil {
					f.output.Print("Inspecting File: ", path)
					contents = temp
				} else {
					return
				}

				var manifest SwaggerFile
				if err := json.Unmarshal(contents, &manifest); err != nil {
					f.output.Print(path, " was not a Swagger file.")
					return
				}
				manifest.Path = strings.TrimPrefix(path, f.Root)
				manifest.Path = strings.TrimLeft(manifest.Path, `/\`)

				title := manifest.Info.Title

				if title == "" {
					return
				}
				f.output.Print("Found! ", path, " is a Swagger.")
				if seenContains(manifest) {
					f.output.Print("Skipping because we've already seen a Swagger that matches ", path)
					return
				} else if versions, ok := seen[manifest.Info.Title]; ok {
					seen[manifest.Info.Title] = append(versions, manifest.Info.Version)
				} else {
					seen[manifest.Info.Title] = []string{manifest.Info.Version}
				}
				select {
				case retval <- manifest:
				case <-cancel:
					return
				}
			}
			return
		})
	}()

	return retval
}
