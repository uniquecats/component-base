package rotate

import "time"

type Rotate struct {
	interval time.Duration
	size     int64
}

func New() *Rotate {
	return &Rotate{}
}
