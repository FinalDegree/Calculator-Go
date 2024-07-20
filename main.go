package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func SplitStringAndGetSymbol(x string) ([]string, string) {
	var operation string
	var array []string
	if strings.Contains(x, "+") || strings.Contains(x, "-") || strings.Contains(x, "*") || strings.Contains(x, "/") {
		if strings.Contains(x, "+") {
			array = strings.Split(x, "+")
			operation = "+"
		} else if strings.Contains(x, "-") {
			array = strings.Split(x, "-")
			operation = "-"
		} else if strings.Contains(x, "*") {
			array = strings.Split(x, "*")
			operation = "*"
		} else if strings.Contains(x, "/") {
			array = strings.Split(x, "/")
			operation = "/"
		}
	} else {
		panic("Введённая строка не является математическим выражением.")
	}
	return array, operation
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

func Calc(a string, b string, symbol string, RomanNum []string) {
	var resultRomanStr string

	x, errorA := strconv.Atoi(a)
	y, errorB := strconv.Atoi(b)
	if errorA == nil && errorB == nil && x <= 10 && x >= 1 && y <= 10 && y >= 1 {
		result := GetResultOfNums(a, b, symbol)
		println(result)
	} else if slices.Contains(RomanNum, a) && slices.Contains(RomanNum, b) {
		result := GetResultOfNums(RomanToArabic(a, RomanNum), RomanToArabic(b, RomanNum), symbol)
		if result <= 0 {
			panic("В римской системе нет отрицательных чисел и нуля")
		}
		keys := []int{100, 90, 50, 40, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
		RomanMap := map[int]string{100: "C", 90: "XC", 50: "L", 40: "XL", 10: "X", 9: "IX", 8: "VIII", 7: "VII", 6: "VI", 5: "V", 4: "IV", 3: "III", 2: "II", 1: "I"}
		for result > 0 {
			for i := 0; i < len(keys); i++ {
				x := result / keys[i]
				if x > 0 {
					resultRomanStr += RomanMap[keys[i]]
					result -= keys[i]
					break
				}
			}
		}
		fmt.Println(resultRomanStr)
	} else {
		panic("Введите два одинаковых числа от 1 до 10 в арабской(1,2,3,4,5...10) или римской (I,II,III,IV,V...X) системе счисления")
	}
}

func main() {
	RomanNum := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	var inputstring, symbol, a, b string
	inputstring, _ = bufio.NewReader(os.Stdin).ReadString('\n') //str := bufio.NewScanner(os.Stdin); str.Scan(); inputstring = str.Text()
	inputstring = strings.TrimSpace(inputstring)
	inputstring = strings.ReplaceAll(inputstring, " ", "")
	var array, operation = SplitStringAndGetSymbol(inputstring)
	if len(array) > 2 || len(array) < 2 {
		panic("Введённая строка не является выражением. Введите два римских (I + I) или арабских числа (1 + 1) от 1 до 10 через пробел с операторами +,-,*,/")
	}
	a = array[0]
	symbol = operation
	b = array[1]
	Calc(a, b, symbol, RomanNum)
}
