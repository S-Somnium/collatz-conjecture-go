package main

import (
	"bufio"
	"collatz-conjecture-go/algorithms"
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
)

func getNumbersFromInput(message string) (int64, error) {
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	return strconv.ParseInt(re.FindString(message), 10, 64)
}

func readMessages(connection net.Conn) {
	defer connection.Close()
	for {
		input, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
			fmt.Println("end of connection")
			return
		}
		fmt.Print("Received message:", input)
		number, err := getNumbersFromInput(input)
		if err != nil {
			connection.Write([]byte("\nArgument is not valid! Expected int\n"))
			fmt.Println(err)
		} else {
			steps := algorithms.CollatzStepsCalc(int(number))
			connection.Write([]byte(strconv.Itoa(steps)))
		}
	}
}

func startService(port string) {
	service, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	for {
		newConnection, err := service.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go readMessages(newConnection)
	}
}

func main() {
	startService("8081")
}
