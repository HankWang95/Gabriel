package client

import (
	"context"
	"testing"
)

var client *ClientGRPC

func TestNewClientGRPC(t *testing.T) {
	client = NewClientGRPC()
}

func TestClientGRPC_UploadFile(t *testing.T) {
	f := "/Users/hank-for-work/Desktop/go/src/github.com/HankWang95/Gabriel/testfile"
	client.UploadFile(context.Background(), f)
}