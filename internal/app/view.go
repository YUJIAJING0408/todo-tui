package app

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// View 渲染当前界面
func (m Model) View() tea.View {
	if m.Quitting {
		return tea.NewView("再见！\n")
	}

	var b strings.Builder

	// 标题
	title := TitleStyle.Render("📋 我的待办事项")
	b.WriteString(title + "\n")

	// 待办列表
	if len(m.Todos) == 0 {
		b.WriteString("  没有待办事项，按 'a' 添加一个吧～\n")
	} else {
		for i, todo := range m.Todos {
			// 完成状态
			text := todo.DisplayText()
			textStyle := TodoTextStyle
			if todo.Done() {
				textStyle = DoneTextStyle
			}

			lineStyle := lipgloss.NewStyle()
			if m.Cursor == i {
				lineStyle = CursorLineStyle
			}
			b.WriteString(lineStyle.Render(textStyle.Render(text)) + "\n")
		}
	}

	b.WriteString("\n")

	// 底部帮助 / 添加输入
	if m.Adding {
		b.WriteString("输入新任务: " + m.Input.View() + "\n")
		b.WriteString("(Enter 确认, Esc 取消)\n")
	} else {
		b.WriteString(HelpStyle.Render("↑/↓ 移动 ｜ 空格 切换完成 ｜ a 添加 ｜ d 删除 ｜ q 退出") + "\n")
	}

	return tea.NewView(b.String())
}
