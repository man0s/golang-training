package main

import (
	"fmt"
	"strings"
)

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for k, v := range c {
		fmt.Println("Hex code for", strings.ToUpper(k), "is", strings.ToUpper(v))
	}
}
