package dataStruct

type elemType int
type TreeNode struct {
	Val elemType
	Left *TreeNode
	Right *TreeNode
}

type Stack struct {
	Val elemType
	Next *Stack
}
