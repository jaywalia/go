package main

import "fmt"

type NugetPackage struct {
	name []string
}

type DotNetPackage struct {
}

type DotNetSolution struct {
}

func main() {
	fmt.Println("Hello World!")

	// read config file
	// get directory name
	// recursively find all sln files
	// for each solution file, find all project files
	// list dependency of project file
	// for each project file
	// list dependencies on other projects
	// list dependencies on .NET libraries
	// list dependencies on .NET packages
}
