package leetcode

import (
    "bytes"
    "fmt"
    "sort"
    "strconv"
)

func TwoSum(nums []int, target int) []int {
    var result []int
    m := make(map[int]int)

    for k,v := range nums {
        d := target - v
        fmt.Printf("target - nums[i] = %d, v = %d, sum = %d\n", d, v, d+v)
        if _, ok := m[d]; ok  {
            result = append(result, m[d])
            result = append(result, k)
            fmt.Println("Found!")
            break
        }
        m[v] = k
        fmt.Println(m)
    }
    return result
}

func TwoSumVersion2(nums []int, target int) []int {

    m := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        temp := int(float64(target - nums[i]))
        if _, ok := m[temp]; ok {
            return []int{m[temp], i}
        } else {

            m[nums[i]] = i
        }
    }
    return []int{}
}

// 9. Palindrome Number
func IsPalindrome(x int) bool {
    //s1 := string(x)
    s1 := strconv.Itoa(x)
    s2 := StringReverse(s1)
    fmt.Println(s1)
    fmt.Println(s2)
    if s1 == s2 {
        return true
    }
    return false
}

func StringReverse(s string) string {
    runes := []rune(s)
    fmt.Println(runes)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    fmt.Println("====>",string(runes))
    return string(runes)
}

// 14. Longest Common Prefix
func LongestCommonPrefix(strs []string) string {
    var result bytes.Buffer
    fmt.Println("strs before sort => ", strs)
    sort.Strings(strs)
    fmt.Println("strs after sort => ", strs)
    if strs[0] == strs[len(strs)-1] {
        return strs[0]
    }
    runesFirst := []rune(strs[0])
    runesLast := []rune(strs[len(strs)-1])

    for num, val := range runesFirst {
        if val != runesLast[num] {
            break
        }
        result.WriteString(string(val))
    }

    return result.String()
}

// 20. Valid Parentheses

