package logger_test

import (
	"testing"

	"github.com/mirzakhany/pkg/logger"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestSetLogLevel(t *testing.T) {
	log := logrus.New()

	err := logger.SetLogLevel(log, "debug")
	assert.Nil(t, err)

	err = logger.SetLogLevel(log, "invalid")
	assert.Equal(t, "not a valid logrus Level: \"invalid\"", err.Error())
}

func TestSetLogOut(t *testing.T) {
	log := logrus.New()

	err := logger.SetLogOut(log, "stdout")
	assert.Nil(t, err)

	err = logger.SetLogOut(log, "stderr")
	assert.Nil(t, err)

	err = logger.SetLogOut(log, "log/access.log")
	assert.Nil(t, err)

	// missing create logs folder.
	err = logger.SetLogOut(log, "logs/access.log")
	assert.NotNil(t, err)
}
