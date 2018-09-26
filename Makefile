init:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure
	cd pkg && go build -o ../bin/cb-sample-server
build:
	cd pkg && go build -o ../bin/cb-sample-server
docker-build:
	docker build --tag=asia.gcr.io/${PROJECT_ID}/cb-sample-server:local .
	docker run -d asia.gcr.io/${PROJECT_ID}/cb-sample-server:local