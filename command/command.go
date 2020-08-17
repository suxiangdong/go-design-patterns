package command

type Tv struct{}

func NewTv() *Tv {
	return &Tv{}
}

func (t *Tv) TurnOn() string {
	return "turn on\n"
}

func (t *Tv) TurnDown() string {
	return "turn down\n"
}

type Command interface {
	Execute() string
}

type TurnOnCommand struct {
	receiver *Tv
}

func NewTurnOnCommand(receiver *Tv) *TurnOnCommand {
	return &TurnOnCommand{
		receiver: receiver,
	}
}

func (c *TurnOnCommand) Execute() string {
	return c.receiver.TurnOn()
}

type TurnDownCommand struct {
	receiver *Tv
}

func NewTurnDownCommand(receiver *Tv) *TurnDownCommand {
	return &TurnDownCommand{
		receiver: receiver,
	}
}

func (c *TurnDownCommand) Execute() string {
	return c.receiver.TurnDown()
}

type Invoker struct {
	command Command
}

func NewInvoker() *Invoker {
	return &Invoker{}
}

func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

func (i *Invoker) ExecuteCommand() string {
	return i.command.Execute()
}
