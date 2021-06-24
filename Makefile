docker:
	docker build -t slandow/pingpong:latest .

docker.push: docker
	docker push slandow/pingpong:latest

server:
	go run ./cmd/server

client:
	go run ./cmd/client
