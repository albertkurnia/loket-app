.PHONY: cover coverage mock test 

mock:
	mockgen -source=modules/event/query/query.go -destination=modules/event/query/mock/event_query_mockgen.go -package=mock
	mockgen -source=modules/event/usecase/usecase.go -destination=modules/event/usecase/mock/event_usecase_mockgen.go -package=mock
	mockgen -source=modules/location/query/query.go -destination=modules/location/query/mock/location_query_mockgen.go -package=mock
	mockgen -source=modules/location/usecase/usecase.go -destination=modules/location/usecase/mock/location_usecase_mockgen.go -package=mock
	mockgen -source=modules/transaction/query/query.go -destination=modules/transaction/query/mock/transaction_query_mockgen.go -package=mock
	mockgen -source=modules/transaction/usecase/usecase.go -destination=modules/transaction/usecase/mock/transaction_usecase_mockgen.go -package=mock

cover: coverage.txt coverage

coverage.txt: coverages/modules.txt
	gocovmerge $^ > $@

coverages/modules.txt:
	go test -race -short -coverprofile=$@ -covermode=atomic ./modules/... | grep -v mock

coverage:
	./codecov

test: SHELL := /bin/bash
test:
	go test ./... -v -race -short | grep -v mock ; exit $${PIPESTATUS[0]}