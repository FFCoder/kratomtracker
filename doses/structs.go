package doses

import "time"

type AddDoseRequest struct {
	DateTaken string `json:"date_taken"`
}

type Dose struct {
	Id        int       `json:"id"`
	DateTaken time.Time `json:"date_taken"`
}
