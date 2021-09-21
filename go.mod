module github.com/teramono/engine-backend

go 1.16

require (
	github.com/gin-gonic/gin v1.7.4
	github.com/teramono/utilities v0.0.0-20210919081101-b247dd3f53c0
	gorm.io/driver/sqlite v1.1.5 // indirect
	rogchap.com/v8go v0.6.0
)

replace github.com/teramono/utilities v0.0.0-20210919081101-b247dd3f53c0 => ../utilities
