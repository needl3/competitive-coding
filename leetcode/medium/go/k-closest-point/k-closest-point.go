package kclosestpoint

import (
	"container/heap"
	"log"
)

type Point struct {
	x        int
	y        int
	distance int
}
type Heap []Point

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i].distance > h[j].distance
}

func (h Heap) Pop() any {
	l := len(h)
	x := h[l-1]
	h = h[:l-1]
	return x
}

func (h Heap) Push(x any) {
	h = append(h, x.(Point))
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func kClosestPoints(points [][]int, k int) [][]int {
	h := Heap{}
	heap.Init(h)
	for _, point := range points {
		h.Push(Point{
			distance: point[0]*point[0] + point[1]*point[1],
			x:        point[0],
			y:        point[1],
		})
		if h.Len() > k {
			h.Pop()
		}
	}

	minPoints := [][]int{}
	for i := 0; i < k; i++ {
		minItem, _ := h.Pop().(Point)
		minPoints = append(minPoints, []int{minItem.x, minItem.y})
	}
	return minPoints
}

func Exec() {
	points := [][]int{{0, 2}, {2, 0}, {2, 2}}
	k := 2

	closestPoints := kClosestPoints(points, k)
	log.Println("Answer: ", closestPoints)
}
