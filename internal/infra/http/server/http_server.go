package http_server

import (
	"github.com/fikrirnurhidayat/dhasar"
	"github.com/fikrirnurhidayat/ulemulem/internal/domain/invitation"
	"github.com/fikrirnurhidayat/x/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func New() (*dhasar.HTTPServer, error) {
	srv, err := dhasar.NewHTTPServer(&dhasar.HTTPServerOption{
		Container: dhasar.NewContainer(),
		Bootstrap: func(srv *dhasar.HTTPServer) error {
			logger := logger.New("v0.0.0", "v0.0.0")
			postgresDatabaseAdapter := dhasar.NewPostgresDatabaseAdapter(logger)
			db, err := postgresDatabaseAdapter.Connect(&dhasar.PostgresDatabaseAdapterOption{
				Username: viper.GetString("database.username"),
				Password: viper.GetString("database.password"),
				Host:     viper.GetString("database.host"),
				Port:     viper.GetString("database.port"),
				Name:     viper.GetString("database.name"),
				SSLMode:  viper.GetString("database.sslmode"),
			})

			if err != nil {
				return err
			}

			sqlDatabaseManager := dhasar.NewSQLDatabaseManager(logger, db)

			srv.Container.Register("Logger", logger)
			srv.Container.Register("PostgresDatabaseAdapter", postgresDatabaseAdapter)
			srv.Container.Register("SQLDatabaseManager", sqlDatabaseManager)
			srv.Container.Register("Echo", srv.Echo)

			return invitation.WireHTTPModule(srv.Container)
		},
	})

	srv.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    AllowOrigins: []string{"http://localhost:5173"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	srv.Echo.Static("/assets", "/home/fain/Codes/ulemulem/assets")

	return srv, err
}
