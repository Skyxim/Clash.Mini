package parser

import (
	"github.com/Clash-Mini/Clash.Mini/cmd"
	"github.com/Clash-Mini/Clash.Mini/cmd/auto"
	"github.com/Clash-Mini/Clash.Mini/cmd/cron"
	"github.com/Clash-Mini/Clash.Mini/cmd/mmdb"
	"github.com/Clash-Mini/Clash.Mini/cmd/protocol"
	"github.com/Clash-Mini/Clash.Mini/cmd/proxy"
	"github.com/Clash-Mini/Clash.Mini/cmd/startup"
	"github.com/Clash-Mini/Clash.Mini/cmd/sys"
	"github.com/Clash-Mini/Clash.Mini/cmd/task"
	"github.com/Clash-Mini/Clash.Mini/log"
)

const (
	logHeader = "cmd.parser"
)

// GetCmdOrDefaultValue 获取命令值或默认值
func GetCmdOrDefaultValue(command cmd.CommandType, defaultValue string) (value cmd.GeneralType) {
	value = GetCmdValue(command, defaultValue)
	if !value.IsValid() {
		value = value.GetDefault()
	}
	return
}

// GetCmdValue 获取命令值
func GetCmdValue(command cmd.CommandType, value string) cmd.GeneralType {
	switch command {
	case cmd.Task:
		return task.ParseType(value)
	case cmd.Protocol:
		return protocol.ParseType(value)
	case cmd.Sys:
		return sys.ParseType(value)
	case cmd.MMDB:
		return mmdb.ParseType(value)
	case cmd.Cron:
		return cron.ParseType(value)
	case cmd.Proxy:
		return proxy.ParseType(value)
	case cmd.Startup:
		return startup.ParseType(value)
	case cmd.Auto:
		return auto.ParseType(value)
	default:
		log.Errorln("[%s] command \"%s\" is not support\n", logHeader, command)
		return cmd.Invalid
	}
}