package repository_test

import (
	"test_go_project/configs"
	"test_go_project/pkg/logger"
	"testing"
)

var testConfig configs.Config

func TestMain(m *testing.M)

func TestSetup(t *testing.T) {
	logger.Debug.Printf("%+v", testConfig)
}
