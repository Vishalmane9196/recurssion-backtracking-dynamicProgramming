package main

func main() {

	/*******Basic Problems*******/
	// res := fibo(4)
	// fmt.Println("fibo ", res)

	// printDecreasingOrder(4)
	// printIncreasingOrder(4)
	// printDecreaingIncreasingOrder(4)
	// printIncreasingDecreasingOrder(0, 4)

	// res := factorial(50)
	// fmt.Println("factorial ", res)

	/*********Array Problem*********/
	// arr := []int{1, 2, 9, 43, 3, 4, 5, 3}

	// printArray(arr, 0, len(arr))
	// printArrayReverse(arr, len(arr)-1)

	// printArrayReverse1(arr, len(arr)-1)
	// printArrayReverse2(&arr, len(arr)-1)

	// max := arr[0]
	// maxOfArray(arr, &max, 0, len(arr))
	// fmt.Println("maxOfArray : ", max)

	// arr := []int{1, 2, 9, 43, 3, 4, 5, 3}
	// ans := -1
	// firstOccurance(arr, 0, len(arr), 3, &ans)
	// fmt.Println("firstOccurance : ", ans)

	// arr := []int{1, 2, 9, 43, 3, 4, 5, 3}
	// ans := -1
	// lastOccurance(arr, len(arr)-1, len(arr), 3, &ans)
	// fmt.Println("lastOccurance : ", ans)

	// arr := []int{1, 2, 9, 43, 3, 4, 5, 3, 5, 2, 3, 2, 3}
	// ans := []int{}
	// allOccurance(arr, 0, len(arr), 3, &ans)
	// fmt.Println("lastOccurance : ", ans)

	/********String Problem*********/
	// res := reverseString("vishal")
	// fmt.Println(" res  : ", res)

	// str := "abc"
	// subString(str, 0)
	// fmt.Println(" res  : ", res)

	// str1 := "12"
	// str2 := "123"
	// res := sumString(str1, str2)
	// fmt.Println(" sumString  : ", res)

	/*********Divide and conquer Problem*********/
	nums := []int{3, 4, 1, 5, 9, 2}
	// test(nums)
	test1(nums, 0, len(nums)-1)

	//one merge sort and binary search

}

/*

Recurssion -- function calls iteself
   solve one case and rest will be take care b y recurssion

problem's --> solution depend on solution of subproblem

There are baically 3 mains steps we need to figure out in problem

1. Base case               --- mandatory
2. Operation/processing    --- optional
3. Recurssive relation     --- mandatory

callstack is used to maintain the history of recurssion and keep the track of function calls

Types of recursion
1.Head recurssion
 In head recursion, the recursive call, when it happens, comes before other processing in the function (think of it happening at the top, or head, of the function).
 If the recursive call occurs at the beginning of a method, it is called a head recursion. The method saves the state before jumping into the next recursive call.

2.Tail recurssion
 In tail recursion, it’s the opposite—the processing occurs before the recursive call. Choosing between the two recursive styles may seem arbitrary, but the choice can make all the difference.
 If the recursive call occurs at the end of a method, it is called a tail recursion. The tail recursion is similar to a loop. The method executes all the statements before jumping into the next recursive call.


A function with a path with a single recursive call at the beginning of the path uses what is called head recursion. The factorial function of a previous exhibit uses head recursion. The first thing it does once it determines that recursion is needed is to call itself with the decremented parameter. A function with a single recursive call at the end of a path is using tail recursion. Refer this article
Example Recursion :

public void tail(int n)         |     public void head(int n)
{                               |     {
	//base case                 |        //base case
    if(n == 1)                  |         if(n == 0)
        return;                 |             return;
    else                        |         else
	//processing                |         //recurssive relation
        System.out.println(n);  |             head(n-1);
                                |
	//recurssive relation       |         //processing
    tail(n-1);                  |         System.out.println(n);
}                               |     }

*/
