package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

type Data struct {
	NumberOfLines int
	Content       []string
}

func (d *Data) validate(data Data) {
	if d.NumberOfLines != data.NumberOfLines {
		panic(fmt.Sprintf("O número de linhas é diferente: Esperava %d, mas recebeu %d\n", d.NumberOfLines, data.NumberOfLines))
	}

	for index, line := range d.Content {
		if line != data.Content[index] {
			panic(fmt.Sprintf("O conteúdo da linha é diferente: Esperava %s, mas recebeu %s\n", line, data.Content[index]))
		}
	}
}

func (d *Data) IsEqualTo(datas ...Data) {
	for _, data := range datas {
		d.validate(data)
	}
}

var wg sync.WaitGroup

func executor(fn func()) {
	startTime := time.Now()
	fn()
	fmt.Println("Duração:", time.Since(startTime))
}

func main() {
	executor(func() {
		files := []string{"./files/ancestral.txt", "./files/atual.txt"}
		size := len(files)

		wg.Add(size)
		channel := make(chan Data, size)

		go readFile(files[0], channel)
		go readFile(files[1], channel)

		wg.Wait()

		validateFiles(getData(size, channel))

		fmt.Println("Sucesso")
	})
}

func getData(size int, channel chan Data) (datas []Data) {
	for i := 0; i < size; i++ {
		datas = append(datas, <-channel)
	}
	return
}

func validateFiles(datas []Data) {
	data1, data2 := datas[0], datas[1]

	data1.validate(data2)
	// if datas[0].NumberOfLines != datas[1].NumberOfLines {
	// 	panic("Numero de linhas diferentes")
	// }

	// for ind, line := range datas[0].Content {
	// 	if line != datas[1].Content[ind] {
	// 		panic("Conteudo de linha diferentes")
	// 	}
	// }
}

func readFile(fileName string, channel chan Data) {
	defer wg.Done()
	file, err := os.OpenFile(fileName, os.O_RDONLY, os.ModeAppend)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	reader := bufio.NewScanner(file)

	reader.Split(bufio.ScanLines)

	var lines []string

	for reader.Scan() {
		lines = append(lines, strings.TrimSpace(reader.Text()))
	}

	data := Data{len(lines), lines}

	channel <- data
}
