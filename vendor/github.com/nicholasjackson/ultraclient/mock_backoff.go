package ultraclient

import (
	"time"

	"github.com/stretchr/testify/mock"
)

type MockBackoffStrategy struct {
	mock.Mock
}

func (m *MockBackoffStrategy) Create(retries int, delay time.Duration) []time.Duration {
	args := m.Called(retries, delay)

	return args.Get(0).([]time.Duration)
}
