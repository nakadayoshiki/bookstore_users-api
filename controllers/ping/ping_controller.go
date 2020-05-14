package ping

import (
    "github.com/gin-gonic/gin"
    "net/http"
)
func ping(c *gin.Context){
    c.String(http.StatusOK, "pong")
}