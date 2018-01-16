PKG_NAME=dwapi
COVERAGE_FILE=/tmp/dwapi.cov

fmt:
	gofmt -s -w ./$(PKG_NAME)

fmtcheck:
	@gofmt -l -s ./$(PKG_NAME)

lint:
	@gometalinter ./$(PKG_NAME)

test: fmtcheck
	@go test ./$(PKG_NAME) -timeout=30s -parallel=4 -cover

coverage-statistics-breakdown:
	@go test ./$(PKG_NAME) -timeout=30s -parallel=4 -coverprofile=${COVERAGE_FILE}; \
	go tool cover -func=${COVERAGE_FILE}
