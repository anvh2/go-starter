# go-starter

[![Build Status](https://travis-ci.org/bakaoh/go-starter.svg?branch=master)](https://travis-ci.org/bakaoh/go-starter) [![codecov](https://codecov.io/gh/bakaoh/go-starter/branch/master/graph/badge.svg)](https://codecov.io/gh/bakaoh/go-starter)


# How to get started with Golang

## Overview
Duration: 1:00

In this tutorial, you will learn how to get started with Golang and use it for new backend services. All the topics below is just an overview. More detail tutorials will come later.

### What you'll learn

* How to install Golang and setup IDE 
* How to manage dependencies with Golang
* How to test Golang services
* How to build a minimal toolkit for yourself

### What you'll need

* The `git` command-line client (that you can install by running `sudo apt install git`)
* A [GitHub](https://github.com) account, for library publication

Let's get started!

## Install Golang
Duration: 5:00

- Download [archive](https://golang.org/dl/) and extract

```bash
$ sudo tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
```

- Edit your ~/.bash_profile to add the following line

```bash
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
```

### First application

- Make the directory `src/hello`

```
$ mkdir $GOPATH/src/hello
$ gedit $GOPATH/src/hello/hello.go
```

- Create a file named `hello.go`

```go
package main

import "fmt"

func main() {
    fmt.Printf("hello, world\n")
}
```

- Build and execute

```bash
$ cd $HOME/go/src/hello
$ go build
$ ./hello
```

You should see `hello, world` printed out!

## Setup IDE
Duration: 5:00

- Download and install Visual Studio Code [here](https://code.visualstudio.com/Download)
- Open VS Code, press `Ctrl+P`, paste the following command `ext install ms-vscode.Go`, and press `enter`
- Install Go tools, from terminal

    * For symbol search in the current workspace
`go get -u -v github.com/ramya-rao-a/go-outline`
    * For symbol search in the current file
`go get -u -v github.com/acroca/go-symbols`
    * For auto-completion
`go get -u -v github.com/nsf/gocode`
    * For the Go to Definition feature
`go get -u -v github.com/rogpeppe/godef`
    * For the documentation that appears on hover
`go get -u -v golang.org/x/tools/cmd/godoc`
    * For the documentation that appears on hover
`go get -u -v github.com/zmb3/gogetdoc`
    * For linting
`go get -u -v github.com/golang/lint/golint`
    * For auto-completion of unimported packages
`go get -u -v github.com/uudashr/gopkgs/cmd/gopkgs`
    * For renaming symbols
`go get -u -v golang.org/x/tools/cmd/gorename`
    * For formatting code
`go get -u -v sourcegraph.com/sqs/goreturns`
    * For formatting code
`go get -u -v golang.org/x/tools/cmd/goimports`
    * For generating unit tests
`go get -u -v github.com/cweill/gotests/...`
    * For the Find all References feature
`go get -u -v golang.org/x/tools/cmd/guru`
    * For generating stubs for interfaces 
`go get -u -v github.com/josharian/impl`
    * For running current file in the Go playground
`go get -u -v github.com/haya14busa/goplay/cmd/goplay`
    * For filling a struct literal with default values
`go get -u -v github.com/davidrjenni/reftools/cmd/fillstruct`

## Manage Dependencies
Duration: 5:00

### Go get

Downloads the packages named by the import paths, along with their dependencies. It then installs the named packages.

```bash
-GOPATH/
├── src/
|   ├── github.com/
|   ├── golang.org/
|   └── hello/
├── bin/
|   └── hello
└── pkg/ 
    └── linux_amd64/
        ├── github.com/
        └── golang.org/
```

### Dep

Dependency management for Go.

```
$ curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
$ dep init
$ dep ensure -add github.com/foo/bar github.com/baz/quu
$ dep ensure
$ dep status
```

```bash
-GOPATH/
└── src/
    ├── github.com/
    ├── golang.org/
    └── hello/
        ├── vendor/
        |   └── golang.org/
        ├── Gopkg.lock
        ├── Gopkg.toml
        ├── main.go
        └── main_test.go
```

## Test and Benchmark
Duration: 5:00

### Test
Any function of the form

```go
func TestXxx(*testing.T)
```

- Run test

```bash
$ go test -v ./... -cover
```

- (Optional) Install `assert` package

```bash
$ go get github.com/stretchr/testify/assert
```

- Fuzz test
```bash
$ go get github.com/dvyukov/go-fuzz
```

- Mock

```bash
$ go get github.com/golang/mock/gomock
$ go install github.com/golang/mock/mockgen
$ mockgen -source=foo.go 
```

### Benchmark 
Functions of the form

```go
func BenchmarkXxx(*testing.B)
```

- Run benchmark

```bash
$ go test -benchmem -bench ^.*$
```

### Profiling

```bash
$ sudo apt-get install graphviz
$ go test -bench=. -benchmem -cpuprofile cpu.prof
$ go tool pprof cpu.prof
(pprof) top10 -cum
(pprof) web
```

- ApacheBench

```bash
$ ab -k -c 8 -n 100000 "http://localhost:8088/ping"
$ go tool pprof http://localhost:8080/debug/pprof/heap
$ go tool pprof http://localhost:8080/debug/pprof/profile
$ go tool pprof http://localhost:8080/debug/pprof/block
```

## Webserver
Duration: 5:00

### Example

```go
package main

import (
    "net/http"
    _ "net/http/pprof"
)

func hiHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hi"))
}

func main() {
    http.HandleFunc("/", hiHandler)
    http.ListenAndServe(":8080", nil)
}
```

## Microservice
Duration: 5:00

### Thrift
- Install thrift package

```bash
$ go get git.apache.org/thrift.git/lib/go/thrift
```

- Gen go thrift

```bash
$ thrift -r --gen go scibe.thrift
```

### gRPC
- Install package

```bash
$ go get -u google.golang.org/grpc
$ go get -u github.com/golang/protobuf/protoc-gen-go
```

- Gen gRPC code

```bash
$ protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld
```

### Sql
- MySql drive `go get github.com/go-sql-driver/mysql`
- Sqlite3 drive `go get github.com/mattn/go-sqlite3`
- Postgres drive `go get github.com/lib/pq`
- Orm `go get github.com/jinzhu/gorm`

### Kafka
### Redis
### Connection Pool?

## Utils
Duration: 5:00

### Log
- log
- glog
- zap `go get github.com/uber-go/zap`

### Config
- flag
- gflag
- go-toml `go get github.com/pelletier/go-toml`
- viper `go get github.com/spf13/viper`

## More infos

* [Golang doc](https://golang.org/doc/)
* [Profiling and optimizing go web application](https://artem.krylysov.com/blog/2017/03/13/profiling-and-optimizing-go-web-applications/)
* [Dep doc](https://golang.github.io/dep/docs/introduction.html)