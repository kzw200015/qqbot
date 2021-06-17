package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&botLogFormat{})
	logrus.SetLevel(logrus.InfoLevel)
}

type botLogFormat struct{}

func (blf *botLogFormat) Format(entry *logrus.Entry) ([]byte, error) {
	s := fmt.Sprintf("%s [%s] %s\n", entry.Time.Format("2006-01-02 15:04:05"), entry.Level, entry.Message)
	return []byte(s), nil
}
