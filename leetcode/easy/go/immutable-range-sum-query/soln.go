package immutablerangesumquery

import "fmt"

type NumArray struct {
   arr []int 
}


func Constructor(nums []int) NumArray {
  cumSum := make([]int, len(nums) + 1)
  cumSum[0] = 0 
  for i, v := range nums{
    cumSum[i+1] = cumSum[i] + v
  }

  return NumArray{
    arr: cumSum,
  }
}


func (this *NumArray) SumRange(left int, right int) int {
  return this.arr[right+1]-this.arr[left]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

func Exec(){
  numArray := Constructor([]int{-2, 0, 3, -5, 2, -1})
  fmt.Println(numArray.SumRange(0, 2)) // return (-2) + 0 + 3 = 1
  fmt.Println(numArray.SumRange(2, 5)) // return 3 + (-5) + 2 + (-1) = -1
  fmt.Println(numArray.SumRange(0, 5)) // return (-2) + 0 + 3 + (-5) + 2 + (-1) = -3
}
