package main

import (
	"fmt"
)

func reverse(s []int) { 
	for i , j := 0, len(s)-1; i < j; i, j = i + 1, j-1 {
		s[i], s[j] = s[j],s[i]
	}
}

func main() {
	months := [...]string{
		1: "January",
		2: "February",
		3: "March",
		4: "April",
		5: "May",
		6: "June",
		7: "July",
		8: "August",
		9: "September",
		10: "October",
		11: "November",
		12: "December",
	}
	
	fmt.Println(months[1:13])
	fmt.Println(months[1:])
	fmt.Println(months[:12])
	fmt.Println(months[:])
	fmt.Println(months[2:3])

	q2 := months[4:7]
	summer := months[6:9]

	fmt.Println(q2)
	fmt.Println(summer)

	fmt.Println(months[8:11])

	numbers := [...]int{1,2,3}

	reverse(numbers[:])
	fmt.Println(numbers)
}