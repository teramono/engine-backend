module github.com/teramono/engine-backend

go 1.16

replace github.com/teramono/utilities v0.0.0-20210919081101-b247dd3f53c0 => ../utilities

replace github.com/teramono/engine-fs v0.0.0-20210924140556-e15c34e7dbcd => ../engine-fs

replace github.com/teramono/engine-db v0.0.0-20210924140608-36967c4678af => ../engine-db

require (
	github.com/gin-gonic/gin v1.7.4
	github.com/go-playground/validator/v10 v10.9.0 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/spf13/viper v1.9.0 // indirect
	github.com/teramono/engine-db v0.0.0-20210924140608-36967c4678af
	github.com/teramono/engine-fs v0.0.0-20210924140556-e15c34e7dbcd
	github.com/teramono/utilities v0.0.0-20210919081101-b247dd3f53c0
	github.com/ugorji/go v1.2.6 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/sys v0.0.0-20210925032602-92d5a993a665 // indirect
	golang.org/x/text v0.3.7 // indirect
	rogchap.com/v8go v0.6.0
)
