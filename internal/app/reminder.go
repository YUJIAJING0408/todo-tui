package app

import (
	"time"

	tea "charm.land/bubbletea/v2"
)

type TickMsg time.Time

func DoTick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}
