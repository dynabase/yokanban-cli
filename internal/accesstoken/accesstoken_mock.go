package accesstoken

import "github.com/stretchr/testify/mock"

// Mock is a mock for the AccessToken.
type Mock struct {
	mock.Mock
}

// Get mock for the original method.
func (m *Mock) Get() string {
	args := m.Called()
	return args.String(0)
}

// Refresh mock for the original method.
func (m *Mock) Refresh() string {
	args := m.Called()
	return args.String(0)
}
