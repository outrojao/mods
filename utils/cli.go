package utils

import "fmt"

// Cria um menu a partir de uma lista de opções exibindo um título opcional
func CreateMenu(options []string, menuTitle ...string) {
	if len(menuTitle) > 0 {
		title := menuTitle[0]
		fmt.Println(title)
	}

	for i, option := range options {
		fmt.Printf("%d. %s\n", i+1, option)
	}
}
