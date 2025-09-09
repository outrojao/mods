package utils

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
)

// Função genérica para obter entrada do usuário
func GetUserInput[T any](message string) (input T) {
	scanner := bufio.NewScanner(os.Stdin)
	letterRegexp := regexp.MustCompile(`^[A-Za-zÀ-ÿ\s]+$`)
	for {
		fmt.Print(message)
		if !scanner.Scan() {
			fmt.Println("Entrada inválida. Tente novamente.")
			continue
		}
		raw := scanner.Text()

		var val any
		switch reflect.TypeOf(input).Kind() {
		case reflect.String:
			matched := letterRegexp.MatchString(raw)
			if !matched {
				fmt.Println("Por favor, digite apenas letras.")
				continue
			}
			val = raw
		case reflect.Int:
			i, err := strconv.Atoi(raw)
			if err != nil {
				fmt.Println("Entrada inválida. Tente novamente.")
				continue
			}
			val = i
		case reflect.Float64:
			f, err := strconv.ParseFloat(raw, 64)
			if err != nil {
				fmt.Println("Entrada inválida. Tente novamente.")
				continue
			}
			val = f
		case reflect.Float32:
			f, err := strconv.ParseFloat(raw, 32)
			if err != nil {
				fmt.Println("Entrada inválida. Tente novamente.")
				continue
			}
			val = float32(f)
		default:
			fmt.Println("Tipo não suportado.")
			return
		}

		input = val.(T)
		return
	}
}

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
