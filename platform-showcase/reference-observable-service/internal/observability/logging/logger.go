package logging

import (
	"os"

	"github.com/TheZeroSlave/zapsentry"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger is a common way to set up the logging process. Needs some cleanup.
func NewLogger() (logger *zap.Logger, syncFunc func(), err error) {
	type syncable interface {
		Sync() error
	}

	syncables := make([]syncable, 0)

	logCfg := zap.NewProductionConfig()
	logCfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	logCfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)

	logger, err = logCfg.Build()
	if err != nil {
		return nil, nil, err
	}

	syncables = append(syncables, logger)

	cfg := zapsentry.Configuration{
		Level: zapcore.ErrorLevel,
		Tags: map[string]string{
			"environment": "local",
			"app":         os.Getenv("SERVICE"),
			"version":     "development",
		},
	}

	sentryDSN := os.Getenv("SENTRY_DSN")
	if sentryDSN != "" {
		core, err := zapsentry.NewCore(cfg, zapsentry.NewSentryClientFromDSN(sentryDSN))
		if err != nil {
			return nil, nil, err
		}

		syncables = append(syncables, core)

		logger = zapsentry.AttachCoreToLogger(core, logger.WithOptions(zap.AddStacktrace(zap.ErrorLevel)))
	}

	zap.ReplaceGlobals(logger)

	return logger, func() {
		for _, s := range syncables {
			s.Sync()
		}
	}, nil
}
