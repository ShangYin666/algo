package linkedlist

import "testing"

func Test_circleSingleLinkedList_Add(t *testing.T) {
	type fields struct {
		size int
		head *singleNode
		last *singleNode
	}
	type args struct {
		index   int
		element interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &circleSingleLinkedList{
				size: tt.fields.size,
				head: tt.fields.head,
				last: tt.fields.last,
			}
			cs.Add(tt.args.index, tt.args.element)
		})
	}
}
