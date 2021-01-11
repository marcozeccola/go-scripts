//Compila file c in binario con lo stesso nome
package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	if len(os.Args) < 2 {

		panic("Nessun file passato")

	} else {

		programma := os.Args[1]
		eseguibile := strings.TrimSuffix(programma, ".c")
		cmd := exec.Command("gcc", programma, "-o", eseguibile)

		err := cmd.Run()

		if err != nil {
			panic("File non compilabile")
		}
	}
}
