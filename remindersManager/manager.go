package remindersManager

import (
	"context"
	"fmt"
	"time"
)

type ReminderManager interface {
	SetReminder(reminder string, time time.Time)
	Start(ctx context.Context)
}

type ReminderManagerImpl struct {
	reminders map[string]time.Time
}

func NewReminderManager() ReminderManager {
	return &ReminderManagerImpl{
		reminders: make(map[string]time.Time),
	}
}

func (r *ReminderManagerImpl) SetReminder(reminder string, time time.Time) {
	r.reminders[reminder] = time
}

func (r *ReminderManagerImpl) Start(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			return
		case <-ticker.C:
			now := time.Now()
			for reminder, tme := range r.reminders {
				if now.After(tme) {
					fmt.Println(reminder)
					delete(r.reminders, reminder)
				}
			}
		}
	}
}
