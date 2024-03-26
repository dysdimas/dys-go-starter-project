package infrastructures

import (
	"dys-go-starter-project/utils/logger"
	"reflect"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

var (
	Factory = ObjFactory{}
)

type (
	ObjFactory struct {
		obj  map[string]interface{}
		Log  logger.ClientLogger
		Smtp *SmtpConnection
		Env  string
	}

	SmtpConnection struct {
		Host     string
		Port     int64
		Sender   string
		User     string
		Password string
	}
)

type CtxOrDb interface {
	*gin.Context | *xorm.Engine
}

func Get[K any, L CtxOrDb](ctx L) (K, error) {
	name := getTypeName[K]()
	var fn = Factory.obj[name].(func(ctx L) (K, error))
	var o, e = fn(ctx)
	return o, e
}

func Bind[K any, L CtxOrDb](fnInstantiate func(ctx L) (K, error)) {
	name := getTypeName[K]()
	Factory.obj[name] = fnInstantiate
}

func Clear() {
	Factory.obj = make(map[string]interface{})
}

func getTypeName[K any]() string {
	elem := reflect.TypeOf((*K)(nil)).Elem()
	result := elem.Name()
	if result == "" {
		result = elem.Elem().Name()
	}
	return result
}
