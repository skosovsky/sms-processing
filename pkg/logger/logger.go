package logger

import (
	"os"
)

const (
	defaultLevel          = LevelInfo
	defaultAddSource      = false
	defaultIsJSON         = false
	defaultsUseMiddleware = false
	defaultSetDefault     = false
)

func NewLogger(opts ...Option) *Logger {
	config := &Options{
		Level:         defaultLevel,
		AddSource:     defaultAddSource,
		IsJSON:        defaultIsJSON,
		UseMiddleware: defaultsUseMiddleware,
		SetDefault:    defaultSetDefault,
	}

	for _, opt := range opts {
		opt(config)
	}

	options := &HandlerOptions{
		AddSource:   config.AddSource,
		Level:       config.Level,
		ReplaceAttr: nil,
	}

	var handler Handler = NewTextHandler(os.Stdout, options)

	if config.IsJSON {
		handler = NewJSONHandler(os.Stdout, options)
	}

	if config.UseMiddleware {
		handler = NewHandlerMiddleware(handler)
	}

	logger := New(handler)

	if config.SetDefault {
		SetDefault(logger)
	}

	return logger
}

type Options struct {
	Level         Level
	AddSource     bool
	IsJSON        bool
	UseMiddleware bool
	SetDefault    bool
}

type Option func(*Options)

// WithLevel logger option sets the log level, if not set, the default level is defaultLevel.
func WithLevel(level string) Option {
	return func(opts *Options) {
		var l Level
		if err := l.UnmarshalText([]byte(level)); err != nil {
			l = LevelInfo
		}

		opts.Level = l
	}
}

// WithAddSource logger option sets the add source option, which will add source file and line number to the log record.
func WithAddSource(addSource bool) Option {
	return func(opts *Options) {
		opts.AddSource = addSource
	}
}

// WithMiddleware logger option sets the usage middleware.
func WithMiddleware(useMiddleware bool) Option {
	return func(opts *Options) {
		opts.UseMiddleware = useMiddleware
	}
}

// WithIsJSON logger option sets the is json option, which will set JSON format for the log record.
func WithIsJSON(isJSON bool) Option {
	return func(opts *Options) {
		opts.IsJSON = isJSON
	}
}

// WithSetDefault logger option sets the set default option, which will set the created logger as default logger.
func WithSetDefault(setDefault bool) Option {
	return func(opts *Options) {
		opts.SetDefault = setDefault
	}
}

// WithAttrs returns logger with attributes.
func WithAttrs(logger *Logger, attrs ...Attr) *Logger {
	for _, attr := range attrs {
		logger = logger.With(attr)
	}

	return logger
}
