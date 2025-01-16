package tools

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// GenerateTimestampedUUID generates a unique UUID that is prefixed with the current timestamp.
// The timestamp is formatted as "DDMMYYYY-HHMMSS" followed by a standard UUID, which ensures that
// the generated string is both time-specific and globally unique.
//
// The generated string combines the formatted current date and time with a UUID, separated by a hyphen.
//
// Returns:
//   - A string representing a timestamped UUID, formatted as "DDMMYYYY-HHMMSS-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx".
//
// Example:
//   For example, calling this function at "2025-01-17 15:30:45" would generate a string like:
//   "17012025-153045-123e4567-e89b-12d3-a456-426614174000".
func GenerateTimestampedUUID() string {
	dt := time.Now()
	formattedTime := dt.Format("02012006-150405")
	guid := uuid.New()

	return fmt.Sprintf("%s-%s", formattedTime, guid)
}
