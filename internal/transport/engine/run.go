package engine

import (
	"context"
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	desc "locomotion-service/pkg/user-service"
	"log"
	"net"
	"net/http"
)

type Server interface {
	RegisterServer(s *grpc.Server)
}

type Engine struct {
	servers []Server
	opts    Opts
}

func NewEngine(impl []Server) (*Engine, error) {
	return &Engine{servers: impl}, nil
}

func (e *Engine) Run() error {
	flag.Parse()

	err := e.runGRPCServer()
	if err != nil {
		return err
	}

	err = e.runPublicHTTP()
	if err != nil {
		return err
	}

	return nil
}

func (e *Engine) runGRPCServer() error {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}

	s := grpc.NewServer()
	for _, server := range e.servers {
		server.RegisterServer(s)
	}

	// Serve gRPC engine
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	return nil
}

// todo возможно переписатть сервер http
func (e *Engine) runPublicHTTP() error {
	// Создаем клиентское подключение к gRPC-серверу, который мы только что запустили
	// Здесь gRPC-Gateway проксирует запросы
	conn, err := grpc.NewClient(
		"0.0.0.0:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial engine:", err)
		return err
	}

	// ServeMux — это мультиплексор запросов для grpc-gateway
	// Он сопоставляет http-запросы с шаблонами и вызывает соответствующий обработчик
	gwmux := runtime.NewServeMux()
	err = desc.RegisterUserServiceHandler(
		context.Background(),
		gwmux,
		conn,
	)
	if err != nil {
		return err
	}

	// Сервер определяет параметры для запуска HTTP-сервера
	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")

	err = gwServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
