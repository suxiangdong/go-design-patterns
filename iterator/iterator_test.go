package iterator

import "testing"

func TestIntIterator(t *testing.T) {
	s := []interface{}{1, 2, 3, 4, 5, 6}
	iterator := NewCustomerIterator(s)
	for ; iterator.HasNext(); iterator.Next() {
		index := iterator.Index()
		if s[index] != iterator.Value() {
			t.Fatal("Iterator test fail")
		}
	}
}
