package withlogrus

import (
	logsapi "github.com/uniquecats/component-base/logs/withlogrus/api/v1"
)

// Options is an alias for LoggingConfiguration to comply with component-base
// conventions.
type Options = logsapi.LoggingConfiguration

// NewOptions is an alias for NewLoggingConfiguration.
var NewOptions = logsapi.NewLoggingConfiguration
