package dataStruct

import "errors"

// 二叉树
type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

/*
 * @brief:	前序遍历
 * @param:	root:二叉树根节点;
 * @param:	flag:遍历方式选择，1为前序，0位中序，-1位后序;
 * @ret: res:前序遍历切片;
 */
func Traversal(root *BinaryTreeNode, flag int) ([]int, error) {
	var perOrder, midOrder, tailOrder func(root *BinaryTreeNode)
	var res []int
	// 前序遍历闭包
	perOrder = func(root *BinaryTreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		perOrder(root.Left)
		perOrder(root.Right)
	}
	// 中序遍历闭包
	midOrder = func(root *BinaryTreeNode) {
		if root == nil {
			return
		}
		midOrder(root.Left)
		res = append(res, root.Val)
		midOrder(root.Right)
	}
	// 后序遍历闭包
	tailOrder = func(root *BinaryTreeNode) {
		if root == nil {
			return
		}
		tailOrder(root.Left)
		res = append(res, root.Val)
		tailOrder(root.Right)
	}

	// 根据flag选择遍历方式
	switch flag {
	case 1:
		perOrder(root)
	case 0:
		midOrder(root)
	case -1:
		tailOrder(root)
	default:
		return nil, errors.New("wrong falg")
	}
	return res, nil
}
