package main

import (
	"fmt"
)

func main() {
	var buff []int
	buf1 := []int{}
	bif2 := []int{42}
	buf3 := make([]int, 0)
	buf4 := make([]int, 5)
	buf5 := make([]int, 5, 10)

	fmt.Println(buff, buf1, bif2, buf3, buf4, buf5)

	buff = append(buff, 9, 10)
	buff = append(buff, 13)

	fmt.Println(buff)

	var user map[string]string = map[string]string{
		"name":    "Sanzhar",
		"surname": "Anarbay",
	}

	if user["name"] == "Sanzhar" {
		fmt.Println("True")
	}

	mapValue := map[string]string{"name": "Sanzhar"}

	if keyValue, keyExist := mapValue["name"]; keyExist {
		fmt.Println(keyValue)
	}

	if _, keyExist := mapValue["name"]; keyExist {
		fmt.Println("It works!")
	}

	number := 1

	if number == 1 {
		fmt.Println("True ", number)
	} else if number == 3 {
		fmt.Println("false")
	}

	check := "hello"
	switch check {
	case "hee":
		fmt.Println("hee")
	case "hello":
		fmt.Println("hello")
	default:
		fmt.Println("Default")
	}

	isRun := true
	for isRun {
		fmt.Println("Loop just one condition")
		isRun = false
	}

	for indx := range buff {
		fmt.Println("range by index", indx)
	}

	for indx, val := range buff {
		fmt.Println("range by index", indx, val)
	}

	for key := range user {
		fmt.Println("Range map by key", key)
	}

	for key, value := range user {
		fmt.Println("Range map by key and value here key", key)
		fmt.Println("Range map by key and value here value", value)
	}

	for _, value := range user {
		fmt.Println("Range map by value", value)
	}
}
