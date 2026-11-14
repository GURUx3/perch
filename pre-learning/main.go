package main

import (
	"fmt"
)

func play(status bool) {

	age := 21

	if age >= 18 {
		fmt.Println("you can drink beer !")
	} else {
		fmt.Println("No no not now !")
	}

	isAdmin := status

	if isAdmin {
		fmt.Println("Access Granted! : Admin")
	} else {
		panic("Unauthorized User Found !")
	}

	// slice IMPORTANT CONCEPT IN GO

	// var s1 []int

	s2 := []string{"apple", "banan", "goa", "orange", "grape", "mango", "plum", "cherry"}

	for _, val := range s2 {
		fmt.Println(val)
	}

	mySlice := make([]int, 3, 5)
	fmt.Println(len(mySlice), cap(mySlice), mySlice)

	mySlice[0] = 10
	fmt.Println(len(mySlice), cap(mySlice), mySlice)
	mySlice[1] = 20

	fmt.Println(len(mySlice), cap(mySlice), mySlice)
	mySlice[2] = 30

	fmt.Println(len(mySlice), cap(mySlice), mySlice)
	mySlice = append(mySlice, 40)

	fmt.Println(len(mySlice), cap(mySlice), mySlice)
	mySlice = append(mySlice, 40)

	fmt.Println(len(mySlice), cap(mySlice), mySlice)

	mySlice = append(mySlice, 60, 70, 80)
	fmt.Println(len(mySlice), cap(mySlice), mySlice)

}

func mp() {

	// maps

	docs := make(map[int]string)

	docs[1] = "Guru"
	docs[2] = "22"
	docs[3] = "guru@gmail.com"

	// fmt.Println(docs)

	for id, val := range docs {
		fmt.Println(id, " : ", val)
	}

	json_data := map[string]string{"foo": "foo1", "foo2": "foo2"}

	fmt.Println(json_data)

}

func fat(n int) int {

	return n * n
}

func factorial_go(n int) int {

	if n < 0 {
		fmt.Errorf("Factorial is not define in neg numbers ")
	}

	if n == 0 || n == 1 {
		return 1
	}

	return n * factorial_go(n-1)
}

func fib(n int) int {

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)

}

func fib_with_count(count, a, b int) {
	if count <= 0 {
		return
	}

	fmt.Println(a, " ")
	fib_with_count(count-1, b, a+b)
}

func main() {
	fmt.Println("Hello World")

	res := fat(12)
	fmt.Println(res)

	fato := factorial_go(5)

	fmt.Println(fato)

	fmt.Println(fib(10))
	fib_with_count(5, 0, 1)

	
	
}
