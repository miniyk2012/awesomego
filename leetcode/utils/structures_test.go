package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestFromSlice2ListNodet(t *testing.T) {
	s1 := []int{1, 2, 3, 4}
	l1 := FromSlice2ListNode(s1)
	fmt.Printf("%v\n", l1)
	s2 := []int{}
	l2 := FromSlice2ListNode(s2)
	fmt.Printf("%v\n", l2)
}

func TestFromListNode2Slice(t *testing.T) {
	cases := []struct {
		s []int
	}{
		{[]int{1, 2, 3, 4}},
		{[]int{1, 2}},
		{[]int{}},
	}
	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(r *testing.T) {
			l := FromSlice2ListNode(c.s)
			t.Logf("%v\n", l)
			actual := FromListNode2Slice(l)
			if !reflect.DeepEqual(c.s, actual) {
				t.Fatalf("actual %v != expected %v", actual, c.s)
			}
		})
	}

}
