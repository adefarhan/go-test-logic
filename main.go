package main

import "fmt"

func main() {

	// {4, 73, 67, 38, 33}
	// {3, 84, 29, 57}
	// {1, 100}
	// {1, 0}
	// {60, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, }

	students := []int{60, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, }

    result := gradingStudents(students)

	fmt.Println(result)
}


func gradingStudents(grades[]int) []int  {
	result := []int{}

	for i := 1; i < grades[0] + 1; i++ {
		num := grades[i]
		closestTimesFive := num + (5 - (num % 5))
		calculate := closestTimesFive - num

		if num >= 38 {
			if calculate < 3 {
				result = append(result, closestTimesFive)
			}else {
				result = append(result, num)
			}
		}else {
			result = append(result, num)
		}
	}

	return result
}