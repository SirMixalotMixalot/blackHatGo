package main

import (
	"fmt"
)
type Person struct {
	name string 
	age int
}
func strlen(s string, c chan int) {
	c <- len(s);
}
func (p *Person) hello() {
	fmt.Println("hello, my name is", p.name);
}
func main() {
	c := make(chan int);
	var words = []string{"hello", "worlds"};
	for _, word := range words {
		go strlen(word, c);
		fmt.Println(word);
	}
	x, y := <-c, <-c;
	fmt.Println(x, y, x + y); //I cannot escape
	//the semicolons





	
}