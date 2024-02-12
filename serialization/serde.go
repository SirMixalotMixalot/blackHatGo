package main

import (
	"encoding/xml"
	"fmt"
)

type Book struct {
	Author string `xml:"id,attr"`
	Title string `xml:"parent>child"`
}

func main() {
	b := Book{"John Morris", "The life of a sigma"}
	data, err := xml.Marshal(b);
	if err != nil {
		fmt.Printf("Error: %+v", err)
		return 
	}
	fmt.Println(string(data));
	xml.Unmarshal(data, &b);
	
}