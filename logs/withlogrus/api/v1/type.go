package v1

import "github.com/uniquecats/component-base/logs/withlogrus/api/v1/hooks/file"

// LoggingConfiguration contains configuration options for logging.
type LoggingConfiguration struct {
	// Level specifies the minimum log level. Valid values are: debug, info, warn, error, dpanic, panic, and fatal.
	Level string `json:"level,omitempty" mapstructure:"level"`
	// DisableCaller specifies whether to include caller information in the log.
	DisableCaller bool `json:"disable_caller,omitempty" mapstructure:"disable_caller"`
	// DisableStacktrace specifies whether to record a stack trace for all messages at or above panic level.
	// DisableStacktrace bool `json:"disable_stacktrace,omitempty" mapstructure:"disable_stacktrace"`
	DisableStdout     bool `json:"disable_stdout"`
	// Format Flag specifies the structure of log messages.
	// Valid values are: text and json.
	// default value of format is `text`
	Format string `json:"format,omitempty"`

	// FileConfiguration specifies log file
	*file.FileConfiguration
}
