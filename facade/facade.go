package facade

type SystemInterface interface {
	TurnDown() string
}

type System struct {
	power *Power
	task  *Task
}

func NewFacade() SystemInterface {
	return &System{
		power: &Power{},
		task:  &Task{},
	}
}

func (s *System) TurnDown() string {
	return s.task.Save() + s.power.Close()
}

type Power struct{}

type PowerInterface interface {
	Close() string
}

func (p *Power) Close() string {
	return "power closed\n"
}

type Task struct{}

type TaskInterface interface {
	Save() string
}

func (t *Task) Save() string {
	return "task saved\n"
}
