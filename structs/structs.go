package main

import (
	"fmt"
)

// type MessageToSend struct {
// 	phoneNumber int
// 	message     string
// }

func main() {
	var numberFromUser int
	fmt.Println("Hi there! Can you give me a number?")
	fmt.Scanln(&numberFromUser)
	fmt.Printf("You've entered: %d\n", numberFromUser)
	fmt.Printf("Is your number divisible by 2? %t\n", (numberFromUser%2 == 0))
	// var m = MessageToSend{
	// 	phoneNumber: 12312312,
	// 	message:     "test",
	// }
	// m.modify()
	// fmt.Println(m)
}

// func (m *MessageToSend) modify() {
// 	m.message = m.message + "1"
// }
