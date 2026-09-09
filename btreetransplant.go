package piscine

/*
package main

import (

	"fmt"

)

	func main() {
		root := &TreeNode{Data: "4"}
		BTreeInsertData(root, "1")
		BTreeInsertData(root, "7")
		BTreeInsertData(root, "5")
		node := BTreeSearchItem(root, "1")
		rplc := &TreeNode{Data: "3"}
		root = BTreeTransplant(root, node, rplc)
		BTreeApplyInorder(root, fmt.Println)
	}

	type TreeNode struct {
		Left, Right, Parent *TreeNode
		Data                string
	}

	func BTreeInsertData(root *TreeNode, data string) *TreeNode {
		if root == nil {
			return &TreeNode{Data: data}
		}

		if data < root.Data {
			root.Left = BTreeInsertData(root.Left, data)
			root.Left.Parent = root
		} else {
			root.Right = BTreeInsertData(root.Right, data)
			root.Right.Parent = root
		}

		return root
	}

	func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
		if root == nil {
			//return &TreeNode{Data: elem}
			return nil
		}

		if elem == root.Data {
			return root
		}
		if elem < root.Data {
			return BTreeSearchItem(root.Left, elem)
		}
		return BTreeSearchItem(root.Right, elem)
	}

	func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
		if root == nil {
			return
		}

		BTreeApplyInorder(root.Left, f)
		f(root.Data)
		BTreeApplyInorder(root.Right, f)
	}
*/
func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node.Parent == nil {
		root = rplc
	} else if node == node.Parent.Left {
		node.Parent.Left = rplc
	} else {
		node.Parent.Right = rplc
	}

	if rplc != nil {
		rplc.Parent = node.Parent
	}

	return root
}
