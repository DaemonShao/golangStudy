package main

import "fmt"

type Vertex struct {
	Lat , Long float64
}

var m map[string]string

func main()  {
	m  = make(map[string]string)
	m["Bell Labs"] = "aaaa"
	fmt.Println(m["Bell Labs"])
}