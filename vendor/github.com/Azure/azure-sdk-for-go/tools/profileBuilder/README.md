# Profile Builder

## Installation

*Note:* These installation notes assume that you have [Go 1.9](https://blog.golang.org/go1.9) or higher, [Glide](http://github.com/Masterminds/glide), and [Git](https://git-scm.com/) installed.

The simplest version of installation is simple, just run the follwoing command:

``` bash
go get -u github.com/lawrencegripper/azure-sdk-for-go/tools/profileBuilder
```

If that causes you trouble, run the following commands:

``` bash
# bash
go get -d github.com/lawrencegripper/azure-sdk-for-go/tools/profileBuilder
cd $GOPATH/src/github.com/lawrencegripper/azure-sdk-for-go/tools/profileBuilder
glide install
go install
```

``` PowerShell
# PowerShell
go get -d github.com/lawrencegripper/azure-sdk-for-go/tools/profileBuilder
cd $env:GOPATH\src\github.com\Azure\azure-sdk-for-go\tools\profileBuilder
glide install
go install
```
Taking things a step further, if you'd like the profileBuilder to stamp the version of itself in the code it generates, you can install as below:

``` bash
# bash
go get -d github.com/lawrencegripper/azure-sdk-for-go/tools/profileBuilder
cd $GOPATH/src/github.com/lawrencegripper/azure-sdk-for-go/tools/profileBuilder
glide install
export currentCommit=$(git rev-parse HEAD)
go install -ldflags "-X main.version=$currentCommit"
```

``` PowerShell
# PowerShell
go get -d github.com/lawrencegripper/azure-sdk-for-go/tools/profileBuilder
cd $env:GOPATH\src\github.com\Azure\azure-sdk-for-go\tools\profileBuilder
glide install
$currentCommit = git rev-parse HEAD
go install -ldflags "-X main.version=$currentCommit"
```