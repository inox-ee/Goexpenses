.PHONY: deploy-image update-lambda-% test ecr-login
include .env

DATE := $(shell date +%Y%m%d%H%M%S)
deploy-image: test
	docker build -t inoxee/goexpenses:latest . --no-cache
	docker tag inoxee/goexpenses:latest $(AWS_ECR_REPOSITORY_URL):$(DATE)
	docker push $(AWS_ECR_REPOSITORY_URL):$(DATE)

update-lambda-%:
	aws lambda update-function-code --function-name goexpenses-slackbot --image-uri $(AWS_ECR_REPOSITORY_URL):${@:update-lambda-%=%}

test:
	go test -v

ecr-login:
	aws ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin $(AWS_ECR_URL)
