package app

import (
	"fmt"
	"strings"
	"time"
)

type TodoType string

const (
	Basic TodoType = "Basic"
	DUE   TodoType = "DUE"
)

type TodoItem interface {
	Text() string // 文字描述
	Done() bool   // 是否完成
	Toggle()
	DueTime() *time.Time // 截止日期，nil表示没有截止日期
	TypeName() TodoType  // 用于显示类型标记，如 "BASIC" / "DUE"
	DisplayText() string // 渲染用文本（可能包含截止时间等信息）
}

// BasicTodo 基础待办
type BasicTodo struct {
	Title    string `json:"title"`
	DoneFlag bool   `json:"done"`
}

func (b *BasicTodo) Text() string {
	return b.Title
}

func (b *BasicTodo) Done() bool {
	return b.DoneFlag
}

func (b *BasicTodo) Toggle() {
	b.DoneFlag = !b.DoneFlag // 反转
}

func (b *BasicTodo) DueTime() *time.Time {
	return nil
}

func (b *BasicTodo) TypeName() TodoType {
	return Basic
}

func (b *BasicTodo) DisplayText() string {
	if b.DoneFlag {
		return b.Title + " [✓]"
	}
	return b.Title + " [ ]"
}

type DueTodo struct {
	Title    string    `json:"title"`
	DoneFlag bool      `json:"done"`
	Due      time.Time `json:"due"`
}

func (d *DueTodo) Text() string {
	return d.Title
}

func (d *DueTodo) Done() bool {
	return d.DoneFlag
}

func (d *DueTodo) Toggle() {
	d.DoneFlag = !d.DoneFlag
}

func (d *DueTodo) DueTime() *time.Time {
	return &d.Due
}

func (d *DueTodo) TypeName() TodoType {
	return DUE
}

func (d *DueTodo) DisplayText() string {
	dueStr := d.Due.Format("01-02 15:04")
	if d.DoneFlag {
		return fmt.Sprintf("%s (截止: %s) [✓]", d.Title, dueStr)
	}
	return fmt.Sprintf("%s (截止: %s) [ ]", d.Title, dueStr)
}

// ParseNewTodo 根据用户输入创建对应的 TodoItem
// 支持格式：
//
//	普通：            "去买菜"
//	带截止时间：      "[2025-12-31 23:59] 完成报告"
func ParseNewTodo(input string) (TodoItem, error) {
	input = strings.TrimSpace(input)
	// 检查是否为带截止日期的格式
	if strings.HasPrefix(input, "[") {
		end := strings.Index(input, "]")
		if end == -1 {
			// 无闭合括号，当作普通任务
			return &BasicTodo{Title: input}, nil
		}
		// 提取括号内的日期字符串
		dateStr := strings.TrimSpace(input[1:end])
		// 解析日期
		loc := time.Local
		dueTime, err := time.ParseInLocation("2006-01-02 15:04", dateStr, loc)
		if err != nil {
			// 日期解析失败，当作普通任务
			return &BasicTodo{Title: input}, nil
		}
		// 提取括号后的任务描述
		desc := strings.TrimSpace(input[end+1:])
		if desc == "" {
			desc = "未命名任务"
		}
		return &DueTodo{
			Title: desc,
			Due:   dueTime,
		}, nil
	}
	// 普通待办
	return &BasicTodo{Title: input}, nil
}
