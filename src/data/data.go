package data

import "fmt"

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

func (d *Data) Compare(datas ...Data) {
	for _, data := range datas {
		d.validate(data)
	}
}
