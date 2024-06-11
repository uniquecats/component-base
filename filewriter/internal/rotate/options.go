package rotate

import (
	"time"
)

type Option func(*Rotate)

// WithRotateInterval sets the interval at which the log file is rotated.
func WithRotateInterval(t time.Duration) Option {
	return func(o *Rotate) {
		o.interval = t
	}
}

// WithRotateSize sets the maximum size of the log file before it is rotated.
func WithRotateSize(size int64) Option {
	return func(o *Rotate) {
		o.size = size
	}
}
