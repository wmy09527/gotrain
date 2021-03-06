package basetrain

import "fmt"

/*
Go 语言切片是对数组的抽象。

var identifier []type

var slice1 []type = make([]type, len)

也可以简写为

slice1 := make([]type, len)

也可以指定容量，其中 capacity 为可选参数。
make([]T, length, capacity)
这里 len 是数组的长度并且也是切片的初始长度。

切片初始化
s :=[] int {1,2,3 } 
直接初始化切片，[] 表示是切片类型，{1,2,3} 初始化值依次是 1,2,3，其 cap=len=3。

s := arr[:] 
初始化切片 s，是数组 arr 的引用。

s := arr[startIndex:endIndex] 
将 arr 中从下标 startIndex 到 endIndex-1 下的元素创建为一个新的切片。

s := arr[startIndex:] 
默认 endIndex 时将表示一直到arr的最后一个元素。

s := arr[:endIndex] 
默认 startIndex 时将表示从 arr 的第一个元素开始。

s1 := s[startIndex:endIndex] 
通过切片 s 初始化切片 s1。

s :=make([]int,len,cap) 
通过内置函数 make() 初始化切片s，[]int 标识为其元素类型为 int 的切片。

*/

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

func number1() {
	/* 创建切片 */
	numbers := []int{0,1,2,3,4,5,6,7,8}  
	printSlice(numbers)
 
	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)
 
	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])
 
	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])
 
	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])
 
	numbers1 := make([]int,0,5)
	printSlice(numbers1)
 
	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)
 
	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)

}

func number2() {
   var numbers []int
   printSlice(numbers)

   /* 允许追加空切片 */
   numbers = append(numbers, 0)
   printSlice(numbers)

   /* 向切片添加一个元素 */
   numbers = append(numbers, 1)
   printSlice(numbers)

   /* 同时添加多个元素 */
   numbers = append(numbers, 2,3)
   printSlice(numbers)
   // append 分配空间是2的倍数，cap值=2的整数倍
   numbers = append(numbers, 4,5,6,7,8,9,10,11,12,13)
   //numbers = append(numbers, 5,6,7)
   printSlice(numbers)

   /* 创建切片 numbers1 是之前切片的两倍容量*/
   numbers1 := make([]int, len(numbers), (cap(numbers))*2)

   /* 拷贝 numbers 的内容到 numbers1 */
   copy(numbers1,numbers)
   printSlice(numbers1) 
   
}

func VarSlice() {
	//fmt.Println()
	//var numbers = make([]int,3,5)
	//printSlice(numbers)

	number1()
	fmt.Println("#######################################")
	number2()
	fmt.Println("#######################################")

	number3 := []int{1,2,3,4,5,6,7}
	printSlice(number3)

	values := [][]int{}
	var values1 [][]int
	fmt.Printf("len=%d cap=%d slice=%v\n",len(values),cap(values),values)
	if(values == nil){
		fmt.Printf("values切片是空的\n")
	 }
	 if(values1 == nil){
		fmt.Printf("values1切片是空的\n")
	 }

	balance3 := [5]int{1000, 2, 3, 7, 5}
	fmt.Printf("len=%d cap=%d slice=%v\n",len(balance3),cap(balance3),balance3)
	//printSlice(balance3)
}