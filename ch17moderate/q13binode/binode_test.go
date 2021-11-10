package q13binode

import (
	"reflect"
	"testing"
)

func TestTreeToList(t *testing.T) {
	cases := []struct {
		tree  *BiNode
		slice []int
	}{
		{nil, []int{}},

		// 0
		{&BiNode{nil, nil, 0}, []int{0}},

		/* 0
		  /
		-1
		*/
		{&BiNode{&BiNode{data: -1}, nil, 0}, []int{-1, 0}},

		/*
			0
			 \
			  1
		*/
		{&BiNode{nil, &BiNode{data: 1}, 0}, []int{0, 1}},

		/* 0
		  / \
		-1   1
		*/
		{&BiNode{&BiNode{data: -1}, &BiNode{data: 1}, 0}, []int{-1, 0, 1}},

		/*   0
		   /   \
		-4       4
		  \     /
		  -2   2
		*/
		{&BiNode{&BiNode{data: -4, node2: &BiNode{data: -2}}, &BiNode{node1: &BiNode{data: 2}, data: 4}, 0}, []int{-4, -2, 0, 2, 4}},

		/*     0
		     /   \
		  -4       4
		  /         \
		-6           6
		*/
		{&BiNode{&BiNode{node1: &BiNode{data: -6}, data: -4}, &BiNode{data: 4, node2: &BiNode{data: 6}}, 0}, []int{-6, -4, 0, 4, 6}},

		/*     0
		     /   \
		  -4       4
		  / \     / \
		-6  -2   2   6
		*/
		{&BiNode{&BiNode{&BiNode{data: -6}, &BiNode{data: -2}, -4}, &BiNode{&BiNode{data: 2}, &BiNode{data: 6}, 4}, 0}, []int{-6, -4, -2, 0, 2, 4, 6}},
	}

	for _, c := range cases {
		// original representation
		treeIterator := TreeIterator(c.tree)
		treeSlice := treeIterator.GetAll()
		if !reflect.DeepEqual(treeSlice, c.slice) {
			t.Errorf("Got %v before conversion, but expected %v", treeSlice, c.slice)
		}

		// conversion
		listIterator := ListIterator(TreeToList(c.tree))
		listSlice := listIterator.GetAll()
		if !reflect.DeepEqual(listSlice, c.slice) {
			t.Errorf("Got %v after conversion, but expected %v", listSlice, c.slice)
		}

		// TODO: check conversion was in-place
	}
}
