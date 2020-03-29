SWAGGER_DIR?=api/swagger

${SWAGGER_DIR}/bpdts-test-app/.codegen: ${SWAGGER_DIR}/bpdts-test-app/swagger.yaml
	swagger generate client --spec $< --target bpdts
	touch $@

${SWAGGER_DIR}/interview-server/.codegen: ${SWAGGER_DIR}/interview-server/swagger.yaml
	swagger generate server --spec $<
	touch $@

.PHONY: build
build: ${SWAGGER_DIR}/bpdts-test-app/.codegen ${SWAGGER_DIR}/interview-server/.codegen
	go build ./
