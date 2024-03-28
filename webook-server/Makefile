.PHONY: make
docker:
	@rm webook || true
	@GOOS=linux GOARCH=arm go build -o webook .
	@docker image rm jannanxj/webook:v1 || true
	@docker build -t jannanxj/webook:v1 .

mock:
# service
	@mockgen -source=./internal/service/user.go -package=svcmock -destination=./internal/service/mocks/user_gen.go
	@mockgen -source=./internal/service/code.go -package=svcmock -destination=./internal/service/mocks/code_gen.go
	@mockgen -source=./internal/service/article.go -package=svcmock -destination=./internal/service/mocks/article_gen.go
# repository
	@mockgen -source=./internal/repository/user.go -package=repomock -destination=./internal/repository/mocks/user_gen.go
	@mockgen -source=./internal/repository/code.go -package=repomock -destination=./internal/repository/mocks/code_gen.go
	@mockgen -source=./internal/repository/article.go -package=repomock -destination=./internal/repository/mocks/article_gen.go
# dao
	@mockgen -source=./internal/repository/dao/user.go -package=daomock -destination=./internal/repository/dao/mocks/user_gen.go
	@mockgen -source=./internal/repository/dao/article.go -package=daomock -destination=./internal/repository/dao/mocks/article_gen.go
# cache
	@mockgen -source=./internal/repository/cache/user.go -package=cachemock -destination=./internal/repository/cache/mocks/user_gen.go
	@mockgen -source=./internal/repository/cache/code.go -package=cachemock -destination=./internal/repository/cache/mocks/code_gen.go
	@mockgen -source=./internal/repository/cache/article.go -package=cachemock -destination=./internal/repository/cache/mocks/article_gen.go
# handler
	@mockgen -source=./internal/web/middleware/auth.go -package=middlewaremock -destination=./internal/web/middleware/mocks/auth_gen.go
	@go mod tidy