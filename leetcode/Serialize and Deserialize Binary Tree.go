package leetcode

// type Codec struct {
// }

// func Constructor() Codec {
// 	return Codec{}
// }

// // Serializes a tree to a single string.
// // result will like: [1 2 3 4 nil nil 5 nil nil nil nil]
// func (this *Codec) serialize(root *TreeNode) string {
// 	nodeList := []string{}
// 	queue := []*TreeNode{root}
// 	for len(queue) > 0 {
// 		node := queue[0]
// 		if node == nil {
// 			nodeList = append(nodeList, "nil")
// 		} else {
// 			nodeList = append(nodeList, strconv.Itoa(node.Val))
// 			queue = append(queue, node.Left, node.Right)
// 		}
// 		queue = queue[1:]
// 	}
// 	return fmt.Sprint(nodeList)
// }

// // Deserializes your encoded data to tree.
// func (this *Codec) deserialize(data string) *TreeNode {
// 	//remove '[' and ']'
// 	data = data[1 : len(data)-1]
// 	if data == "nil" {
// 		return nil
// 	}
// 	nodeList := strings.Split(data, " ")

// 	rootVal, _ := strconv.Atoi(nodeList[0])
// 	root := &TreeNode{Val: rootVal}

// 	queue := []*TreeNode{root}
// 	index := 1
// 	for len(queue) > 0 {
// 		node := queue[0]
// 		//left node
// 		if nodeList[index] != "nil" {
// 			val, _ := strconv.Atoi(nodeList[index])
// 			left := &TreeNode{Val: val}
// 			node.Left = left
// 			queue = append(queue, left)
// 		}
// 		//right node
// 		if nodeList[index+1] != "nil" {
// 			val, _ := strconv.Atoi(nodeList[index+1])
// 			right := &TreeNode{Val: val}
// 			node.Right = right
// 			queue = append(queue, right)
// 		}
// 		queue = queue[1:]
// 		index += 2
// 	}

// 	return root
// }

// /**
//  * Your Codec object will be instantiated and called as such:
//  * ser := Constructor();
//  * deser := Constructor();
//  * data := ser.serialize(root);
//  * ans := deser.deserialize(data);
//  */

// func TestCodec() {
// 	node1 := TreeNode{Val: 1}
// 	node2 := TreeNode{Val: 2}
// 	node3 := TreeNode{Val: 3}
// 	node4 := TreeNode{Val: 4}
// 	node5 := TreeNode{Val: 5}
// 	node1.Left = &node2
// 	node1.Right = &node3
// 	node2.Left = &node4
// 	node3.Right = &node5

// 	ser := Constructor()
// 	deser := Constructor()
// 	data := ser.serialize(&node1)
// 	fmt.Println(data)
// 	ans := deser.deserialize(data)
// 	fmt.Printf("root: %s", ans)
// }
