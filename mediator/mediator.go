package mediator

type Attendee interface {
	SetMediator(*Mediator)
	Talk() string
	Respond() string
}

type John struct {
	mediator *Mediator
}

func (j *John) SetMediator(mediator *Mediator) {
	j.mediator = mediator
}

func (j *John) Talk() string {
	return j.mediator.chat("Lucy")
}

func (j *John) Respond() string {
	return "I am John\n"
}

type Lucy struct {
	mediator *Mediator
}

func (l *Lucy) SetMediator(mediator *Mediator) {
	l.mediator = mediator
}

func (l *Lucy) Talk() string {
	return l.mediator.chat("John")
}

func (l *Lucy) Respond() string {
	return "I am Lucy\n"
}

type Mediator struct {
	John *John
	Lucy *Lucy
}

func (m *Mediator) chat(name string) string {
	if name == "John" {
		return m.John.Respond()
	} else {
		return m.Lucy.Respond()
	}

}
