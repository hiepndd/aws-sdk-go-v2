LINTIGNOREDOT='internal/awstesting/integration.+should not use dot imports'
LINTIGNOREDOC='service/[^/]+/(api|service|waiters)\.go:.+(comment on exported|should have comment or be unexported)'
LINTIGNORECONST='service/[^/]+/(api|service|waiters)\.go:.+(type|struct field|const|func) ([^ ]+) should be ([^ ]+)'
LINTIGNORESTUTTER='service/[^/]+/(api|service)\.go:.+(and that stutters)'
LINTIGNOREINFLECT='service/[^/]+/(api|errors|service)\.go:.+(method|const) .+ should be '
LINTIGNOREINFLECTS3UPLOAD='service/s3/s3manager/upload\.go:.+struct field SSEKMSKeyId should be '
LINTIGNOREDEPS='vendor/.+\.go'
LINTIGNOREPKGCOMMENT='service/[^/]+/doc_custom.go:.+package comment should be of the form'
LINTIGNOREENDPOINTS='aws/endpoints/defaults.go:.+(method|const) .+ should be '
LINTIGNORESINGLEFIGHT='internal/sync/singleflight/singleflight.go:.+error should be the last type'
UNIT_TEST_TAGS="example codegen awsinclude"
ALL_TAGS="example codegen awsinclude integration perftest sdktool"

# SDK's Core and client packages that are compatable with Go 1.9+.
SDK_CORE_PKGS=./aws/... ./private/... ./internal/...
SDK_CLIENT_PKGS=./service/...
SDK_COMPA_PKGS=${SDK_CORE_PKGS} ${SDK_CLIENT_PKGS}

# SDK additional packages that are used for development of the SDK.
SDK_EXAMPLES_PKGS=./example/...
SDK_MODELS_PKGS=./models/...
SDK_ALL_PKGS=${SDK_COMPA_PKGS} ${SDK_EXAMPLES_PKGS} ${SDK_MODELS_PKGS}


all: generate unit

###################
# Code Generation #
###################
generate: cleanup-models gen-test gen-endpoints gen-services gen-external-asserts

gen-test: gen-protocol-test gen-codegen-test

#gen-codegen-test:
#	@echo "Generating SDK API tests"
#	go generate ./private/model/api/codegentest/service

gen-services:
	@echo "Generating SDK clients"
	go generate ./service

gen-protocol-test:
	@echo "Generating SDK protocol tests"
	go generate ./private/protocol/...

gen-endpoints:
	@echo "Generating SDK endpoints"
	go generate ./models/endpoints

gen-codegen-test:
	@echo "Generating SDK API tests"
	go generate ./private/model/api/codegentest/service

gen-external-asserts:
	@echo "Generating SDK external package implementor assertions"
	go generate ./aws/external

cleanup-models:
	@echo "Cleaning up stale model versions"
	@./cleanup_models.sh

###################
# Unit/CI Testing #
###################
build:
	go build -o /dev/null -tags ${ALL_TAGS} ${SDK_ALL_PKGS}

unit: verify build
	@echo "go test SDK and vendor packages"
	@go test -tags ${UNIT_TEST_TAGS} ${SDK_ALL_PKGS}

unit-with-race-cover: verify build
	@echo "go test SDK and vendor packages"
	@go test -tags ${UNIT_TEST_TAGS} -race -cpu=1,2,4 ${SDK_ALL_PKGS}

ci-test: generate unit-with-race-cover ci-test-generate-validate

ci-test-generate-validate:
	@echo "CI test validate no generated code changes"
	git update-index --assume-unchanged go.mod go.sum
	git add . -A
	gitstatus=`git diff --cached --ignore-space-change`; \
	git update-index --no-assume-unchanged go.mod go.sum
	echo "$$gitstatus"; \
	if [ "$$gitstatus" != "" ]; then echo "$$gitstatus"; exit 1; fi

#######################
# Integration Testing #
#######################
integration: core-integ client-integ

core-integ:
	@echo "Integration Testing SDK core"
	AWS_REGION="" go test -count=1 -tags "integration" -v -run '^TestInteg_' ${SDK_CORE_PKGS}

client-integ:
	@echo "Integration Testing SDK clients"
	AWS_REGION="" go test -count=1 -tags "integration" -v -run '^TestInteg_' ./service/...

s3crypto-integ:
	@echo "Integration Testing S3 Cyrpto utility"
	AWS_REGION="" go test -count=1 -tags "s3crypto_integ integration" -v -run '^TestInteg_' ./service/s3/s3crypto

cleanup-integ-buckets:
	@echo "Cleaning up SDK integraiton resources"
	go run -tags "integration" ./internal/awstesting/cmd/bucket_cleanup/main.go "aws-sdk-go-integration"

###################
# Sandbox Testing #
###################
sandbox-tests: sandbox-test-go1.12 sandbox-test-go1.13 sandbox-test-gotip

sandbox-build-go1.13:
	docker build -f ./internal/awstesting/sandbox/Dockerfile.test.go1.13 -t "aws-sdk-go-v2-1.13" .
sandbox-go1.13: sandbox-build-go1.13
	docker run -i -t aws-sdk-go-v2-1.13 bash
sandbox-test-go1.13: sandbox-build-go1.13
	docker run -t aws-sdk-go-v2-1.13

sandbox-build-go1.12:
	docker build -f ./internal/awstesting/sandbox/Dockerfile.test.go1.12 -t "aws-sdk-go-v2-1.12" .
sandbox-go1.12: sandbox-build-go1.12
	docker run -i -t aws-sdk-go-v2-1.12 bash
sandbox-test-go1.12: sandbox-build-go1.12
	docker run -t aws-sdk-go-v2-1.12

sandbox-build-gotip:
	@echo "Run make update-aws-golang-tip, if this test fails because missing aws-golang:tip container"
	docker build -f ./internal/awstesting/sandbox/Dockerfile.test.gotip -t "aws-sdk-go-v2-tip" .
sandbox-gotip: sandbox-build-gotip
	docker run -i -t aws-sdk-go-v2-tip bash
sandbox-test-gotip: sandbox-build-gotip
	docker run -t aws-sdk-go-v2-tip

update-aws-golang-tip:
	docker build --no-cache=true -f ./internal/awstesting/sandbox/Dockerfile.golang-tip -t "aws-golang:tip" .

##################
# Linting/Verify #
##################
verify: lint vet sdkv1check

lint:
	@echo "go lint SDK and vendor packages"
	@lint=`golint ./...`; \
	dolint=`echo "$$lint" | grep -E -v \
	-e ${LINTIGNOREDOC} \
	-e ${LINTIGNORECONST} \
	-e ${LINTIGNORESTUTTER} \
	-e ${LINTIGNOREINFLECT} \
	-e ${LINTIGNOREDEPS} \
	-e ${LINTIGNOREINFLECTS3UPLOAD} \
	-e ${LINTIGNOREPKGCOMMENT} \
	-e ${LINTIGNOREENDPOINTS} \
	-e ${LINTIGNORESINGLEFIGHT}`; \
	echo "$$dolint"; \
	if [ "$$dolint" != "" ]; then exit 1; fi

vet:
	go vet -tags "example codegen awsinclude integration" --all ${SDK_ALL_PKGS}

sdkv1check:
	@echo "Checking for usage of AWS SDK for Go v1"
	@sdkv1usage=`go list -test -f '''{{ if not .Standard }}{{ range $$_, $$name := .Imports }} * {{ $$.ImportPath }} -> {{ $$name }}{{ print "\n" }}{{ end }}{{ range $$_, $$name := .TestImports }} *: {{ $$.ImportPath }} -> {{ $$name }}{{ print "\n" }}{{ end }}{{ end}}''' ./... | sort -u | grep '''/aws-sdk-go/'''`; \
	echo "$$sdkv1usage"; \
	if [ "$$sdkv1usage" != "" ]; then exit 1; fi

################
# Dependencies #
################
get-deps: get-deps-tests get-deps-x-tests get-deps-codegen get-deps-verify
	go get github.com/jmespath/go-jmespath

get-deps-tests:
	@echo "go get SDK testing dependencies"
	go get golang.org/x/net/html
	go get github.com/google/go-cmp

get-deps-x-tests:
	@echo "go get SDK testing golang.org/x dependencies"
	go get golang.org/x/net/http2

get-deps-codegen: get-deps-x-tests
	@echo "go get SDK codegen dependencies"
	go get golang.org/x/net/html

get-deps-verify:
	@echo "go get SDK verification utilities"
	go get golang.org/x/lint/golint

##############
# Benchmarks #
##############
bench:
	@echo "go bench SDK packages"
	@go test -run NONE -bench . -benchmem -tags 'bench' ${SDK_ALL_PKGS}

bench-protocol:
	@echo "go bench SDK protocol marshallers"
	@go test -run NONE -bench . -benchmem -tags 'bench' ./private/protocol/...

#############
# Utilities #
#############
api_info:
	@go run private/model/cli/api-info/api-info.go
