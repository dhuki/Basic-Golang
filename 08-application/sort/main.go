package main

import (
	"fmt"
	"sort"
)

type person struct {
	First string
	Age   int
}

type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByName []person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].First < a[j].First }

func main() {
	xi := []int{4, 5, 2, 10, 1, 3, 6}
	xs := []string{"Dr.", "Saya", "Mau", "Tanya", "Apakah"}

	fmt.Println("sort primitive type")
	// sort int ascending
	fmt.Println(xi)
	sort.Ints(xi) // slice is on top of array it's pointer to make change without return value
	fmt.Println(xi)

	// sort string ascending
	fmt.Println(xs)
	sort.Strings(xs) // slice is on top of array it's pointer to make change without return value
	fmt.Println(xs)
	fmt.Println()

	// custom sort struct
	fmt.Println("custom sort in struct by age")
	people := []person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
	fmt.Println()
	fmt.Println("custom sort in struct by name")
	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)
}
