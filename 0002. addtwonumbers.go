/* Problem two
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/

package main
import (
	"fmt"
	"math/big"
	"os"
)
type ListNode struct {
  Val int
  Next *ListNode
}


func listToInt(l *ListNode) *big.Int {
	var cur *ListNode = l
	res := new(big.Int).SetInt64(0)
	mul := new(big.Int).SetInt64(1)
	for ; cur != nil ; cur = cur.Next {
		x := new(big.Int).SetInt64(0).Mul(mul,new(big.Int).SetInt64(int64(cur.Val)))
		res = res.Add(res,x)
		mul.Mul(mul, new(big.Int).SetInt64(10))
	} 
	return res
}

func intToList(res *big.Int) *ListNode {
	var rc ListNode = ListNode{}
	cur := &rc
	for res.Cmp(new(big.Int).SetInt64(0)) > 0 {
		x := new(big.Int).SetInt64(0)
		res.DivMod(res, new(big.Int).SetInt64(10), x)
		val := int(x.Int64())
		cur.Val = val
		if cur.Next == nil && (res.Cmp(new(big.Int).SetInt64(0))>0) {
			cur.Next = & ListNode{}
			cur = cur.Next
		}
	}
	return &rc
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//Traverse through list and add corresponding numbers, remember to carry over if sum more than 10
	//It runs faster than 99% and occupies less than 92% others
	p1 := l1
	p2 := l2
	i := 0
	rc := ListNode {} 
	cur := &rc ;
	carry := 0
	for i = 0 ;; i++ {
		//If either of lists runs out, break
		if p1 == nil {
			break
		}
		if p2 == nil {
			break
		}
		if i > 0 { //After first one, add node to chain
			cur.Next = &ListNode {}
			cur = cur.Next
		}
		cur.Val = p1.Val + p2.Val + carry //Compute value
		if cur.Val > 9 {
			carry = 1
			cur.Val = cur.Val - 10
		} else {
			carry = 0
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	if p1 != nil {
		cur.Next = p1
	}
	if p2 != nil {
		cur.Next = p2
	}
	//Traverse the chain and change if affected by carry
	for cur.Next != nil {
		cur.Next.Val = cur.Next.Val + carry
		if cur.Next.Val > 9 {
			cur.Next.Val = cur.Next.Val - 10
			carry = 1
		} else {
			carry = 0
		}
		cur = cur.Next
	}
	//Finally, if there is still a carry, add a new node
	if carry > 0 {
		cur.Next = & ListNode {1, nil} //For carry
		}
	return &rc
}

func addTwoNumbers_working_1(l1 *ListNode, l2 *ListNode) *ListNode {
	//Simple approach. Convert lists into ints (big int, because of large sizes), add them, convert back to list and return
	//It works and passes all tests. This solution is in top 20%
	n1 := listToInt(l1)
	n2 := listToInt(l2)
	n1.Add(n1, n2)
	rc := intToList( n1 )
	return rc
}













func main() {
	if len(os.Args) != 3 {
		print("Usage :- <program> <n1> <n2>") 
		return
	}
	//x3 := ListNode {3, nil}
	//x2 := ListNode {4, &x3 }
	//x1 := ListNode {2, &x2 }
	//fmt.Printf("%d", listToInt(&x1) )
	//l1 := intToList(1000000000000000000000000000001)
	//l1 := intToList(342)
	//n1,_ := strconv.ParseFloat("1000000000000000000000000000001",32)
	//n1,_ := new (big.Int).SetString("10000000000000000000000000000001",10)
	//n1,_ := new (big.Int).SetString("542",10)
	n1,_ := new (big.Int).SetString(os.Args[1],10)
	fmt.Printf("%d\n",n1)
	//n2,_ := new (big.Int).SetString("465",10)
	//n2,_ := new (big.Int).SetString("65",10)
	n2,_ := new (big.Int).SetString(os.Args[2],10)
	fmt.Printf("%d\n",n2)
	//fmt.Printf("%70d\n",n1.Add(n1,n2))
	//ten,_ := new (big.Int).SetString("10",10)
	//x,_ := new (big.Int).SetString("0",10)
	//n1.DivMod(n1,ten,x)
	//fmt.Printf("%70d %d\n", n1, x)
	l1 := intToList( n1 )
	l2 := intToList(n2)
	res := addTwoNumbers(l1,l2)
	fmt.Printf("Final result is %d\n", listToInt( res ) )
	//print(listToFloat(rc))
	//n1,_ := strconv.ParseFloat("465",64)
}
