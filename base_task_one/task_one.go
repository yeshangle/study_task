package main

import "strconv"
import "sort"

// 136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。可以使用 for 循环遍历数组，
// 结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。

func singleNumber(nums []int) int {
    
    single := 0 
    for _, num := range nums{
        single ^= num 
    }
   return single
   // 异或 相同为0 相异为1 
    
}


// 回文数给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

func isPalindrome(x int) bool {
    
    // 将数字转化为对应的字符串
    numStr := strconv.Itoa(x)
    i, j := 0 , len(numStr)-1
    for i < j{
        if numStr[i] != numStr[j]{
            return false
        }
        i++
        j--
    }
    return true

}

// 20 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。 

func isValid(s string) bool {
    
    // st := NewStack()
    baseMark := map[byte]byte{
        ')': '(', 
        ']': '[',
        '}': '{',
    }
    stack := []byte{}
    for i:=0; i<len(s); i++{
        if baseMark[s[i]] > 0 {
            if len(stack) == 0 || stack[len(stack)-1] != baseMark[s[i]]{
                return false
            }
            stack = stack[:len(stack)-1]
        }else{
            stack = append(stack, s[i])
        }
    }
    return len(stack) == 0
}
// 14 最长公共前缀 编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，返回空字符串 ""
func longestCommonPrefix(strs []string) string {

    if len(strs) == 0 {
        return ""
    }
    prefix, count := strs[0], len(strs)
   
    for  index := 1; index < count; index++ {
        // 不断获取他的最长公共子串
        prefix = lcp(prefix, strs[index])
        if len(prefix) == 0{
            break
        }
    }
    return prefix
}

func lcp(prefix, str string) string{
    
    count := min(len(prefix), len(str))
    i := 0
    for i < count && prefix[i] == str[i]{
        i ++
    }

    return prefix[:i]

}
// 66 加一 给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。

func plusOne(digits []int) []int {

    n := len(digits)
    for i := n-1; i >= 0; i--{
        if digits[i] < 9{
            digits[i]++
            return digits
        }
        // 如果为9则进一
        digits[i] = 0
    }
    // 到这里说明所有的都是9, 需要使用make进行初始化
    newDigit := make([]int, n)
    return  append([]int{1}, newDigit...)
}

// 26删除有序数组中的重复项
func removeDuplicates(nums []int) int {
    
    // 双指针法用于处理对应的数据
    if len(nums) == 0{
        return 0
    }
    // 同向快慢指针
    slow := 1
    for fast:=1; fast < len(nums); fast++{
        if nums[fast] != nums[fast-1]{
            nums[slow] = nums[fast]
            slow++
        }
    }
    return slow
}

// 56合并区间以数组 intervals 表示若干个区间的集合，
// 其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 
func merge(intervals [][]int) [][]int {
    //首先进行对应的二维数组的排序
    sort.Slice(intervals, func(i, j int) bool{
        if intervals[i][0] == intervals[j][0]{
            return intervals[i][1] < intervals[j][1]
        }
        return intervals[i][0] < intervals[j][0]
    })
    // 以上排序完成
    retArr := [][]int{}
    for _, intNum := range intervals{
        if len(retArr) == 0 || retArr[len(retArr)-1][1] < intNum[0]{
            // 数组中没有或者数组中1位置小于intNum的0位置
            retArr = append(retArr, intNum)
        }else{
            // 找到最大的那个1的位置进行扩张
            retArr[len(retArr)-1][1] = max(retArr[len(retArr)-1][1], intNum[1])
        }
    }
    return retArr
}

// 1 两数之和，需要使用结构体，给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
func twoSum(nums []int, target int) []int {

    // 由于不能动对应的Index所以需要借助对应的结构体
    type Param struct{
        value int
        index int 
    }
    paramList := []Param{}
    for key, v := range nums{
        paramList = append(paramList, Param{value:v, index:key})
    }

    // 先进行相关的排序
    sort.Slice(paramList, func (i, j int) bool{
        return paramList[i].value < paramList[j].value
    })
    // 使用双指针的算法
    sum := 0
    i, j := 0, len(nums)-1
    for i < j{
        sum = paramList[i].value + paramList[j].value
        if sum == target{
            return []int{paramList[i].index, paramList[j].index}
        }else if sum < target{
            i++
        }else{
            j--
        }
    } 
    return nil  
}

func main(){

	ret := singleNumber([]int{4,1,2,1,2})
	println(ret)
}