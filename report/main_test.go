package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateMessage(t *testing.T) {
	now := time.Now()
	message := createMessage(now)
	expect := fmt.Sprintf(messageFormat, now.Format(timeFormat))
	if message != expect {
		t.Fatal("createMessage Failed.")
	}
}
