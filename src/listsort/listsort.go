/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

const MAX int = 9
func main() {
  unsorted := NewNode( 4, NewNode(3, NewNode(1, NewNode(5, NewNode(8, NewNode(7, NewNode(9, NewNode( 2, NewNode( 6, nil)))))))))   // original order 4, 3, 1, 5, 8, 7, 9, 2, 6

  PrintPointerList( unsorted )

  sorted := SortList(BuildListArrayFromPointers(unsorted))

	PrintArrayList( sorted )


}

func NewNode(Val int, Next *ListNode) *ListNode {
  var n ListNode
	n.Val = Val
	n.Next = Next
  return &n
}

func BuildListArrayFromPointers(unsortedList *ListNode) []ListNode {
  var listNodeArray []ListNode = make([]ListNode, MAX)
	var temp = unsortedList
	for i:=0; i < MAX; i++ {
    listNodeArray[i].Val = temp.Val
		listNodeArray[i].Next = temp.Next 
		temp = temp.Next
	}
	return listNodeArray
}


func PrintPointerList(listHead *ListNode) {
	fmt.Printf("Unsorted List elements are: ")
	temp := listHead
	for temp != nil {
		fmt.Printf("%d ", temp.Val)
		temp = temp.Next
	}
	fmt.Printf("\n")
}

func PrintArrayList(sortedList []ListNode) {
	fmt.Printf("Sorted List elements are: ")
	for i := 0; i < len(sortedList); i++ {
		fmt.Printf("%d ", sortedList[i].Val)
	}
	fmt.Printf("\n")
}

func SortList(list []ListNode) []ListNode {
	n := len(list) - 1
	swapped := true
  for swapped {
		swapped = false;
		for i := 0; i < n; i++ {						// bubble sort
			if list[i].Val < list[i+1].Val {
				list[i], list[i+1] = list[i+1], list[i]
			swapped = true
			}
		}
  }
	return list
}

