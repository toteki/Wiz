all: build separate

build:
	go build ./...

separate:
	go build AES.go
	go build Args.go
	go build ASCII.go
	go build Console.go
	go build Ed25519.go
	go build Files.go
	go build Hash.go
	go build Hex.go
	go build HTTP.go
	go build JSON.go
	go build Random.go
	go build SQLite.go
	go build Strings.go
	go build Time.go
	go build Uint64.go
