APP_EXECUTABLE=main

build:
	GOARCH=amd64 GOOS=linux go build -o ${APP_EXECUTABLE} main.go

run: build
	./${APP_EXECUTABLE}

init: build
	./${APP_EXECUTABLE} init

add: build
	./${APP_EXECUTABLE} add .

commit: build
	./${APP_EXECUTABLE} commit -m "first commit"

log: build
	./${APP_EXECUTABLE} log

status: build
	./${APP_EXECUTABLE} status

config-name: build
	./${APP_EXECUTABLE} config user.name "TonyGLL"

config-email: build
	./${APP_EXECUTABLE} config user.email "tonygllambia@gmail.com"

config-list-user: build
	./${APP_EXECUTABLE} config list user

branch-list: build
	./${APP_EXECUTABLE} branch

branch-create: build
	./${APP_EXECUTABLE} branch new-feature

branch-delete: build
	./${APP_EXECUTABLE} branch -d new-feature

lint: ## Runs the linter (golangci-lint) to analyze the code.
	@echo "==> Linting code with golangci-lint..."
	@golangci-lint run