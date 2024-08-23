package notificationmanager

type NotificationObject struct {
	Content string `json:"content"`
}

type NotificationRecord struct {
	Id            int    `json:"id"`
	DatePublished string `json:"date_published"`
	Content       string `json:"content"`
}
