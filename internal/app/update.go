package app

import (
	"strings"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

// Update 处理消息并更新模型
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		return m, nil

	case tea.KeyMsg:
		// 正在添加模式：按键由输入框处理
		if m.Adding {
			switch msg.String() {
			case "enter":
				text := strings.TrimSpace(m.Input.Value())
				if text != "" {
					m.Todos = append(m.Todos, Item{Text: text, Done: false})
					// 新增后保存
					_ = m.Save()
				}
				m.Input.SetValue("")
				m.Adding = false
				return m, nil

			case "esc":
				m.Input.SetValue("")
				m.Adding = false
				return m, nil

			default:
				var cmd tea.Cmd
				m.Input, cmd = m.Input.Update(msg)
				return m, cmd
			}
		}

		// 正常浏览模式
		switch msg.String() {
		case "q", "ctrl+c":
			m.Quitting = true
			// 退出前保存，确保最终状态落地
			_ = m.Save()
			return m, tea.Quit

		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}

		case "down", "j":
			if m.Cursor < len(m.Todos)-1 {
				m.Cursor++
			}

		case "space":
			if len(m.Todos) > 0 {
				m.Todos[m.Cursor].Done = !m.Todos[m.Cursor].Done
				// 切换状态后保存
				_ = m.Save()
			}

		case "a":
			m.Adding = true
			m.Input.Focus()
			return m, textinput.Blink

		case "d":
			if len(m.Todos) > 0 {
				m.Todos = append(m.Todos[:m.Cursor], m.Todos[m.Cursor+1:]...)
				if m.Cursor >= len(m.Todos) {
					m.Cursor = len(m.Todos) - 1
				}
				if m.Cursor < 0 {
					m.Cursor = 0
				}
				// 删除后保存
				_ = m.Save()
			}
		}
	}

	return m, nil
}
