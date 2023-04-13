package levelthree

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/mock"
)

// func Test_GetFile(t *testing.T) {
// 	mfc := mockFileClient{}

// 	err := GetFile(mfc, "some file")
// 	if err != nil {
// 		t.Errorf("error GetFile(), expected no error, got %v", err)
// 	}
// }

type mockFileClient struct {
	mock.Mock
}

func (m mockFileClient) GetObject(ctx context.Context, input *s3.GetObjectInput) (s3.GetObjectOutput, error) {
	args := m.Called(ctx, input)
	resp := args.Get(0).(s3.GetObjectOutput)
	var err error
	errData := args.Get(1)
	if errData != nil {
		err = errData.(error)
	}
	return resp, err
}

func Test_GetFileMock(t *testing.T) {
	mockClient := &mockFileClient{}
	mockClient.On("GetObject", mock.Anything, mock.Anything).Return(s3.GetObjectOutput{}, nil)

	err := GetFile(mockClient, "some file")
	if err != nil {
		t.Errorf("error GetFile(), expected no error, got %v", err)
	}
}
