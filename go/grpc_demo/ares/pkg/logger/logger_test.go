package logger

import (
	zapwarpper "ares/logger/zap"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestInitLogger(t *testing.T) {
	initLogger()
	assert.NotNil(t, Log)
}

func TestSetLogger(t *testing.T) {
	l := zapwarpper.New(zap.NewDevelopmentConfig())
	SetLogger(l)
	assert.Equal(t, l, Log)
}
