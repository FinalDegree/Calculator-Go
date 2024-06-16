package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SplitStringAndGetSymbol(x string) ([]string, string) {
	var operation string
	var words []string
	if strings.Contains(x, "+") || strings.Contains(x, "-") || strings.Contains(x, "*") || strings.Contains(x, "/") {
		if strings.Contains(x, "+") {
			words = strings.Split(x, "+")
			operation = "+"
		} else if strings.Contains(x, "-") {
			words = strings.Split(x, "-")
			operation = "-"
		} else if strings.Contains(x, "*") {
			words = strings.Split(x, "*")
			operation = "*"
		} else if strings.Contains(x, "/") {
			words = strings.Split(x, "/")
			operation = "/"
		}
	} else {
		panic(errors.New("Введённая строка не является математическим выражением."))
	}
	return words, operation
}

func IsArabicNum(str string) bool {
	var b bool = false
	_, err := strconv.Atoi(str)
	if err == nil {
		b = true
	}
	return b
}

func IsRomanNum(str string, RomanNums []string) bool {
	var b bool = false
	for _, v := range RomanNums {
		if v == str {
			b = true
		}
	}
	return b
}

func GetResultOfNums(a string, b string, symb string) int {
	x, _ := strconv.Atoi(a)
	y, _ := strconv.Atoi(b)
	var result int
	switch symb {
	case "+":
		result = x + y
	case "-":
		result = x - y
	case "*":
		result = x * y
	case "/":
		result = x / y
	}
	return result
}

func RomanToArabic(s string, RomanNum []string) string {
	var res string
	for i := 0; i < len(RomanNum); i++ {
		if RomanNum[i] == s {
			res = strconv.Itoa(i)
		}
	}
	return res
}

func GetCharOfRomanNum(a int, RomanNum []string) string {
	return RomanNum[a]
}

func Calc(a string, b string, symbol string, RomanNum []string) {
	var resultRomanStr string
	var x, y int
	x, _ = strconv.Atoi(a)
	y, _ = strconv.Atoi(b)
	if IsArabicNum(a) && IsArabicNum(b) && x <= 10 && x >= 1 && y <= 10 && y >= 1 {
		result := GetResultOfNums(a, b, symbol)
		println(result)
	} else if IsRomanNum(a, RomanNum) && IsRomanNum(b, RomanNum) {
		result := GetResultOfNums(RomanToArabic(a, RomanNum), RomanToArabic(b, RomanNum), symbol)
		if result <= 0 {
			panic(errors.New("В римской системе нет отрицательных чисел и нуля"))
		}
		keys := []int{100, 50, 40, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
		MyMap := map[int]string{100: "C", 50: "L", 40: "XL", 10: "X", 9: "IX", 8: "VIII", 7: "VII", 6: "VI", 5: "V", 4: "IV", 3: "III", 2: "II", 1: "I"}
		for result > 0 {
			for i := 0; i < len(keys); i++ {
				x := result / keys[i]
				if x > 0 {
					resultRomanStr += MyMap[keys[i]]
					result -= keys[i]
					break
				}
			}
		}
		fmt.Println(resultRomanStr)
	} else {
		panic(errors.New("Введите два одинаковых числа от 1 до 10 в арабской(1,2,3,4,5...10) или римской (I,II,III,IV,V...X) системе счисления"))
	}
}

func main() {
	RomanNum := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	var inputstring, symbol, a, b string
	str := bufio.NewScanner(os.Stdin)
	str.Scan()
	inputstring = str.Text() //inputstring, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	inputstring = strings.TrimSpace(inputstring)
	inputstring = strings.ReplaceAll(inputstring, " ", "")
	var words, operation = SplitStringAndGetSymbol(inputstring)
	if len(words) > 2 || len(words) < 2 {
		panic(errors.New("Введённая строка не является выражением. Введите два римских (I + I) или арабских числа (1 + 1) от 1 до 10 через пробел с операторами +,-,*,/"))
	}
	a = words[0]
	symbol = operation
	b = words[1]
	Calc(a, b, symbol, RomanNum)
}
