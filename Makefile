hello:
	echo "Hello Go Api"

build:
	go build -o bin/main main.go

run:
	go run main.go

compile:
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go

docker-run:
	docker-compose down
	docker-compose up -d

docker-build:
	docker-compose down
	docker-compose up -d --build