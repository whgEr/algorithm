package main

import (
	"fmt"
	"time"
	"math/rand"
)

//冒泡排序算法
func bubble_sort(data []int) []int  {
	l := len(data)
	c := l
	for i:=0; i<l-1; i++ {
		c--
		is_sorted:=false
		
		for j:=0;j<c;j++ {
			if data[j] > data[j+1] {
				data[j],data[j+1] = data[j+1],data[j]
				is_sorted = true
			}
		}

		if !is_sorted {
			return data
		}
		
	}
	return data
}

func main()  {
	rand.Seed(time.Now().UnixNano())
	n:=rand.Intn(20)
	data := rand.Perm(n)
	fmt.Println(data)
	fmt.Println(bubble_sort(data[:]))
}
