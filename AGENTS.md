# AGENTS.md

This file provides guidelines for agentic coding agents working in this repository.

本文件为在此仓库工作的 AI 编程代理提供指南。

## Project Overview / 项目概述

**autotag** is a CLI tool for semantic versioning with git tags. It calculates next versions and manages git tags.

**autotag** 是一个用于 Git 标签语义版本管理的 CLI 工具。它计算下一个版本并管理 Git 标签。

- **Language / 语言**: Go 1.26
- **Structure / 结构**: `cmd/` for CLI commands, `tools/` for utilities / `cmd/` 用于 CLI 命令, `tools/` 用于工具函数
- **Dependencies / 依赖**: spf13/cobra (CLI), Masterminds/semver/v3 (semver parsing / 语义版本解析)

## Build/Lint/Test Commands / 构建/检查/测试命令

### Build / 构建
```bash
go build -ldflags "-X main.AppVersion=`autotag current` -X main.AppRevision=`autotag hash` -X main.AppBuildDate=`autotag date`"
```

### Install / 安装
```bash
go install -ldflags "-X main.AppVersion=`autotag current` -X main.AppRevision=`autotag hash` -X main.AppBuildDate=`autotag date`"
```

### Run All Tests / 运行所有测试
```bash
go test ./...
```

### Run Single Test / 运行单个测试
```bash
go test -v -run TestIncPrerelease01 ./tools/
go test -v -run TestNextPhase ./tools/
```

### Format and Vet / 格式化和检查
```bash
go fmt ./...
go vet ./...
```

## Code Style Guidelines / 代码风格指南

### General / 总体原则
- Go 1.26 compatibility; freely use modern Go features up to 1.26 / Go 1.26 兼容性；自由使用 Go 1.26 以内的现代特性
- Keep code simple and readable; avoid unnecessary complexity / 保持代码简洁易读；避免不必要的复杂性
- Use early returns for error conditions / 错误条件使用早期返回

### Package Structure / 包结构
```
cmd/       - CLI command implementations (cobra.Command variables) / CLI 命令实现 (cobra.Command 变量)
tools/     - Utility functions for git operations, version calculations / Git 操作和版本计算的工具函数
main.go    - Entry point, root command setup / 入口点，根命令设置
autotag.go - Version constants (AppVersion, AppRevision, AppBuildDate) / 版本常量
```

### Imports (3 Groups) / 导入（3 个分组）
1. Standard library / 标准库 (`fmt`, `os`, `strings`, `bufio`, etc.)
2. Third-party packages / 第三方包 (`github.com/...`)
3. Internal packages / 内部包 (`github.com/qiuzhanghua/autotag/...`)

Use blank line between groups. / 分组之间使用空行分隔。

Example:
```go
import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"

    "github.com/Masterminds/semver/v3"
    "github.com/qiuzhanghua/autotag/tools"
    "github.com/spf13/cobra"
)
```

### Naming Conventions / 命名约定
- **Files / 文件**: lowercase, simple nouns / 小写，简单名词 (e.g., `version.go`, `hash.go`)
- **Functions / 函数**: PascalCase (e.g., `GitLatestTag`, `IncPrerelease`)
- **Variables / 变量**: camelCase for locals, PascalCase for exported / 局部用驼峰，导出用帕斯卡 (e.g., `latest`, `AppVersion`)
- **Constants / 常量**: PascalCase for exported, camelCase for unexported / 导出具名，局部驼峰
- **Packages / 包**: lowercase, simple words / 小写，简单词 (e.g., `cmd`, `tools`)
- **Test functions / 测试函数**: `Test<FunctionName><Number>` (e.g., `TestIncPrerelease01`)

### Modern Go Patterns (Go 1.21+) / 现代 Go 模式

**Use slices package / 使用 slices 包:**
```go
// Instead of sort.Slice / 替代 sort.Slice
slices.SortFunc(vs, func(a, b semver.Version) int {
    if a.LessThan(&b) {
        return -1
    }
    return 1
})

// Instead of manual contains loop / 替代手动循环
slices.Contains(items, x)
slices.Index(items, x)
```

**Use cmp.Or for defaults / 使用 cmp.Or 设置默认值:**
```go
name := cmp.Or(os.Getenv("NAME"), "default")
```

### Error Handling / 错误处理
- Return errors as last return value / 错误作为最后一个返回值
- Check errors immediately with early return / 立即检查错误并早期返回
- Use `fmt.Printf` for user-facing messages / 用户消息使用 `fmt.Printf`
- Use `log` or `log.Fatalf` for fatal errors in main/context / main/context 中的致命错误使用 `log` 或 `log.Fatalf`
- Test errors with `if err != nil { return ..., err }` pattern / 使用 `if err != nil { return ..., err }` 模式测试错误

```go
func GitLatestTag() string {
    arr, err := GitAllTags()
    if err != nil {
        return "0.0.0"
    }
    // ...
}
```

### Function Comments / 函数注释
- Use godoc-style comments for exported functions / 导出函数使用 godoc 风格注释
- Include blank line after first sentence / 第一句后包含空行
- Reference parameters and return values / 引用参数和返回值

```go
// IncPrerelease increments the prerelease number.
// a.b.c-alpha.x => a.b.c-alpha.(x+1)
// a.b.c-beta.x => a.b.c-beta.(x+1)
// a.b.c => a.b.c    // unchanged
func IncPrerelease(v semver.Version) (vNext semver.Version) {
```

### Package Headers / 包头部
```go
// Package tools
// Copyright © 2021 邱张华 <qiuzhanghua@icloud.com>
package tools
```

### Testing Patterns / 测试模式
- Use `t.Errorf` for assertion failures (not `t.Fatal` in subtests) / 断言失败使用 `t.Errorf`（子测试中不用 `t.Fatal`）
- Compare with `Equal()` method for semver versions / 使用 `Equal()` 方法比较语义版本
- Test naming: `Test<FunctionName><ScenarioNumber>` / 测试命名：`Test<函数名><场景号>`

```go
func TestIncPrerelease02(t *testing.T) {
    ver, _ := semver.NewVersion("v1.2.3-beta.1+build34")
    actual := IncPrerelease(*ver)
    expected, _ := semver.NewVersion("v1.2.3-beta.2+build34")

    if !actual.Equal(expected) {
        t.Errorf("expected: %v, got: '%v'", expected, actual)
    }
}
```

### CLI Commands (Cobra) / CLI 命令
- Define as package-level `var <Name>Cmd = &cobra.Command{}` / 定义为包级变量 `var <名称>Cmd = &cobra.Command{}`
- Use `Use`, `Short`, `Long` fields for command metadata / 使用 `Use`、`Short`、`Long` 字段定义命令元数据
- `Run` function receives `(cmd *cobra.Command, args []string)` / `Run` 函数接收 `(cmd *cobra.Command, args []string)`
- Validate args in `Args` function if needed / 必要时在 `Args` 函数中验证参数

### Printf Formatting / Printf 格式化
- Use `fmt.Printf` for user output with newlines / 用户输出带换行使用 `fmt.Printf`
- Use `fmt.Print` for output without newlines / 不带换行使用 `fmt.Print`
- Use `%s` for strings, `%v` for general values / 字符串用 `%s`，通用值用 `%v`
- Include trailing newline in `Printf` format string / `Printf` 格式字符串包含尾部换行

### File Operations / 文件操作
- Always check errors from `os.Create`, `os.Stat`, etc. / 始终检查 `os.Create`、`os.Stat` 等的错误
- Use `defer file.Close()` after successful `os.Create` / 成功 `os.Create` 后使用 `defer file.Close()`
- Use `bufio.NewWriter` and `writer.Flush()` for writing / 使用 `bufio.NewWriter` 和 `writer.Flush()` 写入
