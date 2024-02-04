
run: devc.go go.mod
	- @go run devc.go build build

build: devc.go cmd internal go.mod 
	- @go build -o build/devc 