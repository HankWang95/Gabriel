package server

import (
	"context"
	"github.com/HankWang95/Gabriel/grpc/proto"
	"github.com/HankWang95/log"
	"google.golang.org/grpc"
	"io"
	"net"
	"os"
	"strconv"
)

type serverGRPC struct {
	server *grpc.Server
	port int
}

func NewListener() *serverGRPC {
	s := new(serverGRPC)
	s.port = 1313

	return s
}

func (s *serverGRPC)Listen()  {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(s.port))
	if err != nil {
		log.Error(err)
		return
	}

	s.server = grpc.NewServer()
	proto.RegisterListenFromPiServer(s.server, s)
	dict := s.server.GetServiceInfo()
	for k,v := range dict{
		log.Debug(k, v)
	}
	err = s.server.Serve(listener)
	if err != nil {
		log.Error(err)
		return
	}
	return
}

// 监听Pi Push 的数据
func (l *serverGRPC) ListenInfo(ctx context.Context, in *proto.InfoMessage) (result *proto.Result, err error) {
	// todo
	return
}


func (l *serverGRPC) Channel(stream proto.ListenFromPi_ChannelServer) error {
	newFile, err := os.Create("./newfile.txt")
	if err != nil {
		log.Error(err)
	}
	for {
		args, err := stream.Recv()
		log.Debug("loop ")
		if err != nil {
			if err == io.EOF {
				break
			}

			// upload file failed, send message to client.
			err := stream.SendAndClose(&proto.UploadStatus{
				Message:"not ok",
				Code:proto.UploadStatusCode_Failed,
			})
			if err != nil {
				log.Error(err)
			}
			return err
		}
		newFile.Write(args.GetValue())

	}


	err = stream.SendAndClose(&proto.UploadStatus{
		Message:"ok",
		Code:proto.UploadStatusCode_Ok,
	})
	if err != nil {
		log.Error(err)
		return err
	}

	newFile.Sync()

	return err
}