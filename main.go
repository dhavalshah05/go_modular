package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	filesToProcess := []string{"file1", "file2", "file3", "file4", "file5"}
	processedFiles := &[]string{}
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	for _, file := range filesToProcess {
		wg.Add(1)
		go processFile(file, processedFiles, wg, mutex)
	}

	wg.Wait()
	fmt.Println(*processedFiles)
}

func processFile(file string, processedFiles *[]string, wg *sync.WaitGroup, mutex *sync.Mutex) {
	time.Sleep(3 * time.Second)
	result := file + "processed"
	fmt.Println(result)
	mutex.Lock()
	*processedFiles = append(*processedFiles, result)
	mutex.Unlock()
	wg.Done()
}
