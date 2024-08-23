package services

import (
	"fmt"
	"kratomTracker/notificationmanager"
)

type ConsoleNotificationServce struct {
}

func (service *ConsoleNotificationServce) SendNotification(n notificationmanager.NotificationObject) error {
	fmt.Println("Notification: " + n.Content)
	return nil
}
