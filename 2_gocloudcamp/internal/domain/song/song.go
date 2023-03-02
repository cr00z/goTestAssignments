package song

import "time"

type Song struct {
	Id       uint64
	Name     string
	Duration time.Duration
}
