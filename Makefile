.PHONY: build

build:
	sam build

deploy:
	sam deploy

local-exec:
	sam local invoke $(FUNCTION_NAME) \
	--region ap-northeast-1 \
	--env-vars .env.json