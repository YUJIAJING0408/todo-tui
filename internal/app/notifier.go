package app

import (
	"log"

	"github.com/gen2brain/beeep"
)

type Notifier interface {
	Notify(item TodoItem) error
}

// ConsoleNotifier 控制台输出（测试用）
type ConsoleNotifier struct{}

func (n *ConsoleNotifier) Notify(item TodoItem) error {
	log.Printf("【提醒】待办事项到期: %s", item.Text())
	return nil
}

// SystemNotifier 系统弹窗通知（依赖 beeep 库）
type SystemNotifier struct{}

func (n *SystemNotifier) Notify(item TodoItem) error {
	item.Toggle()
	return beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
}
