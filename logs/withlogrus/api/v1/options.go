package v1

import (
	"time"

	"github.com/spf13/pflag"
)

// NewLoggingConfiguration returns a struct holding the default logging configuration.
func NewLoggingConfiguration() *LoggingConfiguration {
	c := LoggingConfiguration{}
	SetRecommendedLoggingConfiguration(&c)
	return &c
	// return &LoggingConfiguration{
	// 	DisableCaller:     false,
	// 	DisableStacktrace: false,
	// 	DisableStdout:     false,
	// 	Level:             logrus.InfoLevel.String(),
	// 	Format:            "text",
	// 	FileConfiguration: nil,
	// }
}

// SetRecommendedLoggingConfiguration sets the default logging configuration
// for fields that are unset.
//
// Consumers who embed LoggingConfiguration in their own configuration structs
// may set custom defaults and then should call this function to add the
// global defaults.
func SetRecommendedLoggingConfiguration(c *LoggingConfiguration) {
	if c.Format == "" {
		c.Format = "text"
	}
}

// flagSet is the interface implemented by pflag.FlagSet, with
// just those methods defined which are needed by addFlags.
type flagSet interface {
	BoolVar(p *bool, name string, value bool, usage string)
	DurationVar(p *time.Duration, name string, value time.Duration, usage string)
	StringVar(p *string, name string, value string, usage string)
	Var(value pflag.Value, name string, usage string)
	VarP(value pflag.Value, name, shorthand, usage string)
}

var _ flagSet = &pflag.FlagSet{}
