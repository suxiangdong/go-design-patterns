package bridge

type Phone interface {
	SoftRun() string
}

type Soft interface {
	Run() string
}

type GameSoft struct{}

func (s *GameSoft) Run() string {
	return "game run\n"
}

type MessageSoft struct{}

func (s *MessageSoft) Run() string {
	return "message run\n"
}

func NewPhoneA(soft Soft) *PhoneA {
	return &PhoneA{
		soft,
	}
}

type PhoneA struct {
	Soft
}

func (p *PhoneA) SoftRun() string {
	return p.Run()
}

func NewPhoneB(soft Soft) *PhoneB {
	return &PhoneB{
		soft,
	}
}

type PhoneB struct {
	Soft
}

func (p *PhoneB) SoftRun() string {
	return p.Run()
}
