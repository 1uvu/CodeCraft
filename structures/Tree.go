package structures

type BNode struct {
	Val   interface{}
	Left  *BNode
	Right *BNode
}

//func InorderTraversal(root *BNode) []interface{} {
//	var ans []interface{}
//	var traverse func(node *BNode)
//	traverse = func(root *BNode) {
//		if root == nil {
//			return
//		}
//		traverse(root.Left)
//		ans = append(ans, root.Val)
//		traverse(root.Right)
//	}
//	traverse(root)
//	return ans
//}
//
//func PreorderTraversal(root *BNode) []interface{} {
//	return nil
//}
//
//func PostorderTraversal(root *BNode) []interface{} {
//	return nil
//}
