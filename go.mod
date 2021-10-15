module github.com/teramono/engine-backend

go 1.16

replace github.com/teramono/utilities v0.0.0-20210919081101-b247dd3f53c0 => ../utilities

replace github.com/teramono/tera v0.0.0 => ../tera

require (
	github.com/go-playground/validator/v10 v10.9.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/nats-io/nats.go v1.12.3
	github.com/spf13/viper v1.9.0 // indirect
	github.com/teramono/tera v0.0.0
	github.com/teramono/utilities v0.0.0-20210919081101-b247dd3f53c0
	github.com/ugorji/go v1.2.6 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/sys v0.0.0-20210925032602-92d5a993a665 // indirect
	golang.org/x/text v0.3.7 // indirect
)
