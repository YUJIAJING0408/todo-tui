package app

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

// 待办事项，具体实现
type Item struct {
	Text string
	Done bool
}

// Model 是应用的核心状态
type Model struct {
	Todos    []TodoItem
	Cursor   int
	Input    textinput.Model
	Adding   bool
	Width    int
	Height   int
	Quitting bool

	Notifiers []Notifier // 所有提醒渠道
	LastCheck time.Time  // 防止重复提醒
}

const dataFile = ".todo.json"

// NewModel 创建并返回初始化的 Model
func NewModel() Model {
	ti := textinput.New()
	ti.Placeholder = "输入新任务..."
	ti.CharLimit = 800
	ti.Focus()

	return Model{
		Todos: []TodoItem{
			&BasicTodo{Title: "学习 Go 语言", DoneFlag: false},
			&BasicTodo{Title: "写一个 TUI 应用", DoneFlag: false},
			&BasicTodo{Title: "喝咖啡 ☕", DoneFlag: true},
		},
		Input:  ti,
		Adding: false,
		Notifiers: []Notifier{
			&SystemNotifier{},
		},
		LastCheck: time.Now(),
	}
}

// loadTodos 从 JSON 文件加载待办列表，失败则返回默认示例
func loadTodos() []Item {
	f, err := os.Open(dataFile)
	if err != nil {
		// 文件不存在或无法打开 → 使用默认数据
		return []Item{
			{Text: "学习 Go 语言", Done: false},
			{Text: "写一个 TUI 应用", Done: false},
			{Text: "喝咖啡 ☕", Done: true},
		}
	}
	defer f.Close()

	var items []Item
	if err := json.NewDecoder(f).Decode(&items); err != nil {
		fmt.Fprintf(os.Stderr, "解析数据文件失败: %v，将使用默认数据\n", err)
		return []Item{
			{Text: "学习 Go 语言", Done: false},
			{Text: "写一个 TUI 应用", Done: false},
			{Text: "喝咖啡 ☕", Done: true},
		}
	}
	return items
}

// Save 将当前待办列表写入 JSON 文件
func (m Model) Save() error {
	f, err := os.Create(dataFile)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ")
	return encoder.Encode(m.Todos)
}

// Init 是 Bubble Tea 的初始化命令
func (m Model) Init() tea.Cmd {
	return DoTick()
}
