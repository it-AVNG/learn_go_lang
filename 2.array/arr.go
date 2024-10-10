package main

import "fmt"

func main() {
	intArry := [3]int{2, 3, 4}
	fmt.Println("declairing arrray with name = [length]datatype{value1,value2,..}: ", intArry)
	infArry := [...]int{4, 5, 6}
	fmt.Println("inferred the length by = [...]datatype{value1,value2,..}: ", infArry)
	strArr := [3]string{}
	fmt.Println("initiated with nul value = [3]string{}: ", strArr)
	intArr := [3]int{}
	fmt.Println("initiated with nul value = [3]int{}: ", intArr)
	partStrArr := [3]string{"hi"}
	fmt.Println(" not initiated value will be nul: ", partStrArr)
	partIntArr := [3]int{1}
	fmt.Println("not initiated value will be null = [3]int{1}: ", partIntArr)
	// change element
	partIntArr[1] = 4
	fmt.Println("access array value with '=' operand: ", partIntArr)

	indexArr := [5]int{0: 23, 4: 212, 3: 32}
	fmt.Println("init and assign specific value: ", indexArr)
	fmt.Println("length of array with len(): ", len(indexArr))

	// Slices
	arr1 := [...]int{1, 2, 4, 6, 8}
	fmt.Println("an arr1 is initiated: ", arr1)
	slice1 := []int{45, 62, 7}
	fmt.Println("slice1 init with []int{values}: ", slice1)
	slice2 := arr1[1 : len(arr1)-1]
	fmt.Println("slice2 init with slicing an Array: ", slice2)
	slice3 := make([]int, 4, 5)
	fmt.Println("slice3 init with make([]int, 4, 5): ", slice3)
	fmt.Println("index an out-of-bound slice3[4] is impossible even if the capacity is there")
  slice3 = append(slice3, 9,2)
  fmt.Println("new slice3: ", slice3)
}
