package sensoremu_test

import (
	"fmt"
	"testing"

	"github.com/ekharisma/sensor-emu/sensoremu"
	"github.com/stretchr/testify/assert"
)

func TestTemperatureNotExceed(t *testing.T) {
	testCases := []struct {
		desc     string
		test     float32
		expected bool
	}{
		{
			desc:     "Normal Temperature",
			test:     2,
			expected: true,
		}, {
			desc:     "Exceed Upper Limit",
			test:     12,
			expected: false,
		}, {
			desc:     "Exceed Lower Limit",
			test:     -1,
			expected: false,
		}, {
			desc:     "Normal",
			test:     6,
			expected: true,
		}, {
			desc:     "Normal",
			test:     4,
			expected: true,
		}, {
			desc:     "Normal",
			test:     7,
			expected: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			fmt.Println("Test for ", tC.desc)
			assert.Equal(t, sensoremu.IsDataSafe(tC.test), tC.expected)
		})
	}
}
