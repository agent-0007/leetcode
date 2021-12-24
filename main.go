package main

import (
    "crypto/sha256"
    "fmt"
    "gobook/hashmaps"
    "gobook/leetcode"
    "gobook/structs"
    "math/rand"
    "time"

    // "math/rand"
)

func array_test() {
    var a [3]int
    a[0] = 1
    a[1] = 2
    fmt.Println("first array element => ",a[0])
    fmt.Println("last array element => ", a[len(a)-1])

    hashmaps.Test()

    for i,v := range a {
        fmt.Printf("number: %d, value:%d\n", i, v)
    }

    // вывод только элементов
    for _,v := range a {
        fmt.Printf("just element: %d\n", v)
    }

    var q [3] int = [3]int{1,2,3}
    for i,v := range q {
        fmt.Printf("counter: %d, element: %d\n", i, v)
    }

    var q1 = [...]int{
        1,
        2,
        3,
    }
    fmt.Printf("%T\n", q1)

    type Currency int
    const (
        USD Currency = iota
        EUR
        GBR
        RUR
    )
    symbol := [...]string{USD:"$", EUR:"euro", RUR:"Rub"}
    fmt.Println(RUR, symbol[RUR])

    r := [...]int{99:-1}
    for i,v := range r {
        fmt.Printf("counter: %d, element: %d\n", i, v)
    }

    c1 := sha256.Sum256([]byte("x"))
    c2:= sha256.Sum256([]byte("x"))

    fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1==c2, c1)
}

func reverse(s []int) {
    for i,j := 0, len(s)-1; i< j; i,j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}

func slices() {
    months := [...]string{
        1 : "Январь",
        2 : "Февраль",
        3 : "Март",
        4 : "Апрель",
        5 : "Май",
        6 : "Июнь",
        7 : "Июль",
        8 : "Август",
        9 : "Сентябрь",
        10 : "Октябрь",
        11 : "Ноябрь",
        12 : "Декабрь",
    }
    Q2 := months[4:7]
    All := months[1:13]
    // invalid slice index 14 (out of bounds for 13-element array)
    //Crash := months[1:14]
    test := months[:5]
    testFromTest := test[:3]
    fmt.Println(Q2)
    fmt.Println(All)
    fmt.Println(test)
    fmt.Println(testFromTest)
    fmt.Printf("Q2 == test %t\n", equal(Q2, test))
    if test == nil {
        fmt.Println("test is nil")
    } else {
        fmt.Println("test is not nil")
    }
}
func equal(x,y []string) bool {
    if len(x) != len(y) {
        return false
    }
    for i := range x {
        if x[i] != y[i] {
            return false
        }
    }
    return true
}
func test_append() {
    var runes []rune
    for _,r := range "Hello World!" {
        runes = append(runes, r)
    }
    fmt.Printf("%q\n", runes)
}

func appendInt(x []int, y int) []int {
    var z []int
    zlen := len(x) + 1
    if zlen <= cap(x) {
        fmt.Println("Есть место для роста")
        z = x[:zlen]
    } else {
        fmt.Println("Нет места для роста Выделяем новый масив, увеличиваем в 2 раза")
        zcap := zlen
        if zcap < 2*len(x) {
            zcap = 2 * len(x)
        }
        z = make([]int, zlen, zcap)
        copy(z, x)
    }
    z[len(x)] = y
    return z
}

func non_bool_value_in_if() {
    var level int = 850;

    if level&1 != 0 {
        // do something
    } else {
        // do something else
    }
}

func test(x int) bool {

    if x % 2 == 0 {
        return true
    }
    return false
}

/*
    Задачка на джуна с подъебинго, в слайс кладется указатели на переменную item
    по выходу с цикла items = 5
    поэтому вывод значений будет 5,5,5,5,5
    все положенные ссылки будут вести на одно и то же значение
 */
func avitoTest() {
    numbers := []int{1, 2, 3, 4, 5}
    var a []*int
    for _, item := range numbers {
        a = append(a, &item)
    }
    for _, item := range a {
        fmt.Println(*item)
    }
}


func main() {

    array_test()
    slices()
    non_bool_value_in_if()

    a := [...]int{0,1,2,3,4,5,6,7,8,9}
    fmt.Println("Before sort", a)
    reverse(a[:])
    fmt.Println("After reverse", a)
    test_append()

    var x,y []int
    for i := 0; i < 10; i++ {
        y = appendInt(x, i)
        fmt.Printf("%d cap = %d\t%v\n", i, cap(y), y)
        x = y
    }
    res := test(5)
    fmt.Printf("5 divides by 2 = %t\n", res)
    res = test(4)
    fmt.Printf("4 divides by 2 = %t\n", res)

    fmt.Println("Testing go functions")
    var point structs.Point
    point = structs.Point{1, 2}
    fmt.Println(point)
    structs.MakePoint(&point)
    fmt.Println(point)
    structs.Scale(&point, 5)
    fmt.Println("Point after scale => ", point)

    var w structs.Wheel
    structs.MakeWheel(&w)
    fmt.Println("Wheel => ", w)

    fmt.Println("Test maps")
    hashmaps.Test()
    hashmaps.TestMaps()

    //fmt.Println("JSON test")
    //jsontest.JsonTest()
    //jsontest.JsonTestSimple()


    // LeetCode
    numsSumOfTwo := []int{2,7,11,15}
    targetSumOfTwo := 9
    resultSumOfTwo := leetcode.TwoSum(numsSumOfTwo, targetSumOfTwo)
    fmt.Printf("SumOfTwo, target => %d ", targetSumOfTwo)
    fmt.Println("SumOfTwo => ", resultSumOfTwo)

    resultSumOfTwo_v2 := leetcode.TwoSumVersion2(numsSumOfTwo, targetSumOfTwo)
    fmt.Printf("SumOfTwo_v2, target => %d ", targetSumOfTwo)
    fmt.Println("SumOfTwo_v2 => ", resultSumOfTwo_v2)

    //9. Palindrome Number
    fmt.Println("9. Palindrome Number")
    palindrome := 123321
    notPalindrome := 123456
    fmt.Println(leetcode.IsPalindrome(palindrome))
    fmt.Println(leetcode.IsPalindrome(notPalindrome))

    // 14. Longest Common Prefix
    strArray := []string{"flower","flow","flight"}
    fmt.Printf("Longest prefix is => %s", leetcode.LongestCommonPrefix(strArray))
    strArray2 := []string{"flower","flow","aaaflight"}
    fmt.Printf("Longest prefix is => %s\n", leetcode.LongestCommonPrefix(strArray2))

    // 20. Valid Parentheses
    // TODO: Not works ;(
    parenthesesStr := "{[{}]}"
    parenthesesStrNotCorrect := "{[{}]}}}}}"
    parenthesesStr2 := "()()"
    parenthesesStr3 := "()[]{}"
    fmt.Println("-----------------------")
    leetcode.IsValidParentheses(parenthesesStr)
    fmt.Printf("Valid Parentheses at string: %s is %t\n", parenthesesStr, leetcode.IsValidParentheses(parenthesesStr))
    fmt.Printf("Valid Parentheses at string: %s is %t\n", parenthesesStrNotCorrect, leetcode.IsValidParentheses(parenthesesStrNotCorrect))
    fmt.Printf("Valid Parentheses at string: %s is %t\n", parenthesesStr2, leetcode.IsValidParentheses(parenthesesStr2))
    fmt.Printf("Valid Parentheses at string: %s is %t\n", parenthesesStr3, leetcode.IsValidParentheses(parenthesesStr3))

    // 202. Happy Number
    happyNumber := 19
    unHappyNumber := 2
    maxR := 1000
    minR := 100
    // seed rnd
    rand.Seed(time.Now().UnixNano())
    randNumber := rand.Intn(maxR - minR) + minR
    fmt.Println("Happy Number =======>")
    fmt.Printf("Number %d is %t\n", happyNumber, leetcode.IsHappy(happyNumber))
    fmt.Printf("Number %d is %t\n", unHappyNumber, leetcode.IsHappy(unHappyNumber))
    fmt.Printf("Number %d is %t\n", randNumber, leetcode.IsHappy(randNumber))

    // 58. Length of Last Word
    var strTest58 string = "Hello World"
    var strTest581 string = "   fly me   to   the moon  "
    fmt.Printf("Last world length in str: %s is %d\n", strTest58, leetcode.LengthOfLastWord(strTest58))
    fmt.Printf("Last world length in str: %s is %d\n", strTest581, leetcode.LengthOfLastWord(strTest581))

}