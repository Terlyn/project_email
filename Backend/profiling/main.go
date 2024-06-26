package main

import (
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func main() {
	cpuFile, err := os.Create("cpu_profile.prof")
	if err != nil {
		log.Fatal("No se pudo crear archivo para perfil de CPU: ", err)
	}
	defer cpuFile.Close()

	if err := pprof.StartCPUProfile(cpuFile); err != nil {
		log.Fatal("No se pudo iniciar perfil de CPU: ", err)
	}
	defer pprof.StopCPUProfile()

	doHeavyWork()

	memFile, err := os.Create("heap_profile.prof")
	if err != nil {
		log.Fatal("No se pudo crear archivo para perfil de memoria: ", err)
	}
	defer memFile.Close()

	runtime.GC()
	if err := pprof.WriteHeapProfile(memFile); err != nil {
		log.Fatal("No se pudo escribir perfil de memoria: ", err)
	}
}

func doHeavyWork() {
	for i := 0; i < 1000000; i++ {
		_ = make([]byte, 1024)
	}
	time.Sleep(2 * time.Second)
}
