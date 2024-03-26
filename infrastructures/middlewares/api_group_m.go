package middlewares

import (
	"bytes"
	"dys-go-starter-project/infrastructures"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

func ApiGroup(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("apigroup middleware recovered", r)
			infrastructures.Err500ISE(c, fmt.Sprintf("%v", r))
		}
	}()

	apiPath := strings.ReplaceAll(c.Request.URL.Path, "/api/", "/")

	regex, _ := regexp.Compile("/(v\\d+)/")
	apiVersion := regex.FindString(apiPath)
	apiVersion = strings.ReplaceAll(apiVersion, "/", "")
	appToken := c.GetHeader(infrastructures.HeaderAppToken)
	token := c.GetHeader(infrastructures.HeaderToken)

	headers := make(map[string]string)
	headers["app_token"] = appToken
	headers["token"] = token

	var jsonHeader, _ = json.Marshal(headers)
	_ = jsonHeader

	body, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewReader(body))

	//dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", os.Getenv(infrastructures.EnvDbUser), os.Getenv(infrastructures.EnvDbPass), os.Getenv(infrastructures.EnvDbHost), os.Getenv(infrastructures.EnvDbPort), os.Getenv(infrastructures.EnvDbName))
	dsn := fmt.Sprintf("%s:%s@/%s?charset=utf8", os.Getenv(infrastructures.EnvDbUser), os.Getenv(infrastructures.EnvDbPass), os.Getenv(infrastructures.EnvDbName))
	db, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		infrastructures.VErr(err)
	}

	c.Set(infrastructures.CtxPath, apiPath)
	c.Set(infrastructures.CtxApiVersion, apiVersion)
	c.Set(infrastructures.CtxDb, db)
	w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
	c.Writer = w
	c.Next()

	_ = w.body.String()

}
