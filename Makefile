SWAGGER_DIR?=api/swagger

${SWAGGER_DIR}/bpdts-test-app/.codegen: ${SWAGGER_DIR}/bpdts-test-app/swagger.yaml
	swagger generate client --spec $< --target ${SWAGGER_DIR}/bpdts-test-app
	touch $@

${SWAGGER_DIR}/interview-server/.codegen: ${SWAGGER_DIR}/interview-server/swagger.yaml
	swagger generate server --spec $< --target ${SWAGGER_DIR}/interview-server
	touch $@

.PHONY: build
build: ${SWAGGER_DIR}/bpdts-test-app/.codegen ${SWAGGER_DIR}/interview-server/.codegen
	go build ./
