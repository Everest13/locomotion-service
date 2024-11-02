package app

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	ls "locomotion-service/internal/app/locomotion-server"
	us "locomotion-service/internal/app/user-server"
	kafkaUser "locomotion-service/internal/kafka/user-create-event"
	locomotionService "locomotion-service/internal/service/locomotion"
	userService "locomotion-service/internal/service/user"
	"locomotion-service/internal/transport/engine"
	"locomotion-service/internal/utility/kafka"
)

type App struct {
	engine                  *engine.Engine
	impl                    []engine.Server
	userService             *userService.Service
	locomotionService       *locomotionService.Service
	database                *pg.DB
	quickStartEventProducer *kafka.Producer
	userCreateConsumer      *kafka.Consumer
	userCreateProcessor     kafka.MessageProcessor
}

func NewApp(ctx context.Context) (*App, error) {
	app := &App{}
	err := app.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) Run(ctx context.Context) error {
	a.runBackground(ctx)

	return a.engine.Run()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initDB,
		a.initKafkaProducer,
		a.initService,
		a.initProcessor,
		a.initKafkaConsumer,
		a.initImpl,
		a.initEngine,
	}

	for _, fn := range inits {
		if err := fn(ctx); err != nil {
			fmt.Println(err.Error())
			return err
		}
	}

	return nil
}

func (a *App) initImpl(_ context.Context) error {
	a.impl = []engine.Server{
		us.NewServer(a.userService),
		ls.NewServer(a.locomotionService),
	}

	return nil
}

func (a *App) initEngine(_ context.Context) error {
	var err error
	a.engine, err = engine.NewEngine(a.impl)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initProcessor(_ context.Context) error {
	a.userCreateProcessor = kafkaUser.NewProcessor(a.userService)

	return nil
}

func (a *App) initKafkaProducer(_ context.Context) error {
	topic := "quickstart-events"
	addr := "localhost:9092"

	var err error
	a.quickStartEventProducer, err = kafka.NewProducer(addr, topic)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initKafkaConsumer(_ context.Context) error {
	addr := "localhost:9092"
	partition := 0
	topic := "user-create"

	var err error
	a.userCreateConsumer, err = kafka.NewConsumer(addr, topic, partition, a.userCreateProcessor)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initService(ctx context.Context) error {
	a.userService = userService.NewService(a.database, a.quickStartEventProducer)

	return nil
}

func (a *App) initDB(ctx context.Context) error {
	db := pg.Connect(&pg.Options{
		Addr: ":5432",
		//User:     "user",
		//Password: "pass",
		Database: "locomotion",
	})

	if err := db.Ping(ctx); err != nil {
		return err
	}

	a.database = db

	//todo add into closer
	//defer db.Close()

	return nil
}

func (a *App) runBackground(ctx context.Context) {
	bgCtx, bgCancel := context.WithCancel(ctx)
	_ = bgCancel //todo add into closer

	go a.userCreateConsumer.Run(bgCtx)
}

// todo вынести старт сервиса
//func (a *App) Run(ctx context.Context) error {
//	flag.Parse()
//	// todo Создаем прослушиватель на порту TCP
//	lis, err := net.Listen("tcp", ":8080")
//	if err != nil {
//		log.Fatalf("failed to listen: %v", err)
//	}
//
//	// NewServer создает сервер gRPC, на котором не зарегистрирована ни одна служба, и который еще не начал принимать запросы.
//	s := grpc.NewServer()
//	server := us.NewServer(a.userService)
//
//	// Attach the Locomotion service to the engine
//	pkgUS.RegisterUserServiceServer(s, server)
//	//log.Printf("engine listening at %v", lis.Addr())
//	// Serve gRPC engine
//	go func() {
//		log.Fatalln(s.Serve(lis))
//	}()
//
//	// NewServer создает сервер gRPC, на котором не зарегистрирована ни одна служба, и который еще не начал принимать запросы.
//	s2 := grpc.NewServer()
//	server2 := ls.NewServer(a.locomotionService)
//
//	// Attach the Locomotion service to the engine
//	pkgLS.RegisterLocomotionServiceServer(s2, server2)
//	//log.Printf("engine listening at %v", lis.Addr())
//	// Serve gRPC engine
//	go func() {
//		log.Fatalln(s.Serve(lis))
//	}()
//
//	// Создаем клиентское подключение к gRPC-серверу, который мы только что запустили
//	// Здесь gRPC-Gateway проксирует запросы
//	conn, err := grpc.NewClient(
//		"0.0.0.0:8080",
//		grpc.WithTransportCredentials(insecure.NewCredentials()),
//	)
//	if err != nil {
//		log.Fatalln("Failed to dial engine:", err)
//	}
//
//	// ServeMux — это мультиплексор запросов для grpc-gateway
//	// Он сопоставляет http-запросы с шаблонами и вызывает соответствующий обработчик
//	gwmux := runtime.NewServeMux()
//	err = desc.RegisterUserServiceHandler(
//		context.Background(),
//		gwmux,
//		conn,
//	)
//
//	// Сервер определяет параметры для запуска HTTP-сервера
//	gwServer := &http.Server{
//		Addr:    ":8090",
//		Handler: gwmux,
//	}
//
//	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
//	// ListenAndServe прослушивает сетевой адрес TCP srv.Addr,
//	// а затем вызывает Serve для обработки запросов на входящих соединениях.
//	log.Fatalln(gwServer.ListenAndServe())
//
//	return err
//}
