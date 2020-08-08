package state

type State interface {
	wood() string
	next(worker *Worker)
}

type Worker struct {
	state State
}

func NewWorker() *Worker {
	return &Worker{
		state: &MondayState{},
	}
}

func (w *Worker) Wood() string {
	return w.state.wood()
}

func (w *Worker) Next() {
	w.state.next(w)
}

func (w *Worker) SetDay(state State) {
	w.state = state
}

type MondayState struct{}

func (s *MondayState) wood() string {
	return "星期一综合症\n"
}

func (s *MondayState) next(worker *Worker) {
	worker.state = &Tuesday{}
}

type Tuesday struct{}

func (s *Tuesday) wood() string {
	return "周日怎么还那么远啊\n"
}

func (s *Tuesday) next(worker *Worker) {
	worker.state = &Wednesday{}
}

type Wednesday struct{}

func (s *Wednesday) wood() string {
	return "前不着村后不着店，崩溃\n"
}

func (s *Wednesday) next(worker *Worker) {
	worker.state = &Thursday{}
}

type Thursday struct{}

func (s *Thursday) wood() string {
	return "再忍忍\n"
}

func (s *Thursday) next(worker *Worker) {
	worker.state = &Friday{}
}

type Friday struct{}

func (s *Friday) wood() string {
	return "最后一天，高兴\n"
}

func (s *Friday) next(worker *Worker) {
	worker.state = &Saturday{}
}

type Saturday struct{}

func (s *Saturday) wood() string {
	return "哈哈，周末来了\n"
}

func (s *Saturday) next(worker *Worker) {
	worker.state = &Sunday{}
}

type Sunday struct{}

func (s *Sunday) wood() string {
	return "万恶的周一又要来了\n"
}

func (s *Sunday) next(worker *Worker) {
	worker.state = &MondayState{}
}
