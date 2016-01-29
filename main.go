package main

import (
	"log"
	"runtime"

	"github.com/disintegration/imaging"
	"github.com/nats-io/nats"
	"github.com/nats-io/nats/encoders/protobuf"
	"github.com/rubencosta/nuances-img-go/proto"
)

func resize(path string) (string, error) {
	img, err := imaging.Open(path)
	if err != nil {
		return "", err
	}
	resizedImg := imaging.Resize(img, 400, 400, imaging.Lanczos)
	if err := imaging.Save(resizedImg, path); err != nil {
		return "", err
	}
	return path, nil
}

type request struct {
	URL string `json:"url"`
}

func main() {
	natsc, err := nats.Connect(nats.DefaultURL)
	nc, _ := nats.NewEncodedConn(natsc, protobuf.PROTOBUF_ENCODER)
	if err != nil {
		log.Panic(err)
	}

	nc.Subscribe(">", func(in *imgresizer.ImgUrl) {
		log.Printf("message: %s \r\n received on subject:", in)
		if _, err := resize(in.Url); err != nil {
			log.Fatalln(err)
		}
	})

	runtime.Goexit()
}
