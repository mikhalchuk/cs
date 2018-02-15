package linked_list

import (
	"testing"
)

func TestList_AddFirst(t *testing.T) {

	list := &List{}
	expected := []string{"b", "a"}
	elements := []string{"a", "b"}
	for _, item := range elements {
		list.AddFirst(item)
	}

	for i, item := range list.ToArray() {
		if expected[i] != item.(string) {
			t.Errorf("They should match")
		}
	}
}

func TestList_AddLast(t *testing.T) {

	list := &List{}
	elements := []string{"a", "b"}
	for _, item := range elements {
		list.AddLast(item)
	}

	for i, item := range list.ToArray() {
		if elements[i] != item.(string) {
			t.Errorf("They should match")
		}
	}
}