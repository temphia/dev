package main

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"github.com/temphia/core/backend/plane"
	"github.com/temphia/core/backend/server/app"
	"github.com/temphia/core/backend/server/btypes"
	"github.com/temphia/core/backend/server/btypes/rtypes"
	"github.com/temphia/core/backend/server/registry"

	// db devndors
	_ "github.com/temphia/core/backend/server/store/upper/vendors/postgres"
	_ "github.com/temphia/core/backend/server/store/upper/vendors/ql"
	_ "github.com/temphia/core/backend/server/store/upper/vendors/sqlite"

	// blob providers
	_ "github.com/temphia/core/backend/server/store/cabinet/native"

	// repo provider
	_ "github.com/temphia/core/backend/server/services/pacman/providers/embed"
	_ "github.com/temphia/core/backend/server/services/pacman/providers/gitlab"
	_ "github.com/temphia/core/backend/server/services/pacman/providers/local"

	"github.com/temphia/core/backend/server/engine/executors/goja"
	_ "github.com/temphia/core/backend/server/engine/executors/wasmer2"
	_ "github.com/temphia/executors/backend/dashed"
	_ "github.com/temphia/executors/backend/wizard"
)

func init() {
	os.Chdir("../core")
	pp.Println("$PWD")
	pp.Println(os.Getwd())
	registry.SetExecutor("goja", rtypes.BuilderFunc(goja.NewBuilder))
}

func main() {
	builder := app.NewBuilder(
		app.WithDevConfig(),
		app.WithControlPlane(plane.NewLite()),
		app.WithBeforeListen(devDebug),
	)

	err := builder.Build()
	if err != nil {
		panic(err)
	}

	err = builder.RunApp()
	if err != nil {
		panic(err)
	}

}

func devDebug(a btypes.App, e *gin.Engine) error {
	go func() {
		time.Sleep(time.Second * 5)
		cdb := a.CoreHub()
		pp.Println(cdb.GetUserGroup("default1", "operator"))

	}()

	return nil

}
