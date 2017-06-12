.PHONY: trinity protos

GOBIN = build/bin
INSTALL_DIR = installdir

TRINITY_HOME = $(HOME)/.trinity

BLD_MSG  =`echo "[BUILDING]:   "`
INST_MSG=`echo "[INSTALLING]: "`
UPD_MSG =`echo "[UPDATING]:   "`

INSTALLED_TARGETS = $(INSTALL_DIR)/trinity
INSTALLED_TARGETS += $(INSTALL_DIR)/trinity_gateway


GRPC_CPP_PLUGIN:=$(shell which grpc_cpp_plugin)

all: deps build

build: version trinity_gateway trinity

install: deps build
	@build/env.sh build/static.sh
	@echo "\ninstalling to $(INSTALL_DIR)"
	cp build/bin/* $(INSTALL_DIR)

trinity:
	@echo "$(BLD_MSG) trinity"
	@build/env.sh go build -i -o build/bin/trinity ./trinity

trinity_gateway:
	@echo "$(BLD_MSG) trinity_gateway"
	@build/env.sh go build -i -o build/bin/trinity_gateway ./subsystems/trinity_gateway

protos:
	protoc -I protos protos/trinity.proto --gofast_out=plugins=grpc:api/trinity
	protoc -I protos protos/trinity.proto --grpc_out=protos/out --plugin=protoc-gen-grpc=$(GRPC_CPP_PLUGIN)
	protoc -I protos protos/trinity.proto --cpp_out=protos/out

deps: goget_deps gx_deps

goget_deps:
	@echo "$(INST_MSG) go get dependencies"
	@build/env.sh go get google.golang.org/grpc

gx_deps:
	@echo "$(INST_MSG) gx dependencies"
	@build/env.sh gx install --global > /dev/null

version:
	@echo "$(UPD_MSG) version.go"
	@build/version.sh

clean:
	@echo "\ncleaning up.....\n"
	@build/env.sh rm -fr build/_workspace/ $(GOBIN)/*

installed_dirs:
	@echo $(INSTALLED_TARGETS)


uninstall:
	rm -rf $(INSTALLED_TARGETS) $(TRINITY_HOME)
