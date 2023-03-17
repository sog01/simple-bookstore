upES:
	docker-compose up -d
downES:
	docker-compose down
build:
	go build -o cmd/app/main cmd/app/main.go
run: build
	./cmd/app/main