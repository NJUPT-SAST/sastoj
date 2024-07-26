package data

import (
	"sastoj/app/gojudge/internal/conf"
	"strings"
)

type Commands = map[string]Command

type Command struct {
	Compile       []string
	Run           []string
	Source        string
	Target        string
	CompileConfig *conf.ExecConfig
	RunConfig     *conf.ExecConfig
}

func NewCommands(
	enable []string,
	compile map[string]string,
	run map[string]string,
	source map[string]string,
	target map[string]string,
	configs map[string]*conf.LanguageConfig) Commands {

	res := make(map[string]Command)
	for _, language := range enable {
		r, ok := run[language]
		if !ok {
			r = run["default"]
		}
		c, ok := compile[language]
		if !ok {
			c = ""
			//不编译就添加文件，直接尝试运行
			//c = compile["default"]
		}
		s, ok := source[language]
		if !ok {
			s = source["default"]
		}
		t, ok := target[language]
		if !ok {
			t = target["default"]
		}
		e, ok := configs[language]
		if !ok {
			e = configs["default"]
		}
		res[language] = Command{
			Compile:       strings.Fields(c),
			Run:           strings.Fields(r),
			Source:        s,
			Target:        t,
			CompileConfig: e.Compile,
			RunConfig:     e.Run,
		}
	}
	return res
}
