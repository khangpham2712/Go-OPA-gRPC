go_run:
	go run cmd/main.go

docker_build:
	docker build -t khangpham2712/dummy:latest .

docker_run:
	docker run -d -p 50051:50051 --name khangpham2712-dummy khangpham2712/dummy:latest

docker_push:
	docker push khangpham2712/dummy:latest

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

get_port:
	kubectl get pod $(podName) --template='{{(index (index .spec.containers 0).ports 0).containerPort}}{{"\n"}}'

port_forward:
	kubectl port-forward service/$(svcName) $(hostPort):$(kubectl get pod $(kubectl get pods) --template='{{(index (index .spec.containers 0).ports 0).containerPort}}{{"\n"}}')

minikube_start:
	minikube start
	kubectl create -f $(path)

minikube_end:
	minikube stop
	minikube delete
