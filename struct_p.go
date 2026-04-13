package main

import "fmt"

func main() {
	type person struct {
		firstName string
		lastName  string
	}

	type book struct {
		title  string
		author *person
	}

	p := person{"Van", "Go"}
	b := book{"Hi", &p}
	fmt.Println(p, b)
	b.author.lastName = "Gone"
	fmt.Println(p, b)
}


// {Van Go} {Hi 0x35b043f12020}
// {Van Gone} {Hi 0x35b043f12020}
