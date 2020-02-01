package db

import (
	"github.com/gofrs/uuid"
	"time"
)

func scanDate(date string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05Z", date)
}

func createUUID() string {
	u, _ := uuid.NewV4()
	return u.String()
}
