package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
)

func compute() {
	fp, err := os.Create("cpu.pprof")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	fmt.Println("compute function called, to profile")

	pprof.StartCPUProfile(fp)
	defer pprof.StopCPUProfile()
}

func memory() {
	fp, err := os.Create("memory.pprof")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	fmt.Println("memory function called, to profile")

	runtime.GC()
	defer pprof.WriteHeapProfile(fp)
}


func main() {
	fmt.Println("cpu and memory profiling")
	compute()
	memory()
}
