package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"strconv"
)

var con net.Conn

func getNumbersFromInput(message string) (int64, error) {
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	return strconv.ParseInt(re.FindString(message), 10, 64)
}

func sendMessage(message string) (answer string) {
	con.Write([]byte(message))
	reply := make([]byte, 1024)
	_, err := con.Read(reply)
	if err != nil {
		fmt.Println(err)
	}
	return string(reply)
}

func treatMessage(message string, position *string) {
	switch *position {
	case "start":
		if message[0:1] == "1" || message[0:1] == "2" {
			*position = message[0:1]
		} else {
			fmt.Println(message)
			fmt.Println("Invalid option...")
		}

	case "1":
		limit, err := getNumbersFromInput(message)
		if err != nil {
			fmt.Println("Invalid number..")
			return
		}
		for x := 1; x < int(limit); x++ {
			fmt.Printf("The number of steps is: %s for: %d\n", sendMessage(strconv.Itoa(x)+"\r\n"), x)
		}
		*position = "start"

	case "2":
		fmt.Println("Answer: " + sendMessage(message))
		*position = "start"

	}
}

func positionMessage(position string) {
	fmt.Println("------------------------")
	switch position {
	case "start":
		fmt.Println("Ways of testing \n 1 - range of numbers \n 2 - single value")

	case "1":
		fmt.Println("write the number to query the values from 1 to: ")

	case "2":
		fmt.Println("write the number to see the number of steps: ")

	}
	fmt.Println("------------------------")
}

func main() {

	fmt.Println("Starting script..")
	position := "start"
	for {
		positionMessage(position)
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		treatMessage(line, &position)
	}
}

func init() {
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		log.Fatal(err)
	}
	con = conn
}
