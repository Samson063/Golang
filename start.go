package main

import (
	"fmt"
)

var church, address, parish, area string = "Rccg", "Alapere", "Pec  uliar", "Four"
var profile, profileName, profileAge, profileSchool = 1, "Anthony Mercy", 23, "Nazareth"
var (
	profile2       int    = 2
	profile2Name   string = "Anthony Glory"
	profile2Age    int    = 18
	profile2School string = "FUNAAB"
)

func main() {
	var name string = "Anthony Samson"
	var name2 = "Anthony Glory"
	name3 := "Anthony Samson"
	var school string

	//Typed constant
	const MYNAME string = "Anthony Samson"
	const SCHOOLFEES int = 100000

	const (
		company  string = "NUPRC"
		branches int    = 7
		point           = 3.4
	)

	fmt.Println("Hello World!")

	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(name3)
	fmt.Println(school)

	fmt.Println(company)
	fmt.Println(branches)
	fmt.Println(point)
	fmt.Println(area)

	fmt.Println(profile)
	fmt.Println(profileName)
	fmt.Println(profileAge)
	fmt.Println(profileSchool)

	fmt.Println(profile2)
	fmt.Println(profile2Name)
	fmt.Println(profile2Age)
	fmt.Println(profile2School)

	fmt.Println(MYNAME)
	fmt.Println(SCHOOLFEES)

	var new, old = "Dell latitude", "Hp pavilion"
	var num1, num2 = 20, 80

	fmt.Print(new, old, "\n")
	fmt.Println(new, old)
	fmt.Print(num1, num2)
	fmt.Println(num1, num2)

	//Arrays

	//arr1 := [5]int{}              //not initialized
	arr2 := [5]int{1, 2}          //partially initialized
	arr3 := [5]int{1, 2, 3, 4, 5} //fully initialized

	//fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(len(arr2))

	//Slice

	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)

	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr1[2:4]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))

	//creating sclice using the make() function
	mmyslice1 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", mmyslice1)
	fmt.Printf("length = %d\n", len(mmyslice1))
	fmt.Printf("capacity = %d\n", cap(mmyslice1))

	// with omitted capacity
	mmyslice2 := make([]int, 5)
	fmt.Printf("myslice2 = %v\n", mmyslice2)
	fmt.Printf("length = %d\n", len(mmyslice2))
	fmt.Printf("capacity = %d\n", cap(mmyslice2))

}
