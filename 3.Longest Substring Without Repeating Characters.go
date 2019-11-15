/*
@Time : 2019-11-14 16:31
@Author : silence
@File : 3.Longest Substring Without Repeating Characters
@Software: GoLand
*/
package main

import (
	"fmt"
)

//Given a string, find the length of the longest substring without repeating characters.
//
//Example 1:
//
//Input: "abcabcbb"
//Output: 3
//Explanation: The answer is "abc", with the length of 3.
//Example 2:
//
//Input: "bbbbb"
//Output: 1
//Explanation: The answer is "b", with the length of 1.
//Example 3:
//
//Input: "pwwkew"
//Output: 3
//Explanation: The answer is "wke", with the length of 3.
//Note that the answer must be a substring, "pwke" is a subsequence and not a substring.




// 滑动窗口解法
func lengthOfLongestSubstring(s string) int {
	var (
		lengthMax,firstIndex,secondIndex int
		storeMap map[byte]int
	)

	storeMap = make(map[byte]int)

	for ;firstIndex < len(s) && secondIndex < len(s);{
		_,ok := storeMap[s[secondIndex]]
		if ok {
			if lengthMax < secondIndex - firstIndex {
				lengthMax = secondIndex - firstIndex
			}
			delete(storeMap,s[firstIndex])
			firstIndex ++
		}else{
			storeMap[s[secondIndex]] = secondIndex
			secondIndex ++
		}
	}
	if lengthMax < secondIndex - firstIndex {
		lengthMax = secondIndex - firstIndex
	}
	return lengthMax
}

func main(){
	a := lengthOfLongestSubstring("abba")
	fmt.Println(a)
}







