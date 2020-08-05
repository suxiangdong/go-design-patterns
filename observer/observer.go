package observer

type EventObserver interface {
	handle() string
}

type FirstObserver struct{}

func NewFirstObserver() *FirstObserver {
	return &FirstObserver{}
}

func (o *FirstObserver) handle() string {
	return "first observer\n"
}

type SecondObserver struct{}

func NewSecondObserver() *SecondObserver {
	return &SecondObserver{}
}

func (o *SecondObserver) handle() string {
	return "second observer\n"
}

type EventNotifier interface {
	Register(observer EventObserver)
	Deregister(observer EventObserver)
	Notify() string
}

type Subject struct {
	obs []EventObserver // 通过反射改为map删除复杂度更低
}

func NewSubject() *Subject {
	return &Subject{}
}

func (s *Subject) Register(observer EventObserver) {
	s.obs = append(s.obs, observer)
}

func (s *Subject) Deregister(observer EventObserver) {
	for i := 0; i < len(s.obs); i++ {
		if s.obs[i] == observer {
			s.obs = append(s.obs[:i], s.obs[i+1:]...)
		}
	}
}

func (s *Subject) Notify() string {
	context := ""
	for _, ob := range s.obs {
		context += ob.handle()
	}
	return context
}
