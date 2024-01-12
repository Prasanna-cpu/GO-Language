package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println("Hello Go")
	fmt.Println("Same Same but different")
	fmt.Println("No Money No Honey")
	fmt.Println("Good boy goes to heaven")
	fmt.Println("Bad boy goes to Bangkok")

	var name = "Leo Das"
	fmt.Println(name)
	age := 45
	fmt.Println(age)

	fmt.Println(math.Abs(-9))

	var fruit string = "Apple"
	fmt.Println(fruit)
	var one, two, three int = 1, 2, 3
	fmt.Println(one)
	fmt.Println(two)
	fmt.Println(three)

	const Unchanged int = 6
	fmt.Println(Unchanged)

	var isStudent bool = true
	fmt.Println(isStudent)
	isTrue := false
	fmt.Println(isTrue)

	var num1 float32 = 20.00
	fmt.Println(num1)

	language := "Go Language"

	fmt.Println(language)

	name1 := "KumarKPK"
	name2 := "KumarK"

	fmt.Printf("%c", name1[2])
	fmt.Println("")
	fmt.Printf("%c", name2[1])
	fmt.Println("")

	fmt.Println(len(name1))
	fmt.Println(strings.Compare(name1, name2))

	findChar := "Hey this is a good language"
	fmt.Println(strings.Contains(findChar, "this"))

	password := "1234GetOnTheDanceFloor"

	if len(password) > 40 {
		fmt.Println("Password is valid")
	} else if password == "" {
		fmt.Println("Please Provide a Password")
	} else {
		fmt.Println("Invalid Password")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for j := 5; j > 0; j-- {
		fmt.Println(j)
	}

	for j := 1; j <= 10; j++ {
		if j%2 == 0 {
			continue
		}
		fmt.Println(j)
	}

	// for i := 0; i < 20; i++ {
	// 	fmt.Println("___OUTER___", i)
	// 	for j := 0; j < 20; j++ {
	// 		fmt.Println("___INNER___", j)
	// 	}
	// }

	number := 0

	for number < 10 {
		fmt.Println(number)
		number++
	}

	num := 5

	switch num {
	case 1:
		fmt.Println("One")

	case 2:
		fmt.Println("Two")

	case 3:
		fmt.Println("Three")
	default:
		fmt.Printf("Invalid")

	}

	var arr1 = [3]int{1, 2, 3}
	var arr2 = [5]int{1, 2, 3, 4, 5}

	fmt.Println(arr1)
	fmt.Println(arr2)

	myslice := []int{}
	fmt.Println(len(myslice))
	fmt.Println(cap((myslice)))
	fmt.Println(myslice)

	myslice2 := []string{"Go", "Java", "Python", "C#"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))

}
