package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}
func mul(x, y int) int {
	return x * y
}
func div(x, y int) int {
	return x / y
}
func ParserInt(x, y string) (int, int) {
	var num1, num2 int
	num1, _ = strconv.Atoi(x)
	num2, _ = strconv.Atoi(y)
	return num1, num2
}
func RomanToInt(s string) int {
	var v, lv, sb int
	h := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	for i := len(s) - 1; i >= 0; i-- {
		sb = h[s[i]]
		if sb < lv {
			v -= sb
		} else {
			v += sb
		}
		lv = sb
	}
	return v
}
func IntToRoman(number int) string {
	conv := []struct {
		vaule int
		diget string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	roman := ""
	for _, conv := range conv {
		for number >= conv.vaule {
			roman += conv.diget
			number -= conv.vaule
		}
	}
	return roman
}
func ChekValid(x, oper, y string) {
	inted := [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	roman := [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for i := 0; i < len(inted); i++ {
		for j := 0; j < len(roman); j++ {
			if x == inted[i] && y == inted[j] {

				n1, n2 := ParserInt(x, y)
				switch oper {
				case "+":
					fmt.Println(add(n1, n2))
				case "-":
					fmt.Println(sub(n1, n2))
				case "/":
					fmt.Println(div(n1, n2))
				case "*":
					fmt.Println(mul(n1, n2))
				default:
					fmt.Println("Не существующая мат операция")
				}

			} else if x == roman[i] && y == roman[j] {
				n1 := RomanToInt(x)
				n2 := RomanToInt(y)
				switch oper {
				case "+":
					roman := add(n1, n2)
					fmt.Println(IntToRoman(roman))
				case "-":
					roman := sub(n1, n2)
					if roman < 1 {
						fmt.Println("В римской системе нет нуля и отрицательных чисел")
					}
					fmt.Println(IntToRoman(roman))
				case "/":
					roman := div(n1, n2)
					fmt.Println(IntToRoman(roman))
				case "*":
					roman := mul(n1, n2)
					fmt.Println(IntToRoman(roman))
				default:
					fmt.Println("Не существующая мат операция")
				}
			}
			if (x == roman[i] && y == inted[j]) || (x == inted[i] && y == roman[j]) {
				fmt.Println("Разные типы счисления")
			}
		}
	}
}

func main() {
	var x, oper, y string
	var lin string
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	lin = sc.Text()
	arr := strings.Split(lin, " ")
	if len(arr) != 3 {
		fmt.Println("Математическая операция должна состоять из двух элементов")
		return
	}
	x, oper, y = arr[0], arr[1], arr[2]
	/* x1, y1 := ParserInt(x, y)
	x2 := RomanToInt(x)
	y2 := RomanToInt(y)
	if x1 > 10 || y1 > 10 || x1 < 0 || y1 < 0 {
		fmt.Println("Числа должны находиться в диапазоне от 0 до 10")
	}
	if x2 > 10 || y2 > 10 || x2 <= 0 || y2 <= 0 {
		fmt.Println("Числа должны находиться в диапазоне от 1 до 10")
	}

	*/ChekValid(x, oper, y)
}