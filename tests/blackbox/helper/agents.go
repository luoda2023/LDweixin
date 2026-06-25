//go:build blackbox

// This file registers all agent factories by importing the agent packages.
// Without these blank imports, core.CreateAgent would return "unknown agent".
// Each import triggers the package's init() function which calls core.RegisterAgent.
package helper

import (
	_ "github.com/luoda2023/LDweixin/agent/claudecode"
	_ "github.com/luoda2023/LDweixin/agent/codex"
	_ "github.com/luoda2023/LDweixin/agent/cursor"
	_ "github.com/luoda2023/LDweixin/agent/gemini"
	_ "github.com/luoda2023/LDweixin/agent/opencode"
	_ "github.com/luoda2023/LDweixin/agent/qoder"
)
