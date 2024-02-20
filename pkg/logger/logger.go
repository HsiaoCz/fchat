package logger

import "go.uber.org/zap"

type Field = zap.Field

var Logger *zap.Logger

func InitLogger(path string, level string) error {
	return nil
}
