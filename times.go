package godash

import (
	"strconv"
	"strings"
	"time"
)

// FormatDuration formats a time range into a specified expression, such as MM:DD HH:mm:ss.
func FormatDuration(d time.Duration, formatStr string) string {
	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60
	milliseconds := int(d.Milliseconds()) % 1000

	replacements := map[string]string{
		"YYYY": "0000",
		"MM":   "00",
		"DD":   "00",
		"HH":   PadStart(strconv.Itoa(hours), 2, "0"),
		"mm":   PadStart(strconv.Itoa(minutes), 2, "0"),
		"ss":   PadStart(strconv.Itoa(seconds), 2, "0"),
		"SS":   PadStart(strconv.Itoa(milliseconds), 3, "0"),
	}

	for key, value := range replacements {
		formatStr = strings.ReplaceAll(formatStr, key, value)
	}
	return formatStr
}
