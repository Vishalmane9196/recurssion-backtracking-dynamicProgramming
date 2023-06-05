package main

import (
	"fmt"
	"strconv"
)

/*******Basic Problems*******/
func fibo(n int) int {

	//base case
	if n <= 1 {
		return n
	}
	//recurssive relation and processing
	ans := fibo(n-1) + fibo(n-2)

	return ans
}

func printDecreasingOrder(n int) {
	//base case
	if n < 1 {
		return
	}
	//processing
	fmt.Printf("%v ", n)
	//recurssive relation
	printDecreasingOrder(n - 1)
}

func printIncreasingOrder(n int) {
	//base case
	if n < 1 {
		return
	}
	//recurssive relation
	printDecreasingOrder(n - 1)
	//processing
	fmt.Printf("%v ", n)
}

func printDecreaingIncreasingOrder(n int) {
	//base case
	if n < 1 {
		return
	}
	//processing
	fmt.Printf("%v ", n)
	//recurssive relation
	printDecreaingIncreasingOrder(n - 1)
	//processing
	fmt.Printf("%v ", n)
}

func printIncreasingDecreasingOrder(curIdx, total int) {
	//base case
	if curIdx == total {
		return
	}
	//processing
	fmt.Printf("%v ", curIdx+1)
	//recurssive call
	printIncreasingDecreasingOrder(curIdx+1, total)
	//processing
	fmt.Printf("%v ", curIdx+1)
}

// At one point between 11 to 20 factorial will starts to give wrong answer and find out why is this behaviour
func factorial(n int) int {

	//base case
	if n <= 1 {
		return 1
	}
	//recurssive relation
	ans := n * factorial(n-1)
	//processing
	return ans
}

/*********Array Problem*********/
// print array
func printArray(arr []int, index, size int) {

	//base case
	if index == size {
		return
	}
	//processing
	fmt.Printf("%v ", arr[index])
	//recurssive call
	printArray(arr, index+1, size)
}

// print reverse array
func printArrayReverse(arr []int, index, size int) {

	//approach one
	//traverse till end of last element and while returning from last element print each elment

	//base case
	if index == len(arr) {
		return
	}

	//recurssive call
	printArrayReverse(arr, index+1, size)
	fmt.Printf("%v ", arr[index])
}

// print reverse array
func printArrayReverse1(arr []int, index int) {

	//approach two
	//traverse from last element to 1st element
	//base case
	if index < 0 {
		return
	}
	//recurssive call
	fmt.Printf("%v ", arr[index])
	printArrayReverse1(arr, index-1)
}

// not working need to figure it out
// print reverse array
func printArrayReverse2(arr *[]int, size int) {

	//approach thress
	//pass the array and size and print first element of that array till size become zero
	//base case
	if size == 0 {
		return
	}
	//recurssive call
	fmt.Println((*arr)[0])
	*arr = (*arr)[1:]
	printArrayReverse1(*arr, size-1)
}

// find max of array
func maxOfArray(arr []int, max *int, index, size int) {
	if index == size {
		return
	}
	if arr[index] > *max {
		*max = arr[index]
	}

	maxOfArray(arr, max, index+1, size)
}

// find the first occurance of target
func firstOccurance(arr []int, index, size, target int, ans *int) {
	if index == size {
		return
	}

	if arr[index] == target {
		*ans = index
		return
	}
	firstOccurance(arr, index+1, size, target, ans)
}

// find the last occurance of target
func lastOccurance(arr []int, index, size, target int, ans *int) {
	if index < 0 {
		return
	}

	if arr[index] == target {
		*ans = index
		return
	}
	lastOccurance(arr, index-1, size, target, ans)
}

// find the all occurance of target
func allOccurance(arr []int, index, size, target int, ans *[]int) {
	if index == size {
		return
	}

	if arr[index] == target {
		*ans = append(*ans, index)
	}
	allOccurance(arr, index+1, size, target, ans)
}

/********String Problem*********/
func reverseString(str string) string {

	//approach 1
	// strbyte := []byte(str)
	// reverseStringHelper(strbyte, 0, len(str)-1)
	// return string(strbyte)

	//approach 2
	strbyte := []byte(str)
	var ans []byte
	reverseStringHelper1(strbyte, &ans, 0, len(str))
	return string(ans)
}

// approach 1
// keep the 2 pointer which points to start and end character of the string then swap these character until you reach to middlw of staring
func reverseStringHelper(str []byte, i, j int) {
	if i >= j {
		return
	}

	if str[i] != str[j] {
		str[i], str[j] = str[j], str[i]
		fmt.Println("str : ", string(str))
	}
	reverseStringHelper(str, i+1, j-1)
}

// approach 2
// traverse till end of string and while returning apppend each charater to string
func reverseStringHelper1(str []byte, ans *[]byte, currLength, totalLength int) {
	if currLength == totalLength {
		return
	}
	reverseStringHelper1(str, ans, currLength+1, totalLength)
	*ans = append(*ans, str[currLength])
	fmt.Println("ans  : ", ans)
}

/*
subsequence of substring

relative ordering
include and exclude pattern

2^n where n is the string length


                               "ab"     ""
                              /    \
                 include     /      \  exclude
				 "ab"  "a"         "ab"   ""
		i=0		  ^                       / \
						/ \              /   \
			include    /   \  exclude    "b"   ""
				i=1     "ab"   "a"

				the value at the lwaf nodes will be ans

*/

func subString(str string, index int) {
	strByte := []byte(str)
	subStringHelper(strByte, index, "")
}

func subStringHelper(str []byte, index int, ans string) {

	//base condition
	if index == len(str) {
		fmt.Println("ans : ", ans)
		return
	}

	//include
	subStringHelper(str, index+1, ans+string(str[index]))

	//exclude
	subStringHelper(str, index+1, ans)
}

// need to figure it out why it is not working
func sumString(a, b string) string {
	ans := ""
	sumStringHelper(a, b, len(a)-1, len(b)-1, 0, &ans)
	return ans
}

func sumStringHelper(a, b string, i, j int, carry int, ans *string) {
	//base case
	if i < 0 && j < 0 && carry == 0 {
		return
	}

	first, second := 0, 0

	if i >= 0 {
		first = int(a[i] - '0')
	}

	if j >= 0 {
		second = int(b[j] - '0')
	}

	sum := first + second + carry
	lastDigit := int(sum) % 10
	carry = int(sum) / 10

	*ans = strconv.Itoa(lastDigit) + *ans
	sumStringHelper(a, b, i-1, j-1, carry, ans)
}

//recurssion on Linked list
/*
1. print the linked list

solution==>
   i)
     print the number until head become nil
	 func solve (node *Node) {
		 if head == nil { return }
		 fmt.Println(head.data)
        solve(node.Next)
	 }


2. print the kth element from last/end
      solution ==>
	  trverse till end of linked list and thn while returning do the operations
	  handle if list is empty and k is grater than length of linked list

	    func solve(node *Node, k int) {

			if head == nil { return }
			solve(node.Next, k)
             k = k-1
			if k == 0 {
				fmt.Println(node.data)
			}
		}
*/

/*********Divide and conquer Problem*********/

func test(arr []int) {
	if len(arr) <= 1 {
		fmt.Printf("%v\n", arr[s])
		return
	}
	fmt.Println("arr : ", arr)
	arr = arr[0 : (len(arr)+1)/2]
	test(arr)
}

func test1(arr []int, s, e int) {

	if s >= e {
		fmt.Printf("%v\n", arr[s])
		return
	}

	for i := s; i <= e; i++ {
		fmt.Printf("%v ,", arr[i])
	}
	fmt.Println()

	mid := (s + e) / 2
	test1(arr, s, mid)
}

func mergesort(arr []int, s, e int) {

	//b.c
	if s >= e {
		return
	}

	mid := s + (e-s)/2

	//mergesort on left
	mergesort(arr, s, mid)

	//mergesort on right
	mergesort(arr, mid+1, e)

	merge(arr, s, mid, e)
}

remaining
func merge(arr []int, s, mid, e int) {

	len1 := mid - e + 1
	len2 := e - mid

	first := make([]int, 0)
	second := make([]int, 0)

	index1 := 0
	for i := s; i <= mid; i++ {
		first[index1] = arr[i]
		index1++
	}

	index2 := 0
	for i := mid + 1; i <= e; i++ {
		second[index2] = arr[i]
		index2++
	}

	i:=0
	j:=0
	k:=0

}


//medium level problems
func lastRemaining(n int) int {
    
}