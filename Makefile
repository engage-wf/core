
# Makefile for core

VERSION ?= 0.0.0

test:

test-prerequisites:

install-tools:

### TEST ####################################################################

test-core:
	ginkgo
test-core-watch:
	ginkgo watch
test: test-core
.PHONY: test-core
.PHONY: test
