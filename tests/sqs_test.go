package test

import (
	"testing"
)

func TestSQSExample(t *testing.T) {
	testCloudWatchAlarm(t, "sqs")
}
