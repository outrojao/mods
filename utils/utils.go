package utils

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// Função genérica para obter entrada do usuário
func GetUserInput[T any](message string) (input T) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(message)
		if !scanner.Scan() {
			fmt.Println("Entrada inválida. Tente novamente.")
			continue
		}
		raw := strings.TrimSpace(scanner.Text())
		if raw == "" {
			fmt.Println("Entrada não pode ser vazia. Tente novamente.")
			continue
		}

		var val any
		switch reflect.TypeFor[T]().Kind() {
		case reflect.String:
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

// Função genérica para obter entrada do usuário, aceitando valor nulo
func GetUserInputAcceptNull[T any](message string) (input T) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(message)
		if !scanner.Scan() {
			fmt.Println("Entrada inválida. Tente novamente.")
			continue
		}
		raw := strings.TrimSpace(scanner.Text())

		var val any
		if raw == "" {
			var zero T
			return zero
		}

		switch reflect.TypeFor[T]().Kind() {
		case reflect.String:
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

		v := val.(T)
		input = v
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

// Função para filtragem de uma lista genérica de dados
func Filter[T any](data []T, f func(T) bool) []T {
	fltd := make([]T, 0)
	for _, e := range data {
		if f(e) {
			fltd = append(fltd, e)
		}
	}
	return fltd
}
