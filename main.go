package main

import (
	"UserManagementProject/questions/first_question"
	"UserManagementProject/questions/second_question"
	"UserManagementProject/questions/third_question"
	"fmt"
)

func main() {
	// Question 1
	fmt.Println("Question 1:")
	words := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}
	fmt.Println(first_question.SortWordsByA(words))

	// Question 2
	fmt.Println("\nQuestion 2:")
	second_question.GenerateOutput(9)

	// Question 3: Most repeated data in an array
	fmt.Println("\nQuestion 3:")
	data := []string{"apple", "pie", "apple", "red", "red", "red"}
	mostRepeated := third_question.MostRepeated(data)
	fmt.Println("Most Repeated:", mostRepeated)
}
