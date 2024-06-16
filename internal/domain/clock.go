package domain

import "time"

type Clock interface {
	Now() time.Time
}

type UTCCLock struct {
}

func (c UTCCLock) Now() time.Time {
	return time.Now().UTC()
}
