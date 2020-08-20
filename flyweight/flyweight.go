package flyweight

type CharacterFlyweight struct {
	name string
}

func NewCharacterFlyWeight(name string) *CharacterFlyweight {
	return &CharacterFlyweight{name: name}
}

type Factory struct {
	flyweights map[string]*CharacterFlyweight
}

func NewFlyweightFactory() *Factory {
	return &Factory{
		flyweights: make(map[string]*CharacterFlyweight),
	}
}

func (f *Factory) GetFlyweight(name string) *CharacterFlyweight {
	flyweight, ok := f.flyweights[name]
	if !ok {
		flyweight = NewCharacterFlyWeight(name)
		f.flyweights[name] = flyweight
	}
	return flyweight
}
