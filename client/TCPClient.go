package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func helpPrinter() {
	text := "\n\nPOSSIBLE OPTIONS of Quizz Game" +
		"\n ---- LIST - lists all quizes" +
		"\n ---- SOLVE <quizName> - start solving quiz" +
		"\n ---- RANKING <quizName>" +
		"\n ---- STOP - abort connection to the server"
	fmt.Println(text)
	//os.Exit(0)
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	CONNECT := arguments[1]
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your nickname:")
	nickname, _ := reader.ReadString('\n')
	//TODO:work it
	fmt.Print("Current nickname:" + nickname)
	helpPrinter()
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')

		fmt.Fprintf(c, text+"\n")

		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}
