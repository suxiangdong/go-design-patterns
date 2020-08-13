package memento

// 备忘录管理者
type Caretaker struct {
	memento *Memento
}

func (t *Caretaker) Set(memento *Memento) {
	t.memento = memento
}

func (t *Caretaker) Get() *Memento {
	return t.memento
}

type Memento struct {
	hp int
	mp int
}

type GameRole struct {
	hp int
	mp int
}

func (r *GameRole) HitBoss() {
	r.hp = r.hp / 2
	r.mp = r.mp / 4
}

func (r *GameRole) CreateMemento() *Memento {
	return &Memento{
		hp: r.hp,
		mp: r.mp,
	}
}

func (r *GameRole) ReinstateMemento(m *Memento) {
	r.hp = m.hp
	r.mp = m.mp
}
