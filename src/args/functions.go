package args

import (
	"os"

	"com.github.hugovallada/file-comparer/src/files"
)

type Source string

const (
	dir      Source = "dir"
	multiple Source = "multiple"
	unknow   Source = "unknow"
)

var (
	defaultArgs = []string{"./files/ancestral.txt", "./files/atual.txt"}
)

func ValidateArgs() Source {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case string(dir):
			return dir
		case string(multiple):
			return multiple
		default:
			return unknow
		}
	}
	return unknow
}

func GetArgs(source Source) Arg {
	switch source {
	case dir:
		fileNames := files.ReadFileNamesInDirs(os.Args[2])
		return Arg{NumberOfFiles: len(fileNames), FileNames: fileNames}
	case multiple:
		return Arg{NumberOfFiles: len(os.Args[2:]), FileNames: os.Args[2:]}
	default:
		return Arg{NumberOfFiles: 2, FileNames: defaultArgs}
	}
}
