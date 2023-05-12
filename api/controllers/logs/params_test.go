package logs

import (
	"testing"
	"time"

	"github.com/equinor/radix-common/utils/pointers"
	logservice "github.com/equinor/radix-log-api/services/logs"
	"github.com/stretchr/testify/assert"
)

func Test_AsLogOptions(t *testing.T) {
	type scenarioSpec struct {
		name     string
		param    logParams
		expected logservice.LogOptions
	}
	start, end := time.Now(), time.Now().Add(1*time.Second)
	scenarios := []scenarioSpec{
		{
			name:     "empty param",
			param:    logParams{},
			expected: logservice.LogOptions{},
		},
		{
			name:     "tail param",
			param:    logParams{Tail: pointers.Ptr(1000)},
			expected: logservice.LogOptions{LimitRows: pointers.Ptr(1000)},
		},
		{
			name:     "start param",
			param:    logParams{timeIntervalParams: timeIntervalParams{Start: &start}},
			expected: logservice.LogOptions{Timeinterval: &logservice.TimeInterval{Start: start}},
		},
		{
			name:     "end param",
			param:    logParams{timeIntervalParams: timeIntervalParams{End: &end}},
			expected: logservice.LogOptions{Timeinterval: &logservice.TimeInterval{End: end}},
		},
	}

	for _, scenario := range scenarios {
		scenario := scenario
		t.Run(scenario.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, scenario.expected, scenario.param.AsLogOptions())
		})
	}
}

func Test_AsComponentPodInventoryOptions(t *testing.T) {
	type scenarioSpec struct {
		name     string
		param    inventoryParams
		expected logservice.ComponentPodInventoryOptions
	}
	start, end := time.Now(), time.Now().Add(1*time.Second)
	scenarios := []scenarioSpec{
		{
			name:     "empty param",
			param:    inventoryParams{},
			expected: logservice.ComponentPodInventoryOptions{},
		},

		{
			name:     "start param",
			param:    inventoryParams{timeIntervalParams: timeIntervalParams{Start: &start}},
			expected: logservice.ComponentPodInventoryOptions{Timeinterval: &logservice.TimeInterval{Start: start}},
		},
		{
			name:     "end param",
			param:    inventoryParams{timeIntervalParams: timeIntervalParams{End: &end}},
			expected: logservice.ComponentPodInventoryOptions{Timeinterval: &logservice.TimeInterval{End: end}},
		},
	}

	for _, scenario := range scenarios {
		scenario := scenario
		t.Run(scenario.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, scenario.expected, scenario.param.AsComponentPodInventoryOptions())
		})
	}
}
