package iterator

type Iterator interface {
	Next()
	HasNext() bool
	Value() interface{}
	Index() int
}

type CustomerIterator struct {
	slice []interface{}
	index int
}

func NewCustomerIterator(s []interface{}) *CustomerIterator {
	return &CustomerIterator{
		slice: s,
	}
}

func (i *CustomerIterator) HasNext() bool {
	return len(i.slice) > i.index+1
}

func (i *CustomerIterator) Next() {
	if i.HasNext() {
		i.index++
	}
}

func (i *CustomerIterator) Value() interface{} {
	return i.slice[i.index]
}

func (i *CustomerIterator) Index() int {
	return i.index
}
