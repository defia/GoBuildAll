package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
)

const PROJECTFILENAME = "project.ini"

var threads = runtime.NumCPU()
var ch = make(chan bool, threads)

func main() {
	if len(os.Args) < 2 {
		RunGoDefault("")
		return
	}
	if strings.TrimSpace(os.Args[1]) != "build" {
		RunGoDefault("")
		return
	}
	runtime.GOMAXPROCS(1)
	filename, err := FindFile(PROJECTFILENAME)
	if err != nil {
		RunGoDefault("")
		return
	}
	c, err := Loadconfig(filename)
	if err != nil {
		RunGoDefault("")
		return

	}
	// if error occurs when building current package, don't build main package
	if RunGoDefault("") != nil {
		return
	}

	wg := &sync.WaitGroup{}
	for i := range c {
		if c[i].PackageRoot != "" && IsFolderExist(c[i].PackageRoot) {
			wg.Add(1)
			ch <- true
			go func(c *Package) {
				fmt.Println("processing ", c.PackageName)
				//TODO:the output of each go process may mess due to parallelize build,will it be necessary to group the stdout/stderr?
				RunGoDefault(c.PackageRoot)
				<-ch
				wg.Done()
			}(c[i])

		} else {
			fmt.Println(c[i].PackageName, " err in path")
		}

	}
	wg.Wait()
}

func RunGo(stdout, stderr io.Writer, path string) error {
	cmd := exec.Command("go", os.Args[1:]...)
	cmd.Dir = path
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	return cmd.Run()
}

func RunGoDefault(path string) error {
	return RunGo(os.Stdout, os.Stderr, path)
}
