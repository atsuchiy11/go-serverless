# Makeファイルの環境変数をcompose.yamlの環境変数に渡す
# Ex) make compose-up HANDLER=./handler/hello_world/main.go
compose-up:
	env HANDLER=$(HANDLER) docker compose -f ./compose.yaml up -d

compose-down:
	docker compose down

invoke:
	curl -i -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"Name":"Taro"}'