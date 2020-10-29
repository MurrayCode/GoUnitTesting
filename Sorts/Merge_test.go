package Sorts

import(
	"testing"
	"reflect"
)

func TestMergeSort(t *testing.T){
	var a = []int{9,1,6,8,7,2,10,3,4,5}
	var b = []int{1,2,3,4,5,6,7,8,9,10}
	got := BubbleSort(a)
	want := b

	if !reflect.DeepEqual(got,want){
		t.Errorf("got %v want %v, given: %v", got, want, a)
	}
}