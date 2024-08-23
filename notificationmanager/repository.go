package notificationmanager

import "database/sql"

type NotificationRepository interface {
	FindAll() ([]NotificationRecord, error)
	FindById(id int) (NotificationRecord, error)
	Insert(n NotificationRecord) (NotificationRecord, error)
}

type SqliteNotificationRepository struct {
	Db *sql.DB
}

func NewSqliteNotificationRepository(db *sql.DB) (*SqliteNotificationRepository, error) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS notifications (id INTEGER PRIMARY KEY, date_published TEXT, content TEXT)")
	if err != nil {
		return nil, err
	}
	return &SqliteNotificationRepository{Db: db}, nil
}

func (repo *SqliteNotificationRepository) Close() error {
	return repo.Db.Close()
}

func (repo *SqliteNotificationRepository) FindAll() ([]NotificationRecord, error) {
	rows, err := repo.Db.Query("SELECT id, date_published, content FROM notifications")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	notifications := []NotificationRecord{}
	for rows.Next() {
		var notification NotificationRecord
		err := rows.Scan(&notification.Id, &notification.DatePublished, &notification.Content)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
	}

	return notifications, nil
}

func (repo *SqliteNotificationRepository) FindById(id int) (NotificationRecord, error) {
	var notification NotificationRecord
	err := repo.Db.QueryRow("SELECT date_published, content FROM notifications WHERE id = ?", id).Scan(&notification.DatePublished, &notification.Content)
	if err != nil {
		return NotificationRecord{}, err
	}
	return notification, nil
}

func (repo *SqliteNotificationRepository) Insert(n NotificationRecord) (NotificationRecord, error) {
	result, err := repo.Db.Exec("INSERT INTO notifications (date_published, content) VALUES (?, ?)", n.DatePublished, n.Content)
	if err != nil {
		return NotificationRecord{}, err
	}
	id, _ := result.LastInsertId()
	n.Id = int(id)
	return n, nil
}
