package logger

import (
	"go.uber.org/zap"
)

func GetLogger(cnfg *zap.Config) *zap.Logger {
	return zap.Must(cnfg.Build())
}
