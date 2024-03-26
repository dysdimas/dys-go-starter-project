package infrastructures

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorResponse struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Errors  interface{} `json:"detail,omitempty"`
	Path    string      `json:"path,omitempty"`
}

type SuccessResponse struct {
	Meta interface{} `json:"meta,omitempty"`
	Data interface{} `json:"data,omitempty"`
	Path string      `json:"path,omitempty"`
}

func OkPaging(c *gin.Context, data interface{}, limit int64, offset int64, total int64) {
	meta := make(map[string]interface{}, 3)
	meta["size"] = limit
	meta["page"] = offset
	meta["total"] = total

	c.JSON(http.StatusOK, SuccessResponse{Data: data, Meta: meta, Path: getCtxPath(c)})
}

func Ok(c *gin.Context, data interface{}, meta interface{}) {
	c.JSON(http.StatusOK, SuccessResponse{Data: data, Meta: meta, Path: getCtxPath(c)})
}

func Err500ISE(c *gin.Context, err string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse{Code: 500, Message: err, Path: getCtxPath(c)})
}

func Err500ISEWithDetail(c *gin.Context, err string, detail map[string]interface{}) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse{Code: 500, Message: err, Errors: detail, Path: getCtxPath(c)})
}

func Err400BR(c *gin.Context, err string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{Code: 400, Message: err, Path: getCtxPath(c)})
}

func Err401Unauthorized(c *gin.Context, err string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse{Code: 401, Message: err, Path: getCtxPath(c)})
}

func Err404NF(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, ErrorResponse{Code: 404, Message: "Not Found", Path: getCtxPath(c)})
}

func Err422UE(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, ErrorResponse{Code: 422, Message: "Unprocessable Request", Path: getCtxPath(c)})
}

func getCtxPath(c *gin.Context) string {
	anyPath, exist := c.Get(CtxPath)
	path := ""
	if exist {
		path = anyPath.(string)
	}
	return path
}
