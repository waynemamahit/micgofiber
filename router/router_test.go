package router_test

import (
	"micgofiber/lib"
	"micgofiber/router"
	"os"
	"testing"
)

var TestApp *lib.AppConfig

func TestMain(m *testing.M) {
	TestApp = lib.NewApp()
	router.InitApp(TestApp)
	code := m.Run()
	os.Exit(code)
}
