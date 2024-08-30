package doses

import (
	"database/sql"
	"fmt"
	"kratomTracker/notificationmanager"
	"time"
)

type DoseRepository interface {
	FindAll() ([]Dose, error)
	FindAllToday() ([]Dose, error)
	GetNextDoseTime() (time.Time, error)
	FindById(id int) (Dose, error)
	Add(dose Dose) (Dose, error)
	Update(dose Dose) (Dose, error)
	Delete(id int) error
}

type SqliteDoseRepository struct {
	Db           *sql.DB
	notifManager notificationmanager.NotificationManagerInterface
}

func NewSqliteDoseRepository(db *sql.DB, notifManager notificationmanager.NotificationManagerInterface) (*SqliteDoseRepository, error) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS doses (id INTEGER PRIMARY KEY, date_taken TEXT)")
	if err != nil {
		return nil, err
	}
	return &SqliteDoseRepository{Db: db, notifManager: notifManager}, nil
}

func formatDate(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}

func parseDate(date string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", date)
}

func (repo *SqliteDoseRepository) Close() error {
	return repo.Db.Close()
}

func (repo *SqliteDoseRepository) FindAll() ([]Dose, error) {
	rows, err := repo.Db.Query("SELECT id, date_taken FROM doses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	doses := []Dose{}
	for rows.Next() {
		var dose Dose
		dateStr := ""
		err := rows.Scan(&dose.Id, &dateStr)
		if err != nil {
			return nil, err
		}
		dose.DateTaken, err = parseDate(dateStr)
		if err != nil {
			return nil, err
		}
		doses = append(doses, dose)
	}

	return doses, nil
}

func (repo *SqliteDoseRepository) FindAllToday() ([]Dose, error) {
	rows, err := repo.Db.Query("SELECT id, date_taken FROM doses WHERE date(date_taken) = date('now', 'localtime')")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	doses := []Dose{}
	for rows.Next() {
		var dose Dose
		tempDateTaken := ""
		err := rows.Scan(&dose.Id, &tempDateTaken)
		if err != nil {
			return nil, err
		}
		dose.DateTaken, err = parseDate(tempDateTaken)
		if err != nil {
			return nil, err
		}
		doses = append(doses, dose)
	}

	return doses, nil
}

func (repo *SqliteDoseRepository) FindById(id int) (Dose, error) {
	var dose Dose
	tempDate := ""
	err := repo.Db.QueryRow("SELECT date_taken FROM doses WHERE id = ?", id).Scan(&tempDate)
	if err != nil {
		return Dose{}, err
	}
	dose.DateTaken, err = parseDate(tempDate)
	return dose, nil
}

func (repo *SqliteDoseRepository) Add(dose Dose) (Dose, error) {
	result, err := repo.Db.Exec("INSERT INTO doses (date_taken) VALUES (?)", formatDate(dose.DateTaken))
	if err != nil {
		fmt.Println(err)
		return Dose{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return Dose{}, err
	}
	dose.Id = int(id)
	notif := notificationmanager.NotificationObject{Content: "A new dose was added!"}
	notifErr := repo.notifManager.Publish(notif)
	if notifErr != nil {
		fmt.Println(notifErr)
	}

	return dose, nil
}

func (repo *SqliteDoseRepository) Update(dose Dose) (Dose, error) {
	_, err := repo.Db.Exec("UPDATE doses SET date_taken = ? WHERE id = ?", formatDate(dose.DateTaken), dose.Id)
	if err != nil {
		return Dose{}, err
	}
	return dose, nil
}

func (repo *SqliteDoseRepository) Delete(id int) error {
	_, err := repo.Db.Exec("DELETE FROM doses WHERE id = ?", id)
	return err
}

func (repo *SqliteDoseRepository) GetNextDoseTime() (time.Time, error) {
	// Get the time 2 hours after the last dose
	var lastDoseTime time.Time
	var lastDoseTimeStr string
	queryStr := `
SELECT COALESCE(
    DATETIME(MAX(date_taken), '+2 hours'),
    DATETIME('now', 'localtime')
) AS result_time
FROM doses
WHERE DATE(date_taken) = DATE('now', 'localtime');
`
	err := repo.Db.QueryRow(queryStr).Scan(&lastDoseTimeStr)
	// If there are no doses today, return the current time
	if err != nil {
		return time.Now(), err
	}
	lastDoseTime, err = parseDate(lastDoseTimeStr)
	if err != nil {
		return time.Now(), err
	}
	return lastDoseTime, nil
}
