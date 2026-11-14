package main

import (
	"fmt"
)



func play(status bool){


		age := 21


		if age >= 18 {
			fmt.Println("you can drink beer !")
		} else{
			fmt.Println("No no not now !")
		}




		isAdmin := status


		if isAdmin{
			fmt.Println("Access Granted! : Admin")
		} else{
			panic("Unauthorized User Found !")
		}


		


		// slice IMPORTANT CONCEPT IN GO


		// var s1 []int

		s2 := []string{"apple", "banan", "goa", "orange", "grape" , "mango", "plum" , "cherry"}

		for _ , val := range s2{
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


func main(){
	fmt.Println("Hello World")

	// admin := true
	// play(admin)

	

	
}