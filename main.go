package main

import (
	inf "dys-go-starter-project/infrastructures"
	"dys-go-starter-project/infrastructures/middlewares"
	"dys-go-starter-project/modules/auth/repositories"
	authUserService "dys-go-starter-project/modules/auth/services"
	userRepo "dys-go-starter-project/modules/user/repositories"
	userService "dys-go-starter-project/modules/user/services"
	"dys-go-starter-project/routes"
	"dys-go-starter-project/utils/logger"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"xorm.io/xorm"
)

func main() {
	defer deferMain()
	env, err := loadEnvironment()
	if err != nil {
		panic(err)
	}

	registerLogger(env)
	registerServices()
	apiEngine := gin.Default()

	middlewares.Attach(apiEngine)
	routes.Attach(apiEngine)

	err = http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv(inf.EnvAppPort)), apiEngine)
	fmt.Println(err)
}

func deferMain() {
	if r := recover(); r != nil {
		var resultErr error

		switch x := r.(type) {
		case string:
			resultErr = errors.New(x)
		case error:
			resultErr = x
		default:
			resultErr = errors.New("unknown panic")
		}

		fmt.Println(resultErr)
	} else {
		fmt.Println("application running perfectly")
	}
}

func registerServices() {
	inf.Clear()

	// Auth
	inf.Bind[*authUserService.AuthUserService](func(ctx *gin.Context) (*authUserService.AuthUserService, error) {
		db, exist := ctx.Get(inf.CtxDb)
		if !exist {
			return nil, errors.New("db engine does not exist")
		}

		return authUserService.NewAuthService(db.(*xorm.Engine)), nil
	})

	inf.Bind[repositories.AuthUserRepository](func(db *xorm.Engine) (repositories.AuthUserRepository, error) {
		return repositories.NewAuthUserRepositoryImpl(db), nil
	})

	// User
	inf.Bind[*userService.UserService](func(ctx *gin.Context) (*userService.UserService, error) {
		db, exist := ctx.Get(inf.CtxDb)
		if !exist {
			return nil, errors.New("db engine does not exist")
		}

		return userService.NewUserService(db.(*xorm.Engine)), nil
	})

	inf.Bind[userRepo.UserRepository](func(db *xorm.Engine) (userRepo.UserRepository, error) {
		return userRepo.NewUserRepositoryImpl(db), nil
	})
}

func loadEnvironment() (string, error) {
	envFound := false

	env := os.Getenv(inf.EnvSystemEnvironment)
	if env == "" {
		env = "development"
	}

	err := godotenv.Load(".env." + env + ".local")
	if err == nil {
		envFound = true
	}

	if env == "test" {
		err = godotenv.Load(".env.local")
		if err == nil {
			envFound = true
		}
	}

	err = godotenv.Load(".env." + env)
	if err == nil {
		envFound = true
	}

	if !envFound {
		err = errors.New("environment file not found")
		return env, err
	}

	inf.Factory.Env = env
	return env, nil
}

func registerLogger(env string) {
	fmt.Println(os.Getenv(inf.EnvAppLogPath))
	inf.Factory.Log = &logger.ClientLogrus{
		Config: &logger.ClientLogrusConfig{
			MaxFileSizeMb:  50,
			PrefixFileName: inf.PrefixFileName,
			MaxAgeDay:      30,
			Compress:       true,
			LogDirPath:     os.Getenv(inf.EnvAppLogPath),
		},
	}

	level := logger.DebugLevel
	if env == "staging" {
		level = logger.InfoLevel
	} else if env == "production" {
		level = logger.WarnLevel
	}

	err := inf.Factory.Log.Init(level)
	if err != nil {
		return
	}
}
