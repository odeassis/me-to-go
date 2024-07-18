package main

import (
	"fmt"
)

func sockPairs(socks string) int {
	amount := 0

	m_socks := make(map[rune]int)

	for _, v := range socks {
		m_socks[v] += 1

		if m_socks[v]%2 == 0 {
			amount += 1
		}
	}

	return amount
}

func main() {

	fmt.Println(sockPairs("CABBACCC"))
}
