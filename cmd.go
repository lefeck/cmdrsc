package cmdrsc

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os/exec"
	"strings"
	"time"
)

/*
方式二：这种方式：相对来说也很简介， 只定义了原始数据结构对数据进行了规整，为此， 我将额外附加的参数也对外抽象了出来，这样，使用起来相对更加灵活，好用，用户
可以根据自己的情况去适用特定的场景。
for exmaple:
cmdTmpl: lvremove --yes root
parameter:  --reportformat=json", "--units=b", "--nosuffix
这里的parameter有的人可能需要，有的人可能不需要.
*/

type cmdNameOpt struct {
	cmdName string
}

func (c cmdNameOpt) Apply() string {
	return fmt.Sprintf(strings.TrimSpace(fmt.Sprintf(c.cmdName)))
}

type CmdExecutor interface {
	RefreshCmd(cmdTmpl string) (cmd interface{})
	RunCmd(cmd interface{}, parameter string) (string, string, error)
	SetLevel(level logrus.Level)
}

func NewExecutor(log *logrus.Entry, level logrus.Level) CmdExecutor {
	return &Executor{
		log:   log,
		level: level,
	}
}

type Executor struct {
	log   *logrus.Entry
	level logrus.Level
}

func (e *Executor) SetLevel(level logrus.Level) {
	e.level = level
}

func (e *Executor) RefreshCmd(cmdTmpl string) (cmd interface{}) {
	options := &cmdNameOpt{cmdTmpl}
	return options.Apply()
}

//
func (e *Executor) RunCmd(cmd interface{}, parameter string) (string, string, error) {
	if cmdstr, ok := cmd.(string); ok {
		return e.runCmdFromStr(cmdstr, parameter)
	}
	if cmdstr, ok := cmd.(*exec.Cmd); ok {
		return e.runCmdFromObjects(cmdstr)
	}

	return "", "", fmt.Errorf("could not interpret command from %v", cmd)
}

func isExtraParameter(extraParam string) bool {
	if len(extraParam) == 0 {
		return false
	}
	return true
}

func (e *Executor) runCmdFromStr(cmd string, extraParam string) (string, string, error) {
	args := make([]string, 0)
	fields := strings.Fields(cmd)

	if ok := isExtraParameter(extraParam); ok {
		args = append(args, extraParam)
		fields = append(fields, args...)
	}

	name := fields[0]
	if len(fields) > 1 {
		return e.runCmdFromObjects(exec.Command(name, fields[1:]...))
	}
	return e.runCmdFromObjects(exec.Command(name))
}

func (e *Executor) runCmdFromObjects(cmd *exec.Cmd) (string, string, error) {
	var (
		stdout, stderr bytes.Buffer
		errPart        string
	)
	cmdStartTime := time.Now()
	cmdDuration := time.Since(cmdStartTime)

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		errPart = fmt.Sprintf(", Error: %v", err)
		e.level = logrus.ErrorLevel
	}

	outStr, errStr := stdout.String(), stderr.String()

	e.log.WithFields(logrus.Fields{
		"msg":         strings.Join(cmd.Args, " "),
		"duration":    cmdDuration.String(),
		"duration_ns": cmdDuration.Nanoseconds()}).
		Logf(e.level, "stdout: %s%s%s", outStr, errStr, errPart)

	return outStr, errStr, err
}
