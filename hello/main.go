package main

import (
	"fmt"
)

const (
	langRu = "ru"
	langEn = "en"

	helloPrefixEn = "Hello, "
	helloPrefixRu = "Привет, "
	worldEn       = "World"
	worldRu       = "Мир"
)

func Hello(name, lang string) string {
	if name == "" {
		switch lang {
		case langRu:
			name = worldRu
		case langEn:
			name = worldEn
		}
	}
	switch lang {
	case langRu:
		return helloPrefixRu + name
	case langEn:
		return helloPrefixEn + name
	default:
		return helloPrefixEn + name
	}
}

func main() {
	fmt.Println(Hello("Kirill", "en"))
}
