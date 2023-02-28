package ulog

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestInfo(t *testing.T) {
	w := httptest.NewRecorder()
	c := GetTestGinContext(w)

	ul := &Log{
		C: c,
	}

	// log info type struct
	type TypeStruct struct {
		Int   int                    `json:"int"`
		Map   map[string]interface{} `json:"map"`
		Slice []string               `json:"slice"`
	}

	logStruct := &TypeStruct{
		Int: 1,
		Map: map[string]interface{}{
			"h": "hello",
			"w": "world",
		},
		Slice: []string{"hello log; 你好，日志; Bonjour journal; Hallo Log; ハローログ"},
	}

	ul.Info(logStruct)

	// log info type string
	ul.Info("hello log; 你好，日志; Bonjour journal; Hallo Log; ハローログ")
}

// mock gin context
func GetTestGinContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}
