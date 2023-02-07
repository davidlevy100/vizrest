package main

import "fmt"

func main() {

	c := CreateClient()

	m, err := GetBase(c, "http://127.0.0.1:8580")

	if err == nil {
		for key, val := range m {
			fmt.Println(key, val)
		}
	} else {
		fmt.Println(err)
	}

}
