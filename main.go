package main

import (
	"log"
	"net"

	"golang.org/x/net/context"

	"github.com/disintegration/imaging"
	pb "github.com/rubencosta/nuances-img-go/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8888"
)

type server struct{}

func (s *server) ProcessImg(ctx context.Context, in *pb.ImgUrl) (*pb.ImgUrl, error) {
	log.Printf("received %v", in.Url)
	img, err := imaging.Open(in.Url)
	if err != nil {
		return nil, err
	}
	resizedImg := imaging.Resize(img, 400, 400, imaging.Lanczos)
	newImgPath := in.Url
	if err := imaging.Save(resizedImg, newImgPath); err != nil {
		return nil, err
	}
	log.Printf("return %v", newImgPath)
	return &pb.ImgUrl{Url: newImgPath}, nil
}

func main() {
	list, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Printf("listening on port %v", port)
	}
	s := grpc.NewServer()
	pb.RegisterImgResizerServer(s, &server{})
	s.Serve(list)
}
