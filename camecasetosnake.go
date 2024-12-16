package main

import "fmt"

//Write a function that converts a string from camelCase to snake_case.

// If the string is empty, return an empty string.
//If the string is not camelCase, return the string unchanged.
//If the string is camelCase, return the snake_case version of the string.
//For this exercise you need to know that camelCase has two different writing alternatives that will be accepted:

//Rules for writing in camelCase:

//The word does not end on a capitalized letter (CamelCasE).
//No two capitalized letters shall follow directly each other (CamelCAse).
//Numbers or punctuation are not allowed in the word anywhere (camelCase1).

func CamelToSnakeCase(s string) string {
	
	if len(s) == 0 {
		return s
	}	
	
	isCap := false
	for _, char := range s {
		if !(char >= 'a' && char <= 'z') && !(char >= 'A' && char <= 'Z'){
			return s
		}
		if char >= 'A' && char <= 'Z'{
			
			if isCap {
				return s
			}
			isCap = true
		} else {
			isCap = false
		}
	}

	lastChar := s[len(s)-1]
	if lastChar >= 'A' && lastChar <= 'Z'  {
		return s
	}
	result := ""
	for i := 0; i < len(s); i++ {
	if i != 0 && s[i] >= 'A' && s[i] <= 'Z' {
		result += "_"
	}
	result += string(s[i])

}
return result
}





func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("AbC"))
	fmt.Println(CamelToSnakeCase("hey2"))
	
}
