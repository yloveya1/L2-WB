package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNtpTime(t *testing.T) {
	ntpTime, err := GetTime()
	realTime := time.Now()
	tt := ntpTime.Sub(realTime)
	assert.NoError(t, err)
	assert.Less(t, tt, 5*time.Second)
}
