// composite data type and aggregate data type
// basically is datatype that can hold many of datas
// difference is compsite data can insert multiple datatype using struct
// and aggregate datatype can only insert one datatype

package main

import "fmt"

func main() {
	// defining slice
	// x := type{values} // composite literal
	fmt.Println("A. defining slice")
	x := []int{1, 2, 3}
	for index, value := range x {
		fmt.Println(index, value)
	}
	fmt.Println("===== End of A.")
	fmt.Println()

	// slicing a slice
	fmt.Println("B. slicing a slice")
	y := []int{1, 2, 3, 4}
	fmt.Println(y[1:]) // it's just like python
	fmt.Println(y[1:3])
	fmt.Println("===== End of B.")
	fmt.Println()

	// append a slice
	fmt.Println("C. append a slice")
	z := []int{}
	z = append(z, 1, 2, 3)
	fmt.Println(z)
	z1 := []int{4, 5, 6, 7}
	z = append(z, z1...) // this is variadic parameter, it will take all unlimited number of value and insert to slice
	fmt.Println(z)
	fmt.Println("===== End of C.")
	fmt.Println()

	// deleted a slice
	fmt.Println("D. deleted a slice")
	z = append(z[:1], z[2:]...) // value 2 is deleted
	fmt.Println(z)
	fmt.Println("===== End of D.")
	fmt.Println()

	// make keyword
	fmt.Println("E. make keyword")
	makeSomething := make([]int, 10, 11) // it parameter is : datatype, length, capacity
	fmt.Println("length", len(makeSomething))
	fmt.Println("capacity", cap(makeSomething))
	makeSomething[0] = 10
	makeSomething[9] = 11
	fmt.Println(makeSomething)
	fmt.Println()
	// makeSomething[10] = 12 // it will gettin error since we only had 10 length
	// but we can make this
	makeSomething = append(makeSomething, 12)
	fmt.Println("length", len(makeSomething))
	fmt.Println("capacity", cap(makeSomething))
	fmt.Println(makeSomething)
	fmt.Println()

	// if length is over the capacity it will create another centain capacity
	makeSomething = append(makeSomething, 13)
	fmt.Println("length", len(makeSomething))
	fmt.Println("capacity", cap(makeSomething))
	fmt.Println(makeSomething)
	fmt.Println("===== End of E.")
	fmt.Println()

	// map
	fmt.Println("F. map")
	m := map[string]int{
		"dhuki": 23,
		"anda":  50,
	}
	fmt.Println(m)
	fmt.Println(m["dhuki"])  // m["dhuki"] will return two value : value it self, true/false means exist or not exist
	fmt.Println(m["beliau"]) // so by default if key of map doesn't exist it will return zero value of types (int is 0)
	// so we should check it if we won't print out a value that key doesn't exist
	if v, ok := m["beliau"]; ok {
		fmt.Println(v) // it is idiomatic cheking in golang
	}

	// delete map key
	delete(m, "anda")
	fmt.Println(m)

	fmt.Println("===== End of F.")
	fmt.Println()
}
