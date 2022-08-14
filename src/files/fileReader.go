package files

import (
	"sync"

	"com.github.hugovallada/file-comparer/src/data"
)

type FileReader struct {
	FileName string
	Channel  chan data.Data
	WG       *sync.WaitGroup
}

func GenerateFileReaderFromArgs(fileNames []string, channel chan data.Data, wg *sync.WaitGroup) (fileReaders []FileReader) {
	for _, file := range fileNames {
		fileReaders = append(fileReaders, FileReader{FileName: file, Channel: channel, WG: wg})
	}
	return
}
