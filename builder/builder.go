package builder

// 抽象创建过程接口
type ComputerBuilder interface {
	buildRam() string
	buildHardDisk() string
	buildMonitor() string
}

type Director struct {
	ComputerBuilder
}

func NewDirector(builder ComputerBuilder) *Director {
	return &Director{builder}
}

func (d *Director) CreateComputer() string {
	return d.buildRam() + d.buildHardDisk() + d.buildMonitor()
}

type ComputerBuilderA struct{}

func (c *ComputerBuilderA) buildRam() string {
	return "ram_A"
}

func (c *ComputerBuilderA) buildHardDisk() string {
	return "hardDisk_A"
}

func (c *ComputerBuilderA) buildMonitor() string {
	return "monitor_A"
}

type ComputerBuilderB struct{}

func (c *ComputerBuilderB) buildRam() string {
	return "ram_B"
}

func (c *ComputerBuilderB) buildHardDisk() string {
	return "hardDisk_B"
}

func (c *ComputerBuilderB) buildMonitor() string {
	return "monitor_B"
}
