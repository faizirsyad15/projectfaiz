package main

import (
	"fmt"

	ep "projectHotel/hotel/kamar/endpoint"
	pb "projectHotel/hotel/kamar/grpc"
	svc "projectHotel/hotel/kamar/server"

	cfg "projectHotel/hotel/util/config"
	run "projectHotel/hotel/util/grpc"
	util "projectHotel/hotel/util/microservice"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	logger := util.Logger()

	ok := cfg.AppConfig.LoadConfig()
	if !ok {
		logger.Log(util.LogError, "failed to load configuration")
		return
	}

	discHost := cfg.GetA("discoveryhost", "127.0.0.1:2181")
	ip := cfg.Get("serviceip", "127.0.0.1")
	port := cfg.Get("serviceport", "7001")
	address := fmt.Sprintf("%s:%s", ip, port)

	registrar, err := util.ServiceRegistry(discHost, svc.ServiceID, address, logger)
	if err != nil {
		logger.Log(util.LogError, "cannot find registrar")
		return
	}
	registrar.Register()
	defer registrar.Deregister()

	tracerHost := cfg.Get("tracerhost", "127.0.0.1:9999")
	tracer := util.Tracer(tracerHost)

	var server pb.KamarServiceServer
	{
		//chHost := cfg.Get("chhost", "127.0.0.1:6379")
		//cacher := svc.NewRedisCache(chHost)

		//gmapKey := cfg.Get("gmapkey", "AIzaSyD9tm3UVfxRWeaOy_MQ7tsCj1fVCLfG8Bo")
		//locator := svc.NewLocator(gmapKey)

		dbHost := cfg.Get(cfg.DBhost, "127.0.0.1:3306")
		dbName := cfg.Get(cfg.DBname, "db_hotel")
		dbUser := cfg.Get(cfg.DBuid, "root")
		dbPwd := cfg.Get(cfg.DBpwd, "root")

		//brokers := cfg.GetA("mqbrokers", "127.0.0.1:9092")

		// sebelum ke coding dibawah ini
		// selesaikan di service.go terlebih dahulu

		dbReadWriter := svc.NewDBReadWriter(dbHost, dbName, dbUser, dbPwd)
		//dbRuler := svc.NewDBDispatchRuler(dbReadWriter, locator)
		//notifier := mq.NewAsyncProducer(brokers, nil)

		//auctioneer := svc.NewAuctioneer(dbReadWriter, cacher)
		service := svc.NewKamar(dbReadWriter)
		endpoint := ep.NewKamarEndpoint(service)
		fmt.Println(endpoint)
		server = ep.NewGRPCKamarServer(endpoint, tracer, logger)
	}

	// ca := cfg.Get("capath", "cert/cacert.pem")
	// cert := cfg.Get("certpath", "cert/server.pem")
	// prikey := cfg.Get("keypath", "cert/server.key")
	//
	// tls, err := run.TLSCredentialFromFile(ca, cert, prikey, true)
	// if err != nil {
	// 	logger.Log("tlserr", err)
	// 	return
	// }
	//grpcServer := grpc.NewServer(append(run.Recovery(logger), grpc.Creds(tls))...)
	grpcServer := grpc.NewServer(run.Recovery(logger)...)
	pb.RegisterKamarServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	exit := make(chan bool, 1)
	go run.Serve(address, grpcServer, exit, logger)

	<-exit
}
