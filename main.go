package main

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/YUJIAJING0408/todo-tui/internal/app"
)

func main() {
	p := tea.NewProgram(app.NewModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("运行出错: %v\n", err)
	}
}
