package main

import "fmt"

type Person struct {
	Id int
	Name string
}

func (p *Person) SetName(name string){
	p.Name = name
}

type Account struct{
	Id int
	Name string
	Person
}

type MySlice []int

func (mSl *MySlice) Add(val int)  {
	*mSl = append(*mSl, val)
}

func (mSl *MySlice) Count() int  {
	return len(*mSl)
}

func main() {

	pers:= Person{1, "Sanzhar"}
	pers.SetName("Sanzh")

	var account Account = Account{
		Id: 1,
		Name: "Dl",
		Person: Person{
			Id: 2,
			Name: "Sanzhar Anarbay",
		},
	}

	account.SetName("Ivan Ivanov")

	fmt.Println(account)

	mSl := MySlice([]int{1,2})

	mSl.Add(3)

	fmt.Println(mSl)
	fmt.Println(mSl.Count())

	var buff []int
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
