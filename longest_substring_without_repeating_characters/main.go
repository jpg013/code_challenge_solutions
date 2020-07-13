package main

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

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	// holds current min window slider
	w := 0
	// the max found substring length
	max := 0
	// hash map containing int32 rune characters and the last
	// known index within the string
	hash := make(map[int32]int, len(s))
	// The current sliding substring length
	c := 0
	for idx, ch := range s {
		v, ok := hash[ch]
		if !ok || v < w {
			hash[ch] = idx
			c++
			continue
		}
		max = maxInt(max, c)
		w = v
		c = idx - w
		hash[ch] = idx
	}

	return maxInt(max, c)
}

func main() {
	lengthOfLongestSubstring("abcabcbb")
}
