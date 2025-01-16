package tools

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func GenerateTimestampedUUID() string {
	dt := time.Now()
	formattedTime := dt.Format("02012006-150405")
	guid := uuid.New()

	return fmt.Sprintf("%s-%s", formattedTime, guid)
}
