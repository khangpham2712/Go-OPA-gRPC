go:
	go run cmd/main.go

docker_build:
	docker build -t khangpham2712/dummy:latest .

docker_run:
	docker run -d -p 50000:50000 khangpham2712/dummy:latest

load_test:
	ghz --insecure -c $(curr) -n $(nTime)  \
	--proto ./proto/multiplication.proto \
	--call proto.Multiplication.Multiply \
	-d '{"a":$(a),"b":$(b)}' \
	-m '{"access_token":"$(token)"}' \
	localhost:$(port)

install_ghz:
	brew install ghz
