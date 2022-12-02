package main



/*
*****************************************************************

	@author: Adenugba Adeoluwa
	Date: 2nd  December 2022
	Purpose: To print all possible combinations of a string 
********************************************************************
*/


import (
	"fmt"
	"os"

)

var dictionary_values string = ""

func main1() {

	list := map[byte]string{
		'2': "ABC",
		'3': "DEF",
		'4': "GHI",
		'5': "JKL",
		'6': "MNO",
		'7': "PRS",
		'8': "TUV",
		'9': "WXY",
	}
	var number string
	fmt.Print("Input number(7 digits[2-9]): ")
	fmt.Scan(&number)
	for i := 0; i < len(number); i++ {

		if number[i] == '0' || number[i] == '1' {
			fmt.Println("Contains either 0 or 1")
			os.Exit(1)
		}
	}
	if len(number) == 0 {
		os.Exit(2)
	}
	for _, value := range number {
		conjoin_to_string(list[byte(value)])
	}
	result := permutations_of_string(dictionary_values)
	fmt.Println(result)

}

func conjoin_to_string(value string) {
	for i, _ := range value {
		dictionary_values += string(value[i])
	}
}

func permutations_of_string(testStr string) []string {

	var n func(testStr []rune, p []string) []string
	n = func(testStr []rune, p []string) []string {
		if len(testStr) == 0 {
			return p
		} else {
			result := []string{}
			for _, e := range p {
				result = append(result, join([]rune(e), testStr[0])...)
			}
			return n(testStr[1:], result)
		}
	}

	output := []rune(testStr)
	return n(output[1:], []string{string(output[0])})
}

func join(ins []rune, c rune) (result []string) {
	for i := 0; i <= len(ins); i++ {
		result = append(result, string(ins[:i])+string(c)+string(ins[i:]))
	}
	return
}
