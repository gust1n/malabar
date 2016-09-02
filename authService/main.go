package main

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"sync"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/gust1n/malabar/authService/authservice"
)

// Global user store
var userStore UserStore

type User struct {
	ID       int
	Username string
	Password string
}

type UserStore interface {
	Authenticate(username, password string) (*User, error)
	GetByID(ID int) *User
	GetByUsername(username string) *User
	Insert(username, password string) (*User, error)
}

type inMemoryUserStore struct {
	sync.Mutex
	users     []*User
	currentID int
}

func (us *inMemoryUserStore) Authenticate(username, password string) (*User, error) {
	user := us.GetByUsername(username)
	if user == nil {
		return nil, errors.New("Account not found")
	}
	if user.Password != password {
		return nil, errors.New("Wrong password")
	}

	return user, nil
}

func (us *inMemoryUserStore) GetByID(ID int) *User {
	for _, u := range us.users {
		if u.ID == ID {
			return u
		}
	}
	return nil
}

func (us *inMemoryUserStore) GetByUsername(username string) *User {
	for _, u := range us.users {
		if u.Username == username {
			return u
		}
	}
	return nil
}

func (us *inMemoryUserStore) Insert(username, password string) (*User, error) {
	// Make sure username is unique
	for _, u := range us.users {
		if u.Username == username {
			return nil, errors.New("Username already exists")
		}
	}

	us.Lock()
	us.currentID++
	user := &User{
		ID:       us.currentID,
		Username: username,
		Password: password,
	}
	us.users = append(us.users, user)
	us.Unlock()

	return user, nil
}

type authServiceServer struct{}

func (s *authServiceServer) Authenticate(ctx context.Context, req *pb.AuthReq) (*pb.AuthResp, error) {
	return nil, nil
}

func (s *authServiceServer) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	return nil, nil
}

func main() {
	var (
		caCert     = flag.String("ca-cert", withConfigDir("ca.pem"), "Trusted CA certificate.")
		listenAddr = flag.String("listen-addr", "0.0.0.0:9000", "HTTP listen address.")
		tlsCert    = flag.String("tls-cert", withConfigDir("cert.pem"), "TLS server certificate.")
		tlsKey     = flag.String("tls-key", withConfigDir("key.pem"), "TLS server key.")
		enableTLS  = flag.Bool("tls", true, "Enable TLS")
	)
	flag.Parse()

	log.Println("Auth service starting...")

	userStore = &inMemoryUserStore{}

	var opts []grpc.ServerOption
	if *enableTLS {
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
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	gs := grpc.NewServer(opts...)
	pb.RegisterAuthServiceServer(gs, &authServiceServer{})

	lis, err := net.Listen("tcp", *listenAddr)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Auth service started successfully")
	log.Fatal(gs.Serve(lis))
}

func withConfigDir(path string) string {
	return filepath.Join(os.Getenv("HOME"), ".authService", "server", path)
}
