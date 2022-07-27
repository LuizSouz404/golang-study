package main

import "fmt"

func main() {
	tables := map[string]string{
		"tabela0": "TABLE0",
		"tabela1": "TABLE1",
	}

	dados := []map[string]string{{
		"data": "TABLE1", "hash": "123456",
	}, {
		"data": "TABLE0", "hash": "123456",
	}}

	for _, msg := range dados {
		value, ok := tables[msg["data"]]

		if ok {
			fmt.Printf(value)
		}
	}
}
