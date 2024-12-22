package immutable2drangesumquery

import (
	"fmt"
)

type NumMatrix struct {
	mat [][]int
}

func Constructor(matrix [][]int) NumMatrix {
        cumMatrix := make([][]int, len(matrix)+1)
        for i:=0;i<=len(matrix);i++ {
          cumMatrix[i] = make([]int, len(matrix[0])+1)
          if i==0{
            // Skip all first row as zero
            continue
          }

          rowCum := 0
          for j:=0;j<=len(matrix[0]);j++{
            if j== 0{
              // Skip all first column as zero
              continue
            }

            rowCum += matrix[i-1][j-1]
            cumMatrix[i][j] = cumMatrix[i-1][j] + rowCum
          }
	}
        fmt.Println(cumMatrix)
	return NumMatrix{
		mat: cumMatrix,
	}
}



func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int { 
  return this.mat[row2+1][col2+1] - this.mat[row1][col2+1] - this.mat[row2+1][col1] + this.mat[row1][col1]
}

func Exec() {
	numMatrix := Constructor([][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}})
	fmt.Println(numMatrix.SumRegion(2, 1, 4, 3)) // return 8 (i.e sum of the red rectangle)
	fmt.Println(numMatrix.SumRegion(1, 1, 2, 2)) // return 11 (i.e sum of the green rectangle)
	fmt.Println(numMatrix.SumRegion(1, 2, 2, 4)) // return 12 (i.e sum of the blue rectangle)
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
