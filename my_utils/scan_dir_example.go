package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var wg sync.WaitGroup

func walkDir(dir string)  {
	defer wg.Done()

	visit := func(path string, f os.FileInfo, err error) error {
		if f.IsDir() && path != dir {
			wg.Add(1)
			go walkDir(path)
			return filepath.SkipDir
		}

		if f.Mode().IsRegular() {
			fmt.Printf("Visited:%s; -File name:%s; -Size:%d bytes\n", path, f.Name(), f.Size())
		}
		return nil
	}

	_ = filepath.Walk(dir, visit)
}

func main() {
	t := time.Now()
	path := "D:\\www"

	wg.Add(1)
	walkDir(path)
	wg.Wait()

	elapsed := time.Since(t)
	fmt.Println("扫描完成，耗时：", elapsed)
}
