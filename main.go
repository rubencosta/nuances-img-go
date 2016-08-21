package main

import (
	"log"
	"runtime"

	"github.com/disintegration/imaging"
	"github.com/nats-io/nats"
	"github.com/nats-io/nats/encoders/protobuf"
	"github.com/rubencosta/nuances-img-go/proto"
)

func resize(path string, width int, height int) (string, error) {
	img, err := imaging.Open(path)
	if err != nil {
		return "", err
	}
	resizedImg := imaging.Resize(img, width, height, imaging.Lanczos)
	if err := imaging.Save(resizedImg, path); err != nil {
		return "", err
	}
	return path, nil
}

func main() {
	natsc, err := nats.Connect(nats.DefaultURL)
	nc, _ := nats.NewEncodedConn(natsc, protobuf.PROTOBUF_ENCODER)
	if err != nil {
		log.Panic(err)
	}

	nc.Subscribe("image.resize", func(in *imgresizer.Request) {
		log.Printf("message: %s", in)
		if _, err := resize(in.Url, int(in.Width), int(in.Height)); err != nil {
			log.Fatalln(err)
		}
	})

	runtime.Goexit()
}
