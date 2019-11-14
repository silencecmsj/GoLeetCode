/*
@Time : 2019-10-29 17:55
@Author : silence
@File : 1 两数之和
@Software: GoLand
*/
package main

import "fmt"

//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//
//你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
//
//示例:
//
//给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]

func twoSum1(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0}
		}else{
			return nil
		}
	}
	var temp int
	for index,value := range nums{
		temp = target - value
		for i := index + 1;i < len(nums);i++{
			if temp == nums[i] {
				return []int{index,i}
			}
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0}
		}else{
			return nil
		}
	}
	numMap := make(map[int]int)
	var temp int
	for index,value := range nums{
		numMap[value] = index

	}
	for index,value := range nums{
		temp = target - value
		if i,ok := numMap[temp];ok && index != i {
			return []int{index,i}
		}
	}
	return nil
}

func twoSum3(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0}
		}else{
			return nil
		}
	}
	numMap := make(map[int]int)
	var temp int
	for index,value := range nums{
		temp = target - value
		if i,ok := numMap[temp];ok && index != i {
			return []int{index,i}
		}
		numMap[value] = index
	}
	return nil
}

func main(){
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum3(nums,target))
}
