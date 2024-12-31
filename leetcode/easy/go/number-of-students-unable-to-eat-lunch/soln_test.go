package numberofstudentsunabletoeatlunch

import (
	"testing"
)

func TestSolnQueue(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	q := createQueue(arr)
	if q.head.nodeType != 1 && q.tail.nodeType != 4 {
		t.Fail()
	}

	q.accept()
	if q.head.nodeType != 2 {
		t.Log("Accept failed")
		t.Fail()
	}

	q.requeue()
	if q.head.nodeType != 3 && q.tail.nodeType != 2 {
		t.Log("Requeue failed")
		t.Fail()
	}
}

func TestSoln1(t *testing.T) {
	students := []int{1, 1, 0, 0}
	food := []int{0, 1, 0, 1}

	n := countStudents(students, food)
	if n != 0 {
		t.Log("Failed count: ", n)
		t.Fail()
	}
}

func TestSoln2(t *testing.T) {
	students := []int{1, 1, 1, 0, 0, 1}
	food := []int{1, 0, 0, 0, 1, 1}

	n := countStudents(students, food)
	if n != 3 {
		t.Log("Failed count: ", n)
		t.Fail()
	}
}

func TestSoln3(t *testing.T) {
	students := []int{0, 0, 1, 1, 0, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1, 1}

	food := []int{1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 1, 1, 0}

	n := countStudents(students, food)
	if n != 1 {
		t.Log("Failed count: ", n)
		t.Fail()
	}
}
