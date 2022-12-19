package sgotipy

import (
	"context"
	"google.golang.org/grpc"
	"os"
	"time"
)

func GetStatus() (*SgotipyStatusResponse, error) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(os.Getenv("SGOTIPY_GRPC_URL"), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	sgotService := NewSgotipyClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	sgotipyStatusResponse, err := sgotService.SgotipyStatus(ctx, &SgotipyStatusRequest{})
	if err != nil {
		return nil, err
	}
	return sgotipyStatusResponse, nil
}

func StopSgotipy() error {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(os.Getenv("SGOTIPY_GRPC_URL"), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	sgotService := NewSgotipyClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = sgotService.StopSgotipy(ctx, &StopSgotipyRequest{})
	if err != nil {
		return err
	}
	return nil
}

func StartSgotipy(auth StartSgotipyRequest) error {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(os.Getenv("SGOTIPY_GRPC_URL"), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	sgotService := NewSgotipyClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = sgotService.StartSgotipy(ctx, &auth)
	if err != nil {
		return err
	}
	return nil
}
