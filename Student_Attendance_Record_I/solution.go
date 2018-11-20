package Student_Attendance_Record_I

import (
	"strings"
)

func checkRecord(s string) bool {
	if len(strings.Split(s, "A")) > 2 {
		return false
	}
	if len(strings.Split(s, "LLL")) > 1{
		return false
	}
	return true
}