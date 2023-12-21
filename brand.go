package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

var (
	ErrInvalidString = errors.New("Error: неверный формат")
	ErrSpam          = errors.New("Error: Спам")
)

// Функция проверки на спам
func EqualsBrand(strLeft, strRight string) bool {
	// оставляем только название бренда в правом поле
	strRight = strings.TrimSpace(strings.Map(IgnoreName, strRight))

	if (strLeft == strRight) || (strRight == "") {
		return true
	}

	return false
}

// игнорируем название (символы кириллицы)
func IgnoreName(r rune) rune {
	if (r >= 'А' && r <= 'я') || (r == 'Ё' || r == 'ё') {
		return -1
	}
	return r
}

// функция проверки ввода корректной строки и отсутствия спама
func CheckSpam(str string) (string, error) {
	if strings.Count(str, "|") != 1 {
		return "", ErrInvalidString
	}

	splitStr := strings.Split(str, "|")

	left := strings.TrimSpace(splitStr[0])
	right := strings.TrimSpace(splitStr[1])

	// проверяем отсутствие названия в поле бренда
	for _, char := range left {
		if unicode.Is(unicode.Cyrillic, char) { // Проверяем, является ли символ кириллицей
			return "", ErrSpam
		}
	}

	if EqualsBrand(left, right) {
		return "Ok", nil
	} else {
		return "", ErrSpam
	}
}

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		str := scanner.Text()
		result, err := CheckSpam(str)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(result)
		}
	}
}
