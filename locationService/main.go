package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/gust1n/microServicesArchitecture/locationService/locationservice"
)

type locationServiceServer struct{}

func (lss *locationServiceServer) TrackUser(user *pb.User, stream pb.LocationService_TrackUserServer) error {
	for i := 0; i < 100; i++ {
		if err := stream.Send(&pb.Point{
			Latitude: 1,
			Longitude: 0,
		}); err != nil {
			return err
		}
	}
	return nil
}

func withConfigDir(path string) string {
	return filepath.Join(os.Getenv("HOME"), ".locationService", "server", path)
}

func main() {
	var (
		caCert     = flag.String("ca-cert", withConfigDir("ca.pem"), "Trusted CA certificate.")
		listenAddr = flag.String("listen-addr", "0.0.0.0:9000", "HTTP listen address.")
		tlsCert    = flag.String("tls-cert", withConfigDir("cert.pem"), "TLS server certificate.")
		tlsKey     = flag.String("tls-key", withConfigDir("key.pem"), "TLS server key.")
	)
	flag.Parse()

	log.Println("Location service starting...")

	cert, err := tls.LoadX509KeyPair(*tlsCert, *tlsKey)
	if err != nil {
		log.Fatal(err)
	}

	rawCaCert, err := ioutil.ReadFile(*caCert)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(rawCaCert)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    caCertPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	})

	gs := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterLocationServiceServer(gs, &locationServiceServer{})

	lis, err := net.Listen("tcp", *listenAddr)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Location service started successfully")
	log.Fatal(gs.Serve(lis))
}
