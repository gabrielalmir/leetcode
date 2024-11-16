package main

func longestPalindrome(s string) string {
	start := 0
	end := 0

	for i := 0; i < len(s); i++ {
		odd := expand(s, i, i)
		even := expand(s, i, i+1)

		maxLen := odd

		if even > maxLen {
			maxLen = even
		}

		if maxLen > (end - start) {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}

	return s[start : end+1]
}

func expand(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func main() {
	println(longestPalindrome("cbbd"))
}
