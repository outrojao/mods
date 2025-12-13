package utils

import "time"

// Função para pegar a timestamp atual em segundos
func GetTimestamp() int64 {
	return time.Now().Unix()
}
