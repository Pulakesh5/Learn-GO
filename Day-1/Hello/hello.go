package main

import (
	"Day-1/greetings"
	"fmt"
)

func main() {
	// fmt.Println("Hello, World!")
	// log.SetPrefix("Greetings, Pulakesh: ")
	// log.SetFlags(0)

	fmt.Println(greetings.Hello("Pulakesh"))

	// names := []string{"Pulakesh", "Swastik", "Bikrant", "Sohan"}
	// messages, err := greetings.Hellos(names)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for name, message := range messages {
	// 	fmt.Println("Greetings for ", name, " : ", message)
	// }
}
