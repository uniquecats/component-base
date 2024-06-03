package file

import (
	"github.com/sirupsen/logrus"
)

type Hook struct{}

func (h *Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (fh *Hook) Fire(e *logrus.Entry) error {
	var err error
	return err
}
