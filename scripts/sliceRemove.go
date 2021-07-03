package main

import "fmt"
/*
切片或者列表删除功能
*/
func main(){
	var groupslic =  []int{1,2,3,4,5}
	for i :=0;i<len(groupslic);i++{
		if groupslic[i] == 1{
			groupslic = append(groupslic[:i],groupslic[i+1:]...)
		}
	}
	fmt.Println(groupslic)
}
