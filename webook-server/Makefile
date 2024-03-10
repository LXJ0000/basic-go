.PHONY: docker
docker:
	@rm webook || true
	@GOOS=linux GOARCH=arm go build -o webook .
	@docker image rm jannanxj/webook:v1 || true
	@docker build -t jannanxj/webook:v1 .