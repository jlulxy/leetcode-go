package main

import "fmt"

func main(){
	arr := make([]int,0)
	a1 := append(arr,1)
	a2 := append(a1,2)
	a3 := append(a2,3)
	a4 := append(a3,4)
	a5 := append(a4,5)
	fmt.Println(arr)
	fmt.Printf("%p\n",arr)
	fmt.Printf("%p\n",a1)
	fmt.Printf("%p\n",a2)
	fmt.Printf("%p\n",a3)
	fmt.Printf("%p\n",a4)
	fmt.Printf("%p\n",a5)
}

/**
指针在容量4的时候没有扩容指针不变
[]
0x115f820
0xc420014070
0xc420014080
0xc4200160c0
0xc4200160c0
0xc42001a0c0
 */