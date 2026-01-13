package main

import (
	"fmt"
)

func main() {
	groupAnagrams([]string{"ant", "nat", "apple", "toy", "yot"})
}

type Address struct {
	firstLineAddress string
	city             string
	postcode         string
	udprn            int32
}

type User struct {
	username string
	age      int32
	address  Address
}

func groupAnagrams(words []string) [][]string {

	user := User{
		username: "Michael",
		age:      32,
		address: Address{
			firstLineAddress: "19 Passmore Gardens",
			city:             "London",
			postcode:         "N11 2PE",
			udprn:            123456,
		},
	}

	fmt.Println(user)

	return nil
}
