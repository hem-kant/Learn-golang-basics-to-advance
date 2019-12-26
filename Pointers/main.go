package main

import "fmt"

func main() {
	myslice := []string{"Hi", "there", "how", "are", "you"}
	update(myslice)
	fmt.Println(myslice)
}

func update(s []string) {
	s[0] = "Bye"
}
