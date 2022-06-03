package test

import (
	"testing"
)

func TestStateMachineExample(t *testing.T) {
	testCloudWatchAlarm(t, "state-machine")
}
