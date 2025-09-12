package main

import (
	"fmt"
)

var church, address, parish, area string = "Rccg", "Alapere", "Peculiar", "Four"
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

	const MYNAME string = "Anthony Samson"
	const SCHOOLFEES int = 100000

	fmt.Println("Hello World!")

	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(name3)
	fmt.Println(school)

	fmt.Println(church)
	fmt.Println(address)
	fmt.Println(parish)
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
}
