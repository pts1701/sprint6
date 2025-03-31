package service

import (
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// проверка является ли содержимое файла кодом Морзе
func isMorse(s string) bool {
	morseChar := ".-/"
	for _, ch := range s {
		if !strings.ContainsRune(morseChar, ch) && ch != ' ' {
			return false
		}

	}
	return true
}

// проверка можно ли содержимое файла сконвертировать в код Морзе
func canMorse(s string) bool {

	for _, ch := range s {
		var special string = ".,:?\\-/()\"\n"

		if (ch < 'А' || ch > 'Я') && (ch < 'а' || ch > 'я') && (ch < '0' || ch > '9') && ch != ' ' && ch != 'ё' && ch != 'Ё' && !strings.ContainsRune(special, ch) {
			return false
		}
	}
	return true
}

// функция проверяет контент и конвертирует его соответственно, либо возвращает ошибку
func ConvMorse(input string) (string, error) {

	switch {
	case isMorse(input):
		return morse.ToText(input), nil

	case canMorse(input):
		return morse.ToMorse(input), nil

	default:
		return "", fmt.Errorf("Incorrect input data")
	}

}
