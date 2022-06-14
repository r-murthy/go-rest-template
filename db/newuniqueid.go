package db

import (
	"fmt"
	"strconv"
	"strings"
)

// GenerateID creates string by incrementing latestID by 1 if latestID exists, else starts the pretext with 1
func GenerateID(latestID string, pretext string, uniqueIdformat string) (newID string, err error) {
	var newIDNum int
	if len(latestID) == 0 {
		newIDNum = 1
	} else {
		latestIDParts := strings.Split(latestID, pretext)
		latestIDNum, err := strconv.Atoi(latestIDParts[1])
		if err != nil {
			return newID, err
		}
		newIDNum = latestIDNum + 1
	}

	newID = fmt.Sprintf(pretext+uniqueIdformat, newIDNum)
	return
}
