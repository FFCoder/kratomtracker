package notificationmanager

type NotificationService interface {
	SendNotification(n NotificationObject) error
}

type NotificationManagerInterface interface {
	Publish(n NotificationObject) error
	AddService(s NotificationService) error
}

type NotificationManager struct {
	repo     NotificationRepository
	services []NotificationService
}

func NewNotificationManager(repo NotificationRepository) *NotificationManager {
	return &NotificationManager{repo: repo}
}

func (manager *NotificationManager) Publish(n NotificationObject) error {
	_, err := manager.repo.Insert(NotificationRecord{Content: n.Content})
	if err != nil {
		return err
	}
	for _, s := range manager.services {
		err := s.SendNotification(n)
		if err != nil {
			return err
		}
	}
	return nil
}

func (manager *NotificationManager) AddService(s NotificationService) error {
	manager.services = append(manager.services, s)
	return nil
}
