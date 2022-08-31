package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/zhanguohuang/go-web-awesome/handler"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	router := gin.Default()
	router.GET("/testing", handler.StartPage)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET",
		"/testing?name=appleboy&address=xyz&birthday=1992-03-15&createTime=1562400033000000123&unixTime=1562400033&number=7",
		nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
