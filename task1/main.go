package main

import (
	"sort"
)

// 只出现一次的数字（没找到返回-1）
func appearsOnce(array []int) int {
	mapInt := make(map[int]int)
	for _, v := range array {
		mapInt[v]++
	}
	for k, v := range mapInt {
		if v == 1 {
			return k
		}
	}
	return -1
}

// 有效的括号
func isVerifyParenthetical(str string) bool {
	mapRune := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	strRune := []rune{}
	for _, v := range str {
		switch v {
		case '(', '{', '[':
			strRune = append(strRune, v)
		case ')', '}', ']':
			if len(strRune) == 0 {
				return false
			}
			match := strRune[len(strRune)-1]
			if match != mapRune[v] {
				return false
			}
			strRune = strRune[:len(strRune)-1]
		}
	}
	return len(strRune) == 0
}

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	s0 := strs[0]
	for _, val := range strs {
		i := 0
		for i < len(s0) && i < len(val) && s0[i] == val[i] {
			i++
		}
		s0 = val[:i]
		if len(s0) == 0 {
			return ""
		}
	}
	return s0
}

// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += 1
		if digits[i] < 10 {
			return digits
		}
		digits[i] = 0
	}
	digits = append([]int{1}, digits...)
	return digits
}

// 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

// 合并区间
func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	merged := [][]int{intervals[0]}
	for _, val := range intervals[1:] {
		last := merged[len(merged)-1]
		if val[0] <= last[1] {
			if val[1] > last[1] {
				last[1] = val[1]
			}
		} else {
			merged = append(merged, val)
		}
	}
	return merged
}

// 两数之和
func twoSum(nums []int, target int) []int {
	mapInt := make(map[int]int)
	for index, num := range nums {
		if mIdx, ok := mapInt[target-num]; ok {
			return []int{mIdx, index}
		}
		mapInt[num] = index
	}
	return nil
}

func main() {

}
