module github.com/temphia/dev

go 1.16

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/k0kubun/pp v3.0.1+incompatible
	github.com/temphia/core v0.0.0-20220517064252-007019e74ec1
	github.com/temphia/executors v0.0.0-20220517071942-fb2032e38fff
)

replace github.com/temphia/core => ../core

replace github.com/temphia/executors => ../executors
