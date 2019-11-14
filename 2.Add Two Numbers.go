/*
@Time : 2019-11-11 15:17
@Author : silence
@File : 2.Add Two Numbers
@Software: GoLand
*/
package main

import "fmt"

// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit.
// Add the two numbers and return it as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
// Example:
// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
// Explanation: 342 + 465 = 807.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	     Val int
	     Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var l1p,l2p,lr *ListNode
	var flag bool
	l1p = l1
	l2p = l2
	lr = new(ListNode)
	lrp := lr
	lrpHead := lr
	for ;l1p != nil && l2p != nil;{
		if flag {
			lrp.Val = l1p.Val + l2p.Val + 1
			flag = false
		}else{
			lrp.Val = l1p.Val + l2p.Val
		}
		if lrp.Val > 9 {
			lrp.Val = lrp.Val % 10
			flag = true
		}
		l1p = l1p.Next
		l2p = l2p.Next
		lrp.Next = new(ListNode)
		lrpHead = lrp
		lrp = lrp.Next
	}
	if l1p == nil {
		for ;l2p != nil;{
			if flag {
				lrp.Val = l2p.Val + 1
				flag = false
			}else{
				lrp.Val = l2p.Val
			}
			if lrp.Val > 9 {
				lrp.Val = lrp.Val % 10
				flag = true
			}
			l2p = l2p.Next
			lrp.Next = new(ListNode)
			lrpHead = lrp
			lrp = lrp.Next
		}
	}else{
		for ;l1p != nil;{
			if flag {
				lrp.Val = l1p.Val + 1
				flag = false
			}else{
				lrp.Val = l1p.Val
			}
			if lrp.Val > 9 {
				lrp.Val = lrp.Val % 10
				flag = true
			}
			l1p = l1p.Next
			lrp.Next = new(ListNode)
			lrpHead = lrp
			lrp = lrp.Next
		}
	}
	if  flag {
		lrp.Val = 1
	}else{
		lrpHead.Next = nil
	}
	return lr
}

func TwoSum(l1 *ListNode,l2 *ListNode, flag int) *ListNode {
	if l1 == nil && l2 == nil && flag == 0{
		return nil
	}
	var val1,val2 int
	if l1 == nil {
		val1 = 0
	}else{
		val1 = l1.Val
	}
	if l2 == nil {
		val2 = 0
	}else{
		val2 = l2.Val
	}
	sum := val1 + val2 + flag
	l := new(ListNode)
	l.Val = sum % 10
	if l1 == nil {
		l.Next = TwoSum(nil,l2.Next,sum / 10)
		return l
	}
	if l2 == nil {
		l.Next = TwoSum(l1.Next,nil,sum / 10)
		return l
	}
	l.Next = TwoSum(l1.Next,l2.Next,sum / 10)
	return l
}

func addTwoNumbers_recursive(l1 *ListNode, l2 *ListNode) *ListNode {
	return TwoSum(l1,l2,0)
}

func main(){
	var l1,l2 *ListNode
	l1 = new(ListNode)
	l2 = new(ListNode)
	l1.Val = 9
	//l1.Next = new(ListNode)
	//l1.Next.Val = 4
	//l1.Next.Next = new(ListNode)
	//l1.Next.Next.Val = 3

	l2.Val = 1
	l2.Next = new(ListNode)
	l2.Next.Val = 2
	l2.Next.Next = new(ListNode)
	l2.Next.Next.Val = 3

	l := addTwoNumbers_recursive(l1,l2)
	for ;l != nil;l = l.Next{
		fmt.Println(l.Val)
	}
}