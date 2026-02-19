package config

// LoggerConfig holds logger configuration.
type LoggerConfig struct {
	Activate           string
	SentryDsn          string `json:"-"`
	PerformanceTracing string
	TracesSampleRate   string
}
