
run: main.go go.mod
	@go run .

build: main.go cmd internal go.mod 
	@go build -o build/devc