.PHONY: build create delete push deploy

IMAGE=kaspernissen/hello-go
TAG=latest
ARCH=arm
OS=linux
IP=192.168.1.100
PORT=$(shell kubectl get svc hello  -o json | jq '.spec.ports[].nodePort')

build:
	CGO_ENABLED=0 GOOS=$(OS) GOARCH=$(ARCH) go build hello.go
	docker build -t $(IMAGE):$(TAG) .
	@rm -f hello

push:
	docker push $(IMAGE):$(TAG)

delete:
	kubectl delete -f kubernetes/deployment.yaml
	kubectl delete -f kubernetes/svc.yaml

deploy:
	kubectl apply -f kubernetes/deployment.yaml
	kubectl apply -f kubernetes/svc.yaml

ping:
	while sleep 0.5; do curl http://$(IP):$(PORT)/ping; done

version:
	while sleep 0.5; do curl http://$(IP):$(PORT)/version; done

hostname:
	while sleep 0.5; do curl http://$(IP):$(PORT)/host; done