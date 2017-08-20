package server

import (
	"fmt"
	"testing"

	context "golang.org/x/net/context"

	"github.com/nicholasjackson/bench/plugin/shared/mocks"
	"github.com/nicholasjackson/bench/server/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mockBench *mocks.MockBench

func setupServer(t *testing.T) *GRPCServer {
	mockBench = &mocks.MockBench{}
	mockBench.On("Do").Return(nil)

	return &GRPCServer{
		benchClient: mockBench,
	}
}

func TestExecuteRetrunsDoError(t *testing.T) {
	s := setupServer(t)
	mockBench.ExpectedCalls = make([]*mock.Call, 0)
	mockBench.On("Do").Return(fmt.Errorf("Boom"))

	_, err := s.Execute(context.Background(), &proto.ExecuteRequest{})

	mockBench.AssertCalled(t, "Do")
	assert.Equal(t, fmt.Errorf("Boom"), err)
}
