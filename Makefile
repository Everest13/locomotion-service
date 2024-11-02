run:
	$ go run ./cmd/locomotion-service

GOOSE_BIN:=$(LOCAL_BIN)/goose
.PHONY: install-goose
install-goose:
ifeq ($(wildcard $(GOOSE_BIN)),)
	$(info #Downloading goose)
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.5.0
	go mod tidy
GOOSE_BIN:=$(LOCAL_BIN)/goose
endif

#накатить миграции
goose-up:
	cd migrations
	goose postgres "dbname=locomotion sslmode=disable" up
	cd ../

#
protoc:
	protoc -I ./api \
        --go_out ./pkg --go_opt paths=source_relative \
        --go-grpc_out ./pkg --go-grpc_opt paths=source_relative \
        --grpc-gateway_out ./pkg --grpc-gateway_opt paths=source_relative \
        ./api/user-service/user-service.proto


        protoc -I ./api \
            --go_out ./pkg --go_opt paths=source_relative \
		    --go-grpc_out ./pkg --go-grpc_opt paths=source_relative \
            --grpc-gateway_out ./pkg --grpc-gateway_opt paths=source_relative \
            ./api/locomotion-service/locomotion-service.proto


    protoc -I ./api \
    --go_out ./pkg/user-service --go_opt paths=source_relative \
    ./api/user-service/user-service.proto