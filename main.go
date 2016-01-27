package nuancesimg

import (
	"log"
	"net"

	"golang.org/x/net/context"

	pb "github.com/rubencosta/nuances-img-go/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8888"
)

//server is used to implement imgresizer.ImgResizer
type server struct{}

func (s *server) ProcessImg(ctx context.Context, in *pb.ImgUrl) (*pb.ImgUrl, error) {
	return &pb.ImgUrl{Url: in.Url}, nil
}

func main() {
	list, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterImgResizerServer(s, &server{})
	s.Serve(list)
}
