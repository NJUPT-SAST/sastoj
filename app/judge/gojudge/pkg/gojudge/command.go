package gojudge

import (
	"sastoj/app/judge/gojudge/internal/conf"
	"strings"
)

type Commands = map[string]Command

type Command struct {
	Compile       []string
	Env           []string
	Run           []string
	Source        string
	Target        string
	CompileConfig *conf.ExecConfig
	RunConfig     *conf.ExecConfig
}

func NewCommands(
	enable []string,
	env map[string][]string,
	compile map[string]string,
	run map[string]string,
	source map[string]string,
	target map[string]string,
	configs map[string]*conf.LanguageConfig) Commands {

	res := make(map[string]Command)
	for _, language := range enable {
		e, ok := env[language]
		if !ok {
			e = env["default"]
		}
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
		config, ok := configs[language]
		if !ok {
			config = configs["default"]
		}

		compileCmd := strings.Fields(c)
		index := strings.Index(c, "\"")
		if index != -1 {
			compileCmd = strings.Fields(c[:index])
			compileCmd = append(compileCmd, c[index:])
		}
		runCmd := strings.Fields(r)
		index = strings.Index(r, "\"")
		if index != -1 {
			runCmd = strings.Fields(r[:index])
			runCmd = append(runCmd, r[index:])
		}

		res[language] = Command{
			Env:           e,
			Compile:       compileCmd,
			Run:           runCmd,
			Source:        s,
			Target:        t,
			CompileConfig: config.Compile,
			RunConfig:     config.Run,
		}
	}
	return res
}
