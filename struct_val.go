package main

import "fmt"

func main() {
	type person struct {
		firstName string
		lastName  string
	}

	type book struct {
		title  string
		author person
	}

	p := person{"Van", "Go"}
	b := book{"Hi", p}
	fmt.Println(p, b)
	b.author.lastName = "Gone"
	fmt.Println(p, b)
}


// {Van Go} {Hi {Van Go}}
// {Van Go} {Hi {Van Gone}}
