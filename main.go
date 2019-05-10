package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func cleanup() {
	if r:=recover(); r!=nil {
		log.Fatal(r)
	}
}

func bruteforce(filename string, reader *bufio.Reader) {

}

func one(reader *bufio.Reader) {
	defer cleanup()

	fmt.Print("$: Enter the string: ")
	string, err := reader.ReadString('\n')
	string = strings.TrimSpace(string)
	if err!=nil {
		panic(err.Error())
	}
	fmt.Print("$: Enter the length of password: ")
	length, err := reader.ReadString('\n')
	if err!=nil {
		panic(err.Error())
	}

	cmd := exec.Command("python", "generate_word.py", string, length)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err!=nil {
		panic(err.Error())
	}

	filename := "word-list.txt"
	bruteforce(filename, reader)
}

func two(reader *bufio.Reader) {

}

func main() {
	defer cleanup()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$: 1. Create wordlist and attack 2. Use own wordlist and attack 3. Exit - press Ctrl + c $: ")
		choice,err := reader.ReadString('\n')
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
