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
		raw := scanner.Text()
		if IsEmpty(raw) {
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
			raw = strings.ReplaceAll(raw, ",", ".")
			f, err := strconv.ParseFloat(raw, 64)
			if err != nil {
				fmt.Println("Entrada inválida. Tente novamente.")
				continue
			}
			val = f
		case reflect.Float32:
			raw = strings.ReplaceAll(raw, ",", ".")
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
		raw := scanner.Text()
		if IsEmpty(raw) {
			var zero T
			return zero
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
			raw = strings.ReplaceAll(raw, ",", ".")
			f, err := strconv.ParseFloat(raw, 64)
			if err != nil {
				fmt.Println("Entrada inválida. Tente novamente.")
				continue
			}
			val = f
		case reflect.Float32:
			raw = strings.ReplaceAll(raw, ",", ".")
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

// Função para verificar se a string é vazio retirando os espaços
func IsEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}
