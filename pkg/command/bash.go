package command

//NewBashCommand 创建一个bash的Command实例
func NewBashCommand() *Command {
	return New("bash", "-c")
}
