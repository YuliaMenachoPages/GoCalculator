package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var intToRomanMap = map[int]string{
	1:   "I",
	2:   "II",
	3:   "III",
	4:   "IV",
	5:   "V",
	6:   "VI",
	7:   "VII",
	8:   "VIII",
	9:   "IX",
	10:  "X",
	11:  "XI",
	12:  "XII",
	13:  "XIII",
	14:  "XIV",
	15:  "XV",
	16:  "XVI",
	17:  "XVII",
	18:  "XVIII",
	19:  "XIX",
	20:  "XX",
	21:  "XXI",
	22:  "XXII",
	23:  "XXIII",
	24:  "XXIV",
	25:  "XXV",
	26:  "XXVI",
	27:  "XXVII",
	28:  "XXVIII",
	29:  "XXIX",
	30:  "XXX",
	31:  "XXXI",
	32:  "XXXII",
	33:  "XXXIII",
	34:  "XXXIV",
	35:  "XXXV",
	36:  "XXXVI",
	37:  "XXXVII",
	38:  "XXXVIII",
	39:  "XXXIX",
	40:  "XL",
	41:  "XLI",
	42:  "XLII",
	43:  "XLIII",
	44:  "XLIV",
	45:  "XLV",
	46:  "XLVI",
	47:  "XLVII",
	48:  "XLVIII",
	49:  "XLIX",
	50:  "L",
	51:  "LI",
	52:  "LII",
	53:  "LIII",
	54:  "LIV",
	55:  "LV",
	56:  "LVI",
	57:  "LVII",
	58:  "LVIII",
	59:  "LIX",
	60:  "LX",
	61:  "LXI",
	62:  "LXII",
	63:  "LXIII",
	64:  "LXIV",
	65:  "LXV",
	66:  "LXVI",
	67:  "LXVII",
	68:  "LXVIII",
	69:  "LXIX",
	70:  "LXX",
	71:  "LXXI",
	72:  "LXXII",
	73:  "LXXIII",
	74:  "LXXIV",
	75:  "LXXV",
	76:  "LXXVI",
	77:  "LXXVII",
	78:  "LXXVIII",
	79:  "LXXIX",
	80:  "LXXX",
	81:  "LXXXI",
	82:  "LXXXII",
	83:  "LXXXIII",
	84:  "LXXXIV",
	85:  "LXXXV",
	86:  "LXXXVI",
	87:  "LXXXVII",
	88:  "LXXXVIII",
	89:  "LXXXIX",
	90:  "XC",
	91:  "XCI",
	92:  "XCII",
	93:  "XCIII",
	94:  "XCIV",
	95:  "XCV",
	96:  "XCVI",
	97:  "XCVII",
	98:  "XCVIII",
	99:  "XCIX",
	100: "C",
}

var romanToIntMap = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

var arabicToIntMap = map[string]int{
	"1":  1,
	"2":  2,
	"3":  3,
	"4":  4,
	"5":  5,
	"6":  6,
	"7":  7,
	"8":  8,
	"9":  9,
	"10": 10,
}

var funcByOperatorMap = map[string]func(a, b int) int{
	"+": addition,
	"-": subtraction,
	"/": division,
	"*": multiplication,
}

func main() {
	fmt.Print("Введите выражение в формате 'число операнд число': ")

	exp, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic(err)
	}

	exp = strings.Trim(exp, "\n")

	splittedExp := strings.Split(exp, " ")
	if len(splittedExp) != 3 {
		panic("выпажение должно быть введено в формате 'число операнд число'")
	}

	n1Str, operator, n2Str := splittedExp[0], splittedExp[1], splittedExp[2]

	operatorFunc, ok := funcByOperatorMap[operator]
	if !ok {
		panic("невалидный оператор")
	}

	var n1, n2 int

	if n1, ok = arabicToIntMap[n1Str]; ok {
		if n2, ok = arabicToIntMap[n2Str]; ok {
			fmt.Println(operatorFunc(n1, n2))
		} else {
			panic("разный формат записи чисел(арабские/римские)")
		}
	} else if n1, ok = romanToIntMap[n1Str]; ok {
		if n2, ok = romanToIntMap[n2Str]; ok {
			res := operatorFunc(n1, n2)
			if res < 1 {
				panic("результат работы калькулятора с римскими числами < 1")
			}
			fmt.Println(intToRomanMap[res])
		} else {
			panic("разный формат записи чисел(арабские/римские)")
		}
	} else {
		panic("невалидное число")
	}

}

func addition(a, b int) int {
	return a + b
}

func subtraction(a, b int) int {
	return a - b
}

func division(a, b int) int {
	return a / b
}

func multiplication(a, b int) int {
	return a * b
}
