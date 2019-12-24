package main

import (
	"os"
	"strconv"
	"syscall"
)

func main() {
	filename := "template/"
	for try := 0; try < 2; try++ {
		file, err := os.Create(filename + strconv.Itoa(try))
		defer file.Close()
		if err == nil {
			return
		}
		if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOSPC {
			//deleteTempFiles()  // Recover some space.
			//data, err := ioutil.ReadAll(file)
			//if err != nil {
			//	fmt.Printf("11-error: %v\n", err)
			//	return
			//}
			//fmt.Printf(string(data))
			continue
		}
		return
	}
}
