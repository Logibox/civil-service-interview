BIN_DIR=.bin
SWAGGER_DIR?=swagger

.PHONY: all
all: ${BIN_DIR}/interview-api-server

.PHONE: test
test:
	go test -v ./...

.PHONY: clean
clean:
	rm -rf ${BIN_DIR}

${SWAGGER_DIR}/bpdts-test-app/.codegen: ${SWAGGER_DIR}/bpdts-test-app/swagger.yaml
	swagger generate client --spec $< --target bpdts
	touch $@

${SWAGGER_DIR}/interview-server/.codegen: ${SWAGGER_DIR}/interview-server/swagger.yaml
	swagger generate server --spec $<
	touch $@

${BIN_DIR}:
	mkdir -p ${BIN_DIR}

${BIN_DIR}/interview-api-server: ${BIN_DIR} ${SWAGGER_DIR}/bpdts-test-app/.codegen ${SWAGGER_DIR}/interview-server/.codegen
	go build -o $@ ./cmd/interview-api-server
