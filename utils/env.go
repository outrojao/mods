package utils

import "os"

// Função para pegar valor de env e validar
func EnvOrDefault(key, defaultValue string) string {
	v := os.Getenv(key)
	if v == "" {
		v = defaultValue
	}
	return v
}
