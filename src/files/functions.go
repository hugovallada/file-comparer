package files

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"com.github.hugovallada/file-comparer/src/data"
)

func ReadFile(fileReader FileReader) {
	defer fileReader.WG.Done()

	file, err := os.OpenFile(fileReader.FileName, os.O_RDONLY, os.ModeAppend)

	if err != nil {
		panic(fmt.Sprintf("Não foi possível ler o arquivo: %v", err))
	}

	defer file.Close()

	reader := bufio.NewScanner(file)

	reader.Split(bufio.ScanLines)

	var lines []string
	for reader.Scan() {
		lines = append(lines, strings.TrimSpace(reader.Text()))
	}

	data := data.Data{NumberOfLines: len(lines), Content: lines}
	fileReader.Channel <- data
}

func ReadFileNamesInDirs(path string) []string {
	if isDir(path) {
		err := os.Chdir(path)

		if err != nil {
			fmt.Println("Aqui")
			panic(err)
		}

		files, err := os.ReadDir(".")

		if err != nil {
			panic(err)
		}

		var fileNames []string
		for _, file := range files {
			if !file.IsDir() {
				fileNames = append(fileNames, file.Name())
			}
		}
		return fileNames
	}
	panic("Not a dir")
}

func isDir(path string) bool {
	fileInfo, _ := os.Stat(path)

	return fileInfo.IsDir()
}
