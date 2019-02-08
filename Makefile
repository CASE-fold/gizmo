all: test

deps:
	go get -d -v github.com/case-fold/gizmo/...

updatedeps:
	go get -d -v -u -f github.com/case-fold/gizmo/...

testdeps:
	go get -d -v -t github.com/case-fold/gizmo/...

updatetestdeps:
	go get -d -v -t -u -f github.com/case-fold/gizmo/...

build: deps
	go build github.com/case-fold/gizmo/...

install: deps
	go install github.com/case-fold/gizmo/...

lint: testdeps
	go get -v golang.org/x/lint/golint
	for file in $$(find . -name '*.go' | grep -v '\.pb\.go\|\.pb\.gw\.go\|examples\|pubsub\/aws\/awssub_test\.go' | grep -v 'server\/kit\/kitserver_pb_test\.go'); do \
		golint $${file}; \
		if [ -n "$$(golint $${file})" ]; then \
			exit 1; \
		fi; \
	done

vet: testdeps
	go vet github.com/case-fold/gizmo/...

errcheck: testdeps
	go get -v github.com/kisielk/errcheck
	errcheck -ignoretests github.com/case-fold/gizmo/...

pretest: lint vet # errcheck

test: testdeps pretest
	go test github.com/case-fold/gizmo/...

clean:
	go clean -i github.com/case-fold/gizmo/...

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
