func singleNumber(nums []int) int {
    a := 0
    for _,x:=range nums{
        a ^= x
    }
    return a
}
