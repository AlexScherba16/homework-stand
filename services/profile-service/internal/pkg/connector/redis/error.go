package redis

import "fmt"

type PartialError struct {
	SuccessKeys int
	ErrorKeys   int
}

func (e *PartialError) Error() string {
	return fmt.Sprintf("partial-mget: success=%d error=%d", e.SuccessKeys, e.ErrorKeys)
}
