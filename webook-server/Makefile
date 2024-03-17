.PHONY: make
docker:
	@rm webook || true
	@GOOS=linux GOARCH=arm go build -o webook .
	@docker image rm jannanxj/webook:v1 || true
	@docker build -t jannanxj/webook:v1 .

mock:
	@mockgen -source=./internal/service/user.go -package=svcmock -destination=./internal/service/mocks/user_gen.go
	@mockgen -source=./internal/service/code.go -package=svcmock -destination=./internal/service/mocks/code_gen.go

	@mockgen -source=./internal/repository/user.go -package=repomock -destination=./internal/repository/mocks/user_gen.go
	@mockgen -source=./internal/repository/code.go -package=svcmock -destination=./internal/repository/mocks/code_gen.go

	@mockgen -source=./internal/repository/dao/user.go -package=svcmock -destination=./internal/repository/dao/mocks/user_gen.go

	@mockgen -source=./internal/repository/cache/user.go -package=svcmock -destination=./internal/repository/cache/mocks/user_gen.go
	@mockgen -source=./internal/repository/cache/code.go -package=svcmock -destination=./internal/repository/cache/mocks/code_gen.go

	@go mod tidy