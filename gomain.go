package main

import (
	"fmt"
	"strconv"
	"strings"
)

var err error

/*
isCorrect() --- check is correct input data from user or not
Correct means data in format "number1symbolnumber2"
input:
	str --- input data from
return:
	true if it correct and false if not
*/

func IsCorrect(str string) bool {
	var length int = len(str)
	var symbols_count int = 0

	for i := 0; i < length; i++ {
		if string([]rune(str)[i]) == "+" || string([]rune(str)[i]) == "-" {
			symbols_count++
			if symbols_count > 1 || i == 0 || i == length {
				return false
			}
		} else {
			_, err = strconv.Atoi(string([]rune(str)[i]))

			if err != nil {
				return false
			}
		}
	}
	if symbols_count == 0 {
		return false
	}
	return true
}

/*
getNumbers() --- split data from user into 3 strintgs and make number one length
				by adding "0"
input:
	str --- input data from
return:
	number1, number2, + or -
*/

func GetNumbers(str string) (string, string, string) {
	var num []string
	var symbol string
	var count_to_add int = 0

	if strings.Contains(str, "+") {
		symbol = "+"
	} else {
		symbol = "-"
	}

	num = strings.Split(str, symbol)



	for len(num[0]) > len(num[1]) {
		if len(num[0]) % 4 != 0{
			count_to_add = 4 - len(num[0]) % 4
			for i := 0; i < count_to_add; i++ {
				num[0] = "0" + num[0]
			}
		} 
	}

	for len(num[0]) < len(num[1]) {
		if len(num[1]) % 4 != 0{
			count_to_add = 4 - len(num[1]) % 4
			for i := 0; i < count_to_add; i++ {
				num[1] = "0" + num[1]
			}
		} 
		num[0] = "0" + num[0]
	}

	return num[0], num[1], symbol
}

/*
converToArray() --- convert string to array of int
input:
	str --- string to convert
return:
	arr --- array of int
*/

func ConverToArray(str string, c chan []int) <- chan []int {
	var arr []int
	var element int

	for i := 0; i < len(str); i++ {
		element, err = strconv.Atoi(string([]rune(str)[i]))
		arr = append(arr, element)
	}
	
	c <- arr
	return c
}

/*
converToString() --- convert array of int to string
input:
	arr --- array of int to convert
return:
	str --- string
*/

func ConverToString(arr []int) string {
	var str string
	var i int = 0

	for arr[i] == 0 && i < len(arr)-1 {
		i++
	}
	for i < len(arr) {
		str += strconv.Itoa(arr[i])
		i++
	}

	return str
}

/*
add() --- add numbers
input:
	numb1,numb2 -- numbers to add
return:
	numb1 -- sum
*/
func Add(numb1, numb2 []int, c chan []string) string {

	var length int = len(numb1)

	for i := length - 1; i > -1; i-- {
		numb1[i] += numb2[i]

		if numb1[i] > 9 && i != 0 {
			numb1[i] -= 10
			numb1[i-1]++
		}
	}
	c <- ConverToString(numb1)
	return c
}

func Add_str(numb1, numb2 []int) string {

	var length int = len(numb1)

	for i := length - 1; i > -1; i-- {
		numb1[i] += numb2[i]

		if numb1[i] > 9 && i != 0 {
			numb1[i] -= 10
			numb1[i-1]++
		}
	}

	return ConverToString(numb1)
}
/*
isNegative() --- check can be result of subtraction negative
input:
	numb1,numb2 --- numbers to subtraction
return:
	true if result will be negative and false if not
*/

func IsNegative(numb1, numb2 []int) bool {

	var i = 0

	for numb1[i] == numb2[i] && i != 0 {
		i++
	}
	if i != 0 {
		i++
	}

	if numb1[i] >= numb2[i] {
		return false
	} else {
		return true
	}
}

/*
sub() --- substract numbers
input:
	numb1,numb2 -- numbers to add
return:
	numb1 -- result of subtraction
*/
func Sub(numb1, numb2 []int) string {

	var length int = len(numb1)
	var prefix string = ""

	if IsNegative(numb1, numb2) {
		tmp := numb1
		numb1 = numb2
		numb2 = tmp
		prefix = "-"
	}

	for i := length - 1; i > -1; i-- {
		numb1[i] -= numb2[i]

		if numb1[i] < 0 && i != 0 {
			numb1[i] += 10
			numb1[i-1]--
		}
	}

	return prefix + ConverToString(numb1)
}

func split4(number1 string){
	var result []string 

	for i := 0; i < len(number1); i += 4 {
		fmt.Println("in")
		result = append(result,number1[i:i+4])
	}

	fmt.Println(result)
}

func goAdd(number1 []string, number2 []string) []string{
	var	ch chan []string 
	var splt []string
	

	for i := 0; i < len(number1); i++ {
		go Add(number1[i], number2[i], ch)
		number1 <- ch  		
	}

	for i := 1; i < len(number1); i++ {
		if len(number1[i]) > 4 {
			splt = number1[i][len(number1[i])-4:len(number1[i])]
			number1[i-1] = Add_str(number1[i-1], splt)
		}
	}
	for i := len(number1); i > 0; i-- {
		
	}

}

func main() {
	// var input, num1, num2, symbol, result string

	// fmt.Println("Enter expression: ")
	// fmt.Scanln(&input)
	// c := make(chan []int)

	// if IsCorrect(input) {
	// 	num1, num2, symbol = GetNumbers(input)
	// 	go ConverToArray(num1, c)
	// 	go ConverToArray(num2, c)
	// 	number1, number2 := <-c, <-c
	// 	if symbol == "+" {
	// 		result = Add(number1, number2)
	// 	} else {
	// 		result = Sub(number1, number2)
	// 	}
	// 	fmt.Println("Result\n" + result)
	// } else {
	// 	fmt.Println("Expression is not correct.")
	// }
	goAdd("1234+12345")

}
