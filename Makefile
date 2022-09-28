go:
	go run cmd/main.go
docker_build:
	docker build -t khangpham2712/dummy:latest .
docker_run:
	docker run -d -p 50000:50000 khangpham2712/dummy:latest
