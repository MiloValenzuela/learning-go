package main

import "log"

func main() {
	var fistLine = "Once uppne a time"

	for i, l := range fistLine {
		log.Println(i, ":", l)
	}

}
