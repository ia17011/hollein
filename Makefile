
.PHONY: build
build:
	sam build

.PHONY: deploy
deploy:
	sam deploy

.PHONY: invoke
invoke:
	sam local invoke

local-exec:
	sam local invoke $(FUNCTION_NAME) \
	--region ap-northeast-1 \
	--env-vars .env.json