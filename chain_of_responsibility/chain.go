package chain_of_responsibility

type Handler interface {
	Handle(id int) string
}

type handler struct {
	next Handler
	name string
	id   int
}

func (h *handler) Handle(id int) string {
	if h.id == id {
		return h.name + "\n"
	}
	if h.next != nil {
		return h.next.Handle(id)
	}
	return "no handler\n"
}

func NewHandler(id int, name string, next Handler) Handler {
	return &handler{
		next: next,
		name: name,
		id:   id,
	}
}
