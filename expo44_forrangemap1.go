package main

import "fmt"

func main() {
	var weekday map[int]string
	weekday = map[int]string{
		0: "Sunday",
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday"}

	for key, value := range weekday {
		fmt.Printf("key is: %d - value is: %s\n", key, value)
		switch value {
		case "Tuesday", "Hollyday":
			fmt.Println("Bingo! %s", value)
		default:
			fmt.Println("Nothing!")
		}
	}
}
