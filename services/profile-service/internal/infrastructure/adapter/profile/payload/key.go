package payload

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	Prefix = "profile"
)

func ProfileKey(userID int64, version string) string {
	// {version}:{prefix}:{uuid}
	return fmt.Sprintf("%s:%s:%d", version, Prefix, userID)
}

func ProfileID(key string) int64 {
	parts := strings.Split(key, ":")
	if len(parts) != 3 {
		return 0
	}

	userID, _ := strconv.ParseInt(parts[2], 10, 64)
	return userID
}
