/*
Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3 
Explanation: The answer is "abc", with the length of 3. 
Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3. 
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/
package main
import ( 
	"os" 
	"fmt"
)

//This is a brute force methods. Code works but I am at bottom 5%
func lengthOfLongestSubstring_brute_force(s string) int {
	max := 0
	n := len(s)
	//println("main string",s,"len",n)
	if n < 2 {
		return(n)
	}
	for i:= 0 ; i < n-1; i++ {
		subs := s[i:]
		//println("sub string",i+1,subs)
		nn := len(subs)
		chars := make(map[byte]int,0)
		j := 0
		for j = 0 ; j < nn ; j++ {
			chars[subs[j]]++
			if chars[subs[j]] > 1 {
				//println("found unique string of len",j, subs[0:j])
				if j > max {
					max = j
				}
				break
			}
		}
		if j == nn { //End of string
			//println("this ending string is unique")
			if j > max {
				max = j
			}
		}
		//fmt.Println(chars)
	}
	return(max)
}

//This one is much better. Runs within 8 ms. Now I am in top 23%
func lengthOfLongestSubstring(s string) int {
	//A hashmap to hold the position of the character. E.g. if incoming string is abcd, this will store {a:0,b:1,c:2,d:3}
	chars := make(map[rune]int,0)

	deleted_until := 0
	max := 0 //This is the final return value

	var i int
	var c rune
	//Pull out characters from string one at a time
	for i,c = range(s) { //i is index, c is the character
		val,exists := chars[c] //Does it exist in map
		if exists {
			//If so, number of characters in map, is the current max string
			ln := len(chars)

			//Is it more than cumulative max
			if ln > max {
				max = ln
			}
			//Now delete some characters from map e.g. if the string is abcdefc, as soon as 2nd c is encountered, delete abc from map
			for j:=deleted_until;j<=val;j++ {
				delete(chars,rune(s[j])) //Delete x characters of incoming string from map
			}
			deleted_until = val+1 //Move the deletion index ahead like a sliding window
			chars[c] = i //Now add the character back into the map but with new position
		} else {
			chars[c] = i
		}
	}
	if i == 0 {
		return len(s)
	}
	if len(chars) > max {
		max = len(chars)
	}
	return(max)
}


func main() {
	for i,par := range(os.Args) {
		if i == 0 {
			continue
		}
		fmt.Printf("%s %d\n",par,lengthOfLongestSubstring(par))
	}
}
