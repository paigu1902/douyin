# start the environment of douyin
.PHONY: start
start:
	docker-compose up -d

# stop the environment of douyin
.PHONY: stop
stop:
	docker-compose down

.PHONY: api-gateway
api-gateway:
	go run ../service/api-gateway/

.PHONY: rpc-user-info
rpc-user-info:
	go run ../service/rpc-user-info/

.PHONY: rpc-user-relation
rpc-user-relation:
	go run ../service/rpc-user-relation/