// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/mrtc0/seccamp-2023/pkg/config"
	"github.com/mrtc0/seccamp-2023/pkg/controllers/book"
	"github.com/mrtc0/seccamp-2023/pkg/controllers/healthcheck"
	"github.com/mrtc0/seccamp-2023/pkg/db"
	"github.com/mrtc0/seccamp-2023/pkg/usecases"
)

// Injectors from wire.go:

func Initialize(conf config.Config) *gin.Engine {
	healthcheckController := healthcheck.NewHealthcheckController()
	gormDB := db.NewDatabaseCon(conf)
	bookRepository := db.NewBookRepository(gormDB)
	listBooks := usecases.NewListBooks(bookRepository)
	bookController := book.NewBookController(listBooks)
	engine := NewServer(conf, healthcheckController, bookController)
	return engine
}

// wire.go:

var APISet = wire.NewSet(healthcheck.NewHealthcheckController, book.NewBookController)

var DBSet = wire.NewSet(db.NewDatabaseCon, db.NewBookRepository)

var UsecaseSet = wire.NewSet(usecases.NewListBooks)
