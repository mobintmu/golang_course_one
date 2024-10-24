## Golang Roadmap

https://roadmap.sh/golang
https://roadmap.sh/backend

## Go Commmand

https://pkg.go.dev/cmd/go

```
go 
Go is a tool for managing Go source code.
```

## go mod init
The go mod init command is used in the Go programming language to initialize a new module in your project. It creates a go.mod file

## go mod tidy
go mod tidy is like tidying up your room but for your Go project dependencies. It ensures that your go.mod file has exactly what you need:

Adds missing modules: If your code imports a package that isn't in go.mod, go mod tidy will add it.

Removes unused modules: If there are dependencies listed in go.mod that your project doesn't actually use, it'll clean those out.

## go test 

go test is the command used in Go to run test files, typically those ending in _test.go. It automates the execution of tests, making sure your code works as expected. Here's a quick look at how it works:

go test

## go run

go run is a handy command that compiles and executes Go programs in a single step. Here's what happens:

Compiles the code: It turns your Go code into an executable.

Runs the executable: It immediately executes the compiled code.

```
go run main.go
go run *.go
```

## go build

To build your Go application for Linux in a single command, you can combine setting the environment variable and building the binary like this:
```
GOOS=linux go build -o first
ls
./first
 sudo chmod 744 ./first 
```

for windows:
```
GOOS=windows go build -o yourappname.exe
```

go build -o bin/hello


## how-to-build-and-install-go-programs

https://www.digitalocean.com/community/tutorials/how-to-build-and-install-go-programs

## tools

https://book.gofarsi.ir/toolchain/


## go env GOPATH

go env GOPATH

go env GOPATH shows the current value of the GOPATH environment variable. This variable defines the root directory of your workspace, which is where Go looks for your source files, binaries, and dependencies. It’s kind of like the main hub for all your Go projects.

## Why is it recommended to use `go build` instead of `go run` when running a Go app in production?

https://stackoverflow.com/questions/61060768/why-is-it-recommended-to-use-go-build-instead-of-go-run-when-running-a-go-ap


A Go application does not require the Go toolchain, just as a C application doesn't require gcc. If you use go build, you can just deploy a binary; if you use go run, you have to install the toolchain.
go run compiles the application unnecessarily every time it's run, increasing startup times.
go run forks the application to a new process, making process management unnecessarily complicated by obscuring PIDs and exit statuses.
go run has the potential to unexpectedly absorb code changes when you're only intending to run an application. Using go build only when you want a fresh binary allows you to re-run the same, consistent binary every time with no unexpected changes.
go run ignores *.syso files while go build recognizes and links them.

## go doc

go doc is your trusty reference guide within the Go programming environment. It displays documentation for Go packages, functions, types, and more right in your terminal. Think of it as your personal manual for all things Go code, keeping you in the know without having to leave your development environment. 

```
go doc fmt
go doc fmt.Println
```

https://pkg.go.dev/fmt


## go module dependency

https://www.youtube.com/watch?v=7xSxIwWJ9R4

dependency requirements are stored in go.mod file


## Adding a Remote Module as a Dependency

Go modules are distributed from version control repositories, commonly Git repositories. When you want to add a new module as a dependency to your own, you use the repository’s path as a way to reference the module you’d like to use. When Go sees the import path for these modules, it can infer where to find it remotely based on this repository path.

For this example, you willl add a dependency on the github.com/spf13/cobra library to your module. Cobra is a popular library for creating console applications, but we won’t address that in this tutorial.

https://www.digitalocean.com/community/tutorials/how-to-use-go-modules#adding-a-remote-module-as-a-dependency

```
go get github.com/spf13/cobra
```

https://github.com/spf13/cobra

Cobra is a library for creating powerful modern CLI applications.

Cobra is used in many Go projects such as Kubernetes, Hugo, and GitHub CLI to name a few. This list contains a more extensive list of projects using Cobra.


## Module in Golang

In Go, a module is a collection of related Go packages that are versioned together as a single unit. Modules are the basis for managing dependencies in Go projects.
Key Points:

    Definition: A module is defined by a `go.mod` file that specifies the module path and its dependencies.

    Dependency Management: Modules allow you to manage project dependencies efficiently, ensuring that your project uses the correct versions of external packages.

    Versioning: Modules are versioned, meaning you can use specific versions of packages and update them as needed.

Example go.mod File:

Here's a simple example of what a go.mod file might look like:
go

```
module github.com/yourusername/yourproject

go 1.18

require (
    github.com/spf13/cobra v1.2.1
    github.com/stretchr/testify v1.7.0
)
```

Benefits:

    Reproducibility: Ensures that builds are consistent by locking dependencies to specific versions.

    Isolation: Avoids dependency conflicts by managing each module's dependencies separately.

    Version Control: Facilitates the use of version control systems to track changes in dependencies over time.



## package in Golang

In Go, a package is a collection of related Go source files that are organized together under a single directory. Packages are the fundamental building blocks of Go programs, providing modularity and code reuse.
Key Features:

    Modularity: Packages allow you to organize your code into reusable and maintainable units.

    Encapsulation: They encapsulate functionality, enabling you to expose only the parts of the code you want to be accessible.

    Namespace: Packages provide a namespace to avoid naming conflicts.

Basic Structure:

    Package Declaration: Every Go source file begins with a package declaration.

    Import: You can import other packages to use their functionality.

Example:

Here's a simple example of a Go package:

```

// mathutil/mathutil.go
package mathutil

// Add function adds two integers and returns the result.
func Add(a, b int) int {
    return a + b
}

```

Usage:

You can use this package in another Go program as follows:


```
// main.go
package main

import (
    "fmt"
    "yourmodule/mathutil"
)

func main() {
    result := mathutil.Add(3, 4)
    fmt.Println("Result:", result)
}
```

In this example:

    The mathutil package is defined with a single function Add.

    The main package imports mathutil and uses the Add function.

Packages are essential for structuring Go programs and promoting code reuse.