package postgres

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type configTester struct {
	captured map[string]string
}

func (c configTester) captureConfig(key, value string) {
	c.captured[key] = value
}

func TestConfig(t *testing.T) {
	tt := []struct {
		name       string
		configFunc func(configBuilder)
		expected   map[string]string
	}{
		{
			name:       "credentials",
			configFunc: Credentials("someuser", "somepass"),
			expected: map[string]string{
				"user":     "someuser",
				"password": "somepass",
			},
		},
		{
			name:       "host and port",
			configFunc: HostAndPort("somehost", 1234),
			expected: map[string]string{
				"host": "somehost",
				"port": "1234",
			},
		},
		{
			name:       "dbname",
			configFunc: DbName("somedb"),
			expected: map[string]string{
				"dbname": "somedb",
			},
		},
		{
			name:       "ssl disabled",
			configFunc: SslDisabled(),
			expected: map[string]string{
				"sslmode": "disable",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tester := configTester{make(map[string]string)}
			tc.configFunc(tester.captureConfig)

			assert.Equal(t, tc.expected, tester.captured)
		})
	}
}
