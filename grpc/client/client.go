package client

import (
	"context"
	"github.com/HankWang95/Gabriel/grpc/proto"
	"github.com/HankWang95/log"
	"google.golang.org/grpc"
	"io"
	"os"
	"time"
)

type ClientGRPC struct {
	client proto.ListenFromPiClient
	chunkSize int
}

type Stats struct {
	StartedAt  time.Time
	FinishedAt time.Time
}

func NewClientGRPC() *ClientGRPC {
	c := &ClientGRPC{}
	conn,err := grpc.Dial("localhost:1313", grpc.WithInsecure())
	if err != nil {
		log.Error(err)
	}

	c.client = proto.NewListenFromPiClient(conn)
	c.chunkSize = 1 << 12
	return c
}

func (c *ClientGRPC) UploadFile(ctx context.Context, f string) (stats Stats, err error) {

	var (
		writing = true
		buf     []byte
		n       int
		file    *os.File
		status  *proto.UploadStatus
	)
	// Get a file handle for the file we
	// want to upload
	file, err = os.Open(f)
	if err != nil {
		log.Error(err)
	}

	// Open a stream-based connection with the
	// gRPC server
	stream, err := c.client.Channel(ctx)
	if err != nil {
		log.Error(err)
	}

	// Start timing the execution
	stats.StartedAt = time.Now()

	// todo 加入tls
	//credentials.NewClientTLSFromFile()

	// Allocate a buffer with `chunkSize` as the capacity
	// and length (making a 0 array of the size of `chunkSize`)
	buf = make([]byte, c.chunkSize)
	for writing {
		// put as many bytes as `chunkSize` into the
		// buf array.
		n, err = file.Read(buf)
		if err != nil {
			if err == io.EOF {
				writing = false
				err = nil
				continue
			}
			log.Error(err)
			// todo 写传输错误处理
			//err = errors.Wrapf(err,
			//	"errored while copying from file to buf")
			return
		}


		stream.Send(&proto.Bytes{
			// because we might've read less than
			// `chunkSize` we want to only send up to
			// `n` (amount of bytes read).
			// note: slicing (`:n`) won't copy the
			// underlying data, so this as fast as taking
			// a "pointer" to the underlying storage.
			Value: buf[:n],
		})
	}

	// keep track of the end time so that we can take the elapsed
	// time later
	stats.FinishedAt = time.Now()

	// close
	status, err = stream.CloseAndRecv()
	if err != nil {
		log.Error(err)
		return
	}

	log.Info(status)

	if status.Code != proto.UploadStatusCode_Ok {
		log.Error("upload fail")
	} else {
		log.Info("upload ok")
	}

	return
}
