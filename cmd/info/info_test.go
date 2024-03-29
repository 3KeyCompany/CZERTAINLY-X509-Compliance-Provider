package info

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"testing"
)

var logger *zap.Logger
var tService Service

func init() {
	logger, _ = zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	tService = NewService(sugar)
}

func TestService_GetInfo(t *testing.T) {
	infoResponse, _ := tService.GetInfo(mux.NewRouter())
	if len(infoResponse) == 0 {
		t.Error("End Points not detected")
	}
}
