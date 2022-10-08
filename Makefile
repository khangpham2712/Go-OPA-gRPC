go:
	go run cmd/main.go

docker_build:
	docker build -t khangpham2712/dummy:latest .

docker_run:
	docker run -d -p 50051:50051 --name khangpham2712-dummy khangpham2712/dummy:latest

docker_remove_dangling_images:
	docker image rm -f $(docker images --format "{{.ID}}" --filter "dangling=true")

docker_remove_dangling_containers:
	docker container prune

load_test:
	ghz --insecure -c $(curr) -n $(nTime)  \
	--proto ./proto/multiplication.proto \
	--call proto.Multiplication.Multiply \
	-d '{"a":$(a),"b":$(b)}' \
	-m '{"access_token":"$(token)"}' \
	localhost:$(port)

install_ghz:
	brew install ghz
