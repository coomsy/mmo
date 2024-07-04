# Copyright 2024
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

GO           ?= go
GOFMT        ?= $(GO)fmt
GOHOSTOS     ?= $(shell $(GO) env GOHOSTOS)
GOHOSTARCH   ?= $(shell $(GO) env GOHOSTARCH)

GO_VERSION        ?= $(shell $(GO) version)
GO_VERSION_NUMBER ?= $(word 3, $(GO_VERSION))

GOTEST := $(GO) test
GOTEST_DIR :=

PKGS = cmd/api # when have all tests "./..."

.PHONY: help
help: 
	@echo "run 'make api-serve' in one terminal, \
	then 'make ui-serve' in another terminal. \
	UI will be at http://localhost:3000 "

include web/Makefile

.PHONY: go-serve
api-serve:
	MMO_DEBUG=true go run cmd/api/api.go

.PHONY: test
test: go-test-full ui-test-full

##
# common go cmds, eventually make into makefile
##

.PHONY: go-test-full
go-test-full: go-test

$(GOTEST_DIR):
	@mkdir -p $@

.PHONY: go-test
go-test: $(GOTEST_DIR)
	@echo ">> go tests"
	$(GOTEST) $(PKGS)

