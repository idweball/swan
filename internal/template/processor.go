package template

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"swan/pkg/command"
	"text/template"
)

//Config 应用模板配置
type Config struct {
	Mode      os.FileMode `toml:"mode"`
	User      string      `toml:"user"`
	Group     string      `toml:"group"`
	Src       string      `toml:"src"`
	Dst       string      `toml:"dst"`
	Keys      []string    `toml:"keys"`
	ReloadCmd string      `toml:"reload_cmd"`
}

//Processor 模板处理器
type Processor struct {
	tpl *template.Template
	cfg Config
}

//NewProcessor 创建模板处理器
func NewProcessor(cfg Config) (*Processor, error) {
	var (
		tpl *template.Template
		err error
	)
	tpl = template.New(filepath.Base(cfg.Src))
	tpl = tpl.Funcs(funcMap)
	if tpl, err = tpl.ParseFiles(cfg.Src); err != nil {
		return nil, err
	}
	return &Processor{tpl: tpl, cfg: cfg}, nil
}

//tmpRender 进行临时渲染，避免模板出错，导致渲染结果不正确，然后人为reload, 引发错误的问题
func (p *Processor) tmpRender() error {
	name := filepath.Join(os.TempDir(), filepath.Base(p.cfg.Dst))
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, p.cfg.Mode)
	if err != nil {
		return err
	}
	defer f.Close()
	defer os.Remove(name)

	err = p.tpl.Execute(f, nil)
	if err != nil {
		return err
	}

	return nil
}

//GetTemplate 获取processor处理的模板
func (p *Processor) GetTemplate() string {
	return p.cfg.Src
}

//Render 渲染模板
func (p *Processor) Render() error {
	if err := p.tmpRender(); err != nil {
		return err
	}

	f, err := os.OpenFile(p.cfg.Dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, p.cfg.Mode)
	if err != nil {
		return err
	}
	defer f.Close()

	err = p.tpl.Execute(f, nil)
	if err != nil {
		return err
	}

	if err := chown(p.cfg.Dst, p.cfg.User, p.cfg.Group); err != nil {
		return err
	}

	return p.reload()
}

func (p *Processor) reload() (err error) {
	if len(p.cfg.ReloadCmd) == 0 {
		return nil
	}

	var output string
	output, err = command.NewBashCommand().AddArgs(p.cfg.ReloadCmd).Run()
	if err != nil {
		return fmt.Errorf("failed to execute reload cmd: %s %s %v", p.cfg.ReloadCmd, output, err)
	}
	return nil
}

func chown(name, user, group string) error {
	var (
		uid int
		gid int
		err error
	)
	if uid, err = getUID(user); err != nil {
		return &ChownError{Err:err}
	}
	if gid, err = getGID(group); err != nil {
		return &ChownError{Err:err}
	}
	if err = os.Chown(name, uid, gid); err != nil {
		return &ChownError{Err:err}
	}
	return nil
}

func getUID(name string) (uid int, err error) {
	var u *user.User
	u, err = user.Lookup(name)
	if err != nil {
		return
	}
	return strconv.Atoi(u.Uid)
}

func getGID(name string) (gid int, err error) {
	var g *user.Group
	g, err = user.LookupGroup(name)
	if err != nil {
		return
	}
	return strconv.Atoi(g.Gid)
}

type ChownError struct {
	Err error
}

func (c *ChownError) Error() string {
	return "chown error: " + c.Err.Error()
}
