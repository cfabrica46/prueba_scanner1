package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	origen, err := os.OpenFile("origen.txt", os.O_RDWR, 0644)

	if err != nil {
		log.Fatal(err)
	}

	scaner := bufio.NewScanner(origen)

	for scaner.Scan() {

		fmt.Println(scaner.Text())

		fmt.Println("Salto de linea")

	}

}
