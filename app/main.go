package main

import (
	"fmt"
	"sync"

	"com.github.hugovallada/file-comparer/src/args"
	"com.github.hugovallada/file-comparer/src/data"
	exec "com.github.hugovallada/file-comparer/src/executor"
	"com.github.hugovallada/file-comparer/src/files"
)

func main() {
	exec.TimedExecution(func() {
		source := args.ValidateArgs()
		args := args.GetArgs(source)

		channel := make(chan data.Data, args.NumberOfFiles)

		var ww sync.WaitGroup

		ww.Add(args.NumberOfFiles)

		fileReaders := files.GenerateFileReaderFromArgs(args.FileNames, channel, &ww)

		for _, fileReader := range fileReaders {
			go files.ReadFile(fileReader)
		}

		ww.Wait()

		allData := getDatas(args.NumberOfFiles, channel)
		validateAllFiles(allData)
		fmt.Println("Sucesso")
	})
}

func getDatas(size int, channel chan data.Data) (datas []data.Data) {
	for i := 0; i < size; i++ {
		datas = append(datas, <-channel)
	}
	return
}

func validateAllFiles(datas []data.Data) {
	datas[0].Compare(datas[1:]...)
}
