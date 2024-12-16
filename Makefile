BIN_DIR=$(PWD)/bin
MLAN_DIR=$(PWD)/cmd/mlan
CC=gcc
CXX=g++
ANTLR_VERSION=4.13.2

.PHONY: mlan
mlan: gen-parser
	@mkdir -p ${BIN_DIR}
	@echo "Building mlan..."
	@CC=${CC} CXX=${CXX} go build -o ${BIN_DIR}/mlan ${MLAN_DIR}

.PHONY: test-mlan
test-mlan: mlan
	@echo "Testing language..."
	@${BIN_DIR}/mlan --file tests/assert.mlan
	@echo "  assert: OK"
	@${BIN_DIR}/mlan --file tests/recursion_max_depth.mlan 1>/dev/null 2>/dev/null 2>/dev/null || true
	@echo "  recursion_max_depth: OK"
	@${BIN_DIR}/mlan --file tests/cast.mlan 1>/dev/null 2>/dev/null 2>/dev/null
	@echo "  cast: OK"
	@${BIN_DIR}/mlan --file tests/files.mlan
	@echo "  files: OK"

.PHONY: gen-parser
gen-parser: download-antlr
	@echo "Generating parser..."
	@java -jar files/antlr-${ANTLR_VERSION}.jar -o pkg/parser -visitor -package parser -Dlanguage=Go ./Mlan.g4

download-antlr:
	@echo "Downloading antlr ${ANTLR_VERSION}..."
	@curl -s -o files/antlr-${ANTLR_VERSION}.jar https://www.antlr.org/download/antlr-${ANTLR_VERSION}-complete.jar