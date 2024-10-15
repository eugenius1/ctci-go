package q13binode

import (
	"reflect"
	"testing"
)

func TestTreeToList(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		tree  *BiNode
		slice []int
	}{
		{
			name:  "nil returns empty slice",
			tree:  nil,
			slice: []int{},
		},

		// 0
		{
			name:  "single node",
			tree:  &BiNode{nil, nil, 0},
			slice: []int{0},
		},

		/* 0
		  /
		-1
		*/
		{
			name:  "left child only",
			tree:  &BiNode{&BiNode{data: -1}, nil, 0},
			slice: []int{-1, 0},
		},

		/*
			0
			 \
			  1
		*/
		{
			name:  "right child only",
			tree:  &BiNode{nil, &BiNode{data: 1}, 0},
			slice: []int{0, 1},
		},

		/* 0
		  / \
		-1   1
		*/
		{
			name:  "left and right children",
			tree:  &BiNode{&BiNode{data: -1}, &BiNode{data: 1}, 0},
			slice: []int{-1, 0, 1},
		},

		/*   0
		   /   \
		-4       4
		  \     /
		  -2   2
		*/
		{
			name:  "inward tree",
			tree:  &BiNode{&BiNode{data: -4, node2: &BiNode{data: -2}}, &BiNode{node1: &BiNode{data: 2}, data: 4}, 0},
			slice: []int{-4, -2, 0, 2, 4},
		},

		/*     0
		     /   \
		  -4       4
		  /         \
		-6           6
		*/
		{
			name:  "outward tree",
			tree:  &BiNode{&BiNode{node1: &BiNode{data: -6}, data: -4}, &BiNode{data: 4, node2: &BiNode{data: 6}}, 0},
			slice: []int{-6, -4, 0, 4, 6},
		},

		/*     0
		     /   \
		  -4       4
		  / \     / \
		-6  -2   2   6
		*/
		{
			name:  "full tree",
			tree:  &BiNode{&BiNode{&BiNode{data: -6}, &BiNode{data: -2}, -4}, &BiNode{&BiNode{data: 2}, &BiNode{data: 6}, 4}, 0},
			slice: []int{-6, -4, -2, 0, 2, 4, 6},
		},
	}

	for _, tt := range tests {
		// original representation
		treeIterator := TreeIterator(tt.tree)

		treeSlice := treeIterator.GetAll()
		if !reflect.DeepEqual(treeSlice, tt.slice) {
			t.Errorf("Got %v before conversion, but expected %v", treeSlice, tt.slice)
		}

		// conversion
		listIterator := ListIterator(TreeToList(tt.tree))

		listSlice := listIterator.GetAll()
		if !reflect.DeepEqual(listSlice, tt.slice) {
			t.Errorf("Got %v after conversion, but expected %v", listSlice, tt.slice)
		}

		// TODO: check conversion was in-place
	} //nolint:wsl
}
