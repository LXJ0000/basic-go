// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"webook-server/internal/handler"
	"webook-server/internal/repository"
	"webook-server/internal/repository/cache"
	"webook-server/internal/repository/dao"
	"webook-server/ioc"
)

// Injectors from wire.go:

func InitWebServer() *gin.Engine {
	v := ioc.InitGinMiddlewares()
	db := ioc.InitDB()
	userDao := dao.NewUserDao(db)
	cmdable := ioc.InitRedis()
	userCache := cache.NewUserCache(cmdable)
	userRepository := repository.NewUserRepository(userDao, userCache)
	codeCache := cache.NewCodeCache(cmdable)
	codeRepository := repository.NewCodeRepository(codeCache)
	service := ioc.InitSMSService()
	userHandler := handler.NewUserHandler(userRepository, codeRepository, service)
	engine := ioc.InitWebServer(v, userHandler)
	return engine
}
