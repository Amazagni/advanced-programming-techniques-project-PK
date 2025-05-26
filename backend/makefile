compile:
	@go build -o ./bin/backend ./cmd

run:
	@make compile
	./bin/backend backend

generate-swagger:
	swag init --generalInfo cmd/main.go

docker-build:
	docker build -t go-backend .

docker-run: docker-build
	docker run -p 50001:50001 go-backend