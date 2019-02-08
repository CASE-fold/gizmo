all: test

deps:
	go get -d -v github.com/CASE-fold/gizmo/v2/...

updatedeps:
	go get -d -v -u -f github.com/CASE-fold/gizmo/v2/...

testdeps:
	go get -d -v -t github.com/CASE-fold/gizmo/v2/...

updatetestdeps:
	go get -d -v -t -u -f github.com/CASE-fold/gizmo/v2/...

build: deps
	go build github.com/CASE-fold/gizmo/v2/...

install: deps
	go install github.com/CASE-fold/gizmo/v2/...

lint: testdeps
	go get -v golang.org/x/lint/golint
	for file in $$(find . -name '*.go' | grep -v '\.pb\.go\|\.pb\.gw\.go\|examples\|pubsub\/aws\/awssub_test\.go' | grep -v 'server\/kit\/kitserver_pb_test\.go'); do \
		golint $${file}; \
		if [ -n "$$(golint $${file})" ]; then \
			exit 1; \
		fi; \
	done

vet: testdeps
	go vet github.com/CASE-fold/gizmo/v2/...

errcheck: testdeps
	go get -v github.com/kisielk/errcheck
	errcheck -ignoretests github.com/CASE-fold/gizmo/v2/...

pretest: lint vet # errcheck

test: testdeps pretest
	go test github.com/CASE-fold/gizmo/v2/...

clean:
	go clean -i github.com/CASE-fold/gizmo/v2/...

coverage: testdeps
	./coverage.sh --coveralls

.PHONY: \
	all \
	deps \
	updatedeps \
	testdeps \
	updatetestdeps \
	build \
	install \
	lint \
	vet \
	errcheck \
	pretest \
	test \
	clean \
	coverage
