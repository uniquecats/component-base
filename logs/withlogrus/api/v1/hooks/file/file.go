package file

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// defaultMaxSize is the default maximum size of a log file in bytes.
const defaultMaxSize = 100 << 20

// program is the base name of a log file.
var program = filepath.Base(os.Args[0])

// logName returns a new log file name containing tag, with start time t, and
// the name for the symlink for tag.
func logName(tag string, t time.Time) (name, link string) {
	name = fmt.Sprintf("%s.%s.%04d%02d%02d-%02d%02d%02d.log",
		program,
		tag,
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
	)
	return name, program + "." + tag
}
