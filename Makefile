# Set these to the desired values
ARTIFACT_ID=hsa
VERSION=0.1.0

GOTAG=1.21.1
MAKEFILES_VERSION=8.3.1
.DEFAULT_GOAL:=default
LDFLAGS=-ldflags "-X main.Version=${VERSION}"

include build/make/variables.mk
include build/make/dependencies-gomod.mk
include build/make/build.mk
include build/make/mocks.mk
include build/make/test-common.mk
include build/make/test-unit.mk
include build/make/static-analysis.mk
include build/make/clean.mk
include build/make/digital-signature.mk
include build/make/self-update.mk

default: help

.PHONY: run
run:
	@go run .