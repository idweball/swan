package command

import (
	"bytes"
	"context"
	"io"
	"os"
	"os/exec"
	"time"
)

//Command 命令实例
type Command struct {
	name    string
	args    []string
	envs    []string
	dir     string
	timeout time.Duration
	process *os.Process
}

//New 创建Command实例
func New(name string, args ...string) *Command {
	return &Command{
		name: name,
		args: args,
	}
}

//AddArgs 添加命令行参数
func (c *Command) AddArgs(args ...string) *Command {
	c.args = append(c.args, args...)
	return c
}

//AddEnvs 添加运行时环境变量
func (c *Command) AddEnvs(envs ...string) *Command {
	c.envs = append(c.envs, envs...)
	return c
}

//RunDir 切换到指定目录运行
func (c *Command) RunDir(dir string) *Command {
	c.dir = dir
	return c
}

//SetTimeout 设置运行超时时间
func (c *Command) SetTimeout(timeout time.Duration) *Command {
	c.timeout = timeout
	return c
}

//RunWithPipe 运行指令，将指令的标准输出定向到stdout, 标准错误输出定向到stderr
func (c *Command) RunWithPipe(stdout, stderr io.Writer) (err error) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	if c.timeout == 0 {
		ctx, cancel = context.WithCancel(context.TODO())
	} else {
		ctx, cancel = context.WithTimeout(context.TODO(), c.timeout)
	}
	defer cancel()

	cmd := exec.CommandContext(ctx, c.name, c.args...)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	cmd.Env = append(os.Environ(), c.envs...)
	cmd.Dir = c.dir

	err = cmd.Start()
	if err != nil {
		return
	}

	c.process = cmd.Process

	return cmd.Wait()
}

//Run 运行指令，将标准错误，标准错误合并到output
func (c *Command) Run() (output string, err error) {
	buffer := new(bytes.Buffer)
	err = c.RunWithPipe(buffer, buffer)
	return buffer.String(), err
}

//Stop 终止正在运行的指令
func (c *Command) Stop() error {
	return c.process.Kill()
}
