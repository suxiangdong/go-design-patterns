package abstract_factory

type IFactory interface {
	CreatePhone() IPhone
	CreateComputer() IComputer
}

type IPhone interface {
	Call() string
}

type IComputer interface {
	Work() string
}

type MacFactory struct{}

func NewMacFactory() *MacFactory {
	return &MacFactory{}
}

func (f *MacFactory) CreatePhone() IPhone {
	return &MacPhone{}
}

func (f *MacFactory) CreateComputer() IComputer {
	return &MacComputer{}
}

type MacPhone struct{}

func (p *MacPhone) Call() string {
	return "Mac phone call\n"
}

type MacComputer struct{}

func (c *MacComputer) Work() string {
	return "Mac computer work\n"
}

type HuaWeiFactory struct{}

func NewHuaWeiFactory() *HuaWeiFactory {
	return &HuaWeiFactory{}
}

func (f *HuaWeiFactory) CreatePhone() IPhone {
	return &HuaWeiPhone{}
}

func (f *HuaWeiFactory) CreateComputer() IComputer {
	return &HuaWeiComputer{}
}

type HuaWeiPhone struct{}

func (p *HuaWeiPhone) Call() string {
	return "HuaWei phone call\n"
}

type HuaWeiComputer struct{}

func (c *HuaWeiComputer) Work() string {
	return "HuaWei computer work\n"
}
