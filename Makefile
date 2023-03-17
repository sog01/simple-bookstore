upES:
	docker-compose up -d
populateData:
	curl -s -H "Content-Type: application/x-ndjson" -XPOST localhost:9200/_bulk --data-binary "@populate-books"
downES:
	docker-compose down
build:
	go build -o cmd/app/main cmd/app/main.go
run: build
	./cmd/app/main