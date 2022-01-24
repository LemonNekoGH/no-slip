package main

import (
	"os/exec"

	"github.com/alecthomas/kong"
	"github.com/sqweek/dialog"
)

// 定义命令行参数
type CLI struct {
	ExecPath string   `required:"" help:"要防止手滑的可执行文件路径"`
	Times    int      `default:"1" help:"启动之前的确认次数，默认为 1"`
	Args     []string `help:"要传递给可执行文件的参数"`
}

func main() {
	ctx := kong.Parse(&CLI{})
	flags := ctx.Flags()
	nameFlagMap := map[string]interface{}{}
	// 转换 map 方便获取参数
	for _, flag := range flags {
		nameFlagMap[flag.Name] = ctx.FlagValue(flag)
	}
	// 获取可执行文件路径
	execPath := nameFlagMap["exec-path"].(string)
	if execPath == "" {
		dialog.Message("可执行文件路径为空").Error()
		return
	}
	// 获取确认次数
	times := nameFlagMap["times"].(int)
	if times < 1 {
		dialog.Message("确认次数小于 1 不就没有意义了嘛？").Error()
		return
	}
	// 获取要传递的参数
	args := nameFlagMap["args"].([]string)
	// 开始阻止手滑
	confirmedTimes := 0
	for confirmedTimes < times {
		if dialog.Message("为了防止手滑，需要点确认 %d 次后才会启动\n%s", times-confirmedTimes, execPath).YesNo() {
			confirmedTimes++
		} else {
			dialog.Message("谢谢使用").Info()
			return
		}
	}
	// 确认要启动了，开始启动
	exec.Command(execPath, args...).Run()
}
