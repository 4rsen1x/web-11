package main

import (
	"flag"
	"log"
	authAPI "web-11/internal/auth/api"
	authConfig "web-11/internal/auth/config"
	authProvider "web-11/internal/auth/provider"
	authUsecase "web-11/internal/auth/usecase"
)

func main() {
	authConfigPath := flag.String("config-path", "../../configs/hello_example.yaml", "Path to auth config file")
	flag.Parse()
	authCfg, err := authConfig.LoadConfig(*authConfigPath)
	if err != nil {
		log.Fatalf("Failed to load auth config: %v", err)
	}
	authProvider, err := authProvider.NewProvider(authCfg.DB.Host, authCfg.DB.Port, authCfg.DB.User, authCfg.DB.Password, authCfg.DB.DBName)
	if err != nil {
		log.Fatalf("Failed to initialize auth provider: %v", err)
	}
	authUsecase := authUsecase.NewUsecase(authProvider)
	authServer := authAPI.NewServer(authCfg.IP, authCfg.Port, authCfg.SecretKey, authUsecase)
	authServer.Run()
}
