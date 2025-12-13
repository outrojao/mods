package utils

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
