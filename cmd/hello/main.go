package main

import (
	"flag"
	"log"
	authAPI "web-11/internal/auth/api"
	authConfig "web-11/internal/auth/config"
	"web-11/internal/hello/api"
	"web-11/internal/hello/config"
	"web-11/internal/hello/provider"
	"web-11/internal/hello/usecase"

	_ "github.com/lib/pq"
)

func main() {
	// Считываем аргументы командной строки
	configPath := flag.String("config-path", "../../configs/hello_example.yaml", "путь к файлу конфигурации")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}
	authConfigPath := flag.String("auth-config", "../../configs/hello_example.yaml", "Path to auth config file")
	authCfg, err := authConfig.LoadConfig(*authConfigPath)
	if err != nil {
		log.Fatalf("Failed to load auth config: %v", err)
	}
	authMiddleware := authAPI.JWTMiddleware(authCfg.SecretKey)
	prv := provider.NewProvider(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBname)
	use := usecase.NewUsecase(cfg.Usecase.DefaultMessage, prv)
	srv := api.NewServer(cfg.IP, cfg.Port, cfg.API.MaxMessageSize, use)
	srv.AddMiddleware(authMiddleware)
	srv.Run()
}