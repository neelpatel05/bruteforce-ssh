package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func cleanup() {
	if r:=recover(); r!=nil {
		log.Fatal(r)
	}
}

func one(reader *bufio.Reader) {

}

func two(reader *bufio.Reader) {

}

func main() {
	defer cleanup()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$: 1. Create wordlist and attack 2. Use own wordlist and attack 3. Exit - press Ctrl + c $: ")
		choice,err := reader.ReadString("\n")
		if err!=nil {
			panic(err.Error())
		}
		switch(choice) {
		case "1":
			one(reader)
		case "2":
			two(reader)
		}
	}
}
