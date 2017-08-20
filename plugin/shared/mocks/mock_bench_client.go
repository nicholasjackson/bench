package mocks

import "github.com/stretchr/testify/mock"

// MockBench implements the Bench interface and
// can be used for testing
type MockBench struct {
	mock.Mock
}

// Do implements the interface method Do
func (m *MockBench) Do() error {
	args := m.Called()

	return args.Error(0)
}
