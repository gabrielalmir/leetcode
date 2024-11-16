# LeetCode Solutions

This repository contains solutions to various LeetCode problems implemented in Go.

## Problem 5: Longest Palindromic Substring

Sure! Let's break down the solution for finding the longest palindromic substring in a given string.

### Summary

The solution efficiently finds the longest palindromic substring by expanding around each character and each pair of characters in the string, ensuring that the time complexity remains O(n^2). This approach is effective for the given problem and provides the correct result. I optimized the algorithm by increasing space complexity to reduce time complexity. This means I used more memory to store intermediate results, which allowed the algorithm to run faster.

![Problem 5: Longest Palindromic Substring Results](./assets/p5_longest_palindrome.png)

### Solution

The solution to the Longest Palindromic Substring problem can be found in the `p5_longest_palindrome.go` file. The approach used is expanding around the center to find the longest palindrome.

### Usage

To run the solution, use the following command:

```sh
go run solutions/p5_longest_palindrome.go
```

### Example

For the input string `"cbbd"`, the output will be:

```
bb
```

### Author

This solution was implemented by [gabrielalmir](https://github.com/gabrielalmir).

### License

This project is licensed under the MIT License.
