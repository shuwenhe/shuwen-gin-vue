# 配置
service_name = shuwen-checklist
config_path = config/dev.yaml
tag = dev

.PHONY:help
help: ##@other Show this help.
	@perl -e '$(HELP_FUN)' $(MAKEFILE_LIST)

.PHONY:run
run: ##@run 启动服务.
	@echo "启动服务..."
	go run -race . --config=$(config_path)

.PHONY:build-linux
build-linux: ##@build 构建二进制文件.
	@echo "构建二进制文件"
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $(service_name) .
	upx $(service_name)

HELP_FUN = \
	%help; \
	while(<>) { push @{$$help{$$2 // 'options'}}, [$$1, $$3] if /^([a-zA-Z\-]+)\s*:.*\#\#(?:@([a-zA-Z\-]+))?\s(.*)$$/ }; \
	print "usage: make [target]\n\n"; \
	for (sort keys %help) { \
		print "${WHITE}$$_:${RESET}\n"; \
		for (@{$$help{$$_}}) { \
			$$sep = " " x (32 - length $$_->[0]); \
			print "  ${YELLOW}$$_->[0]${RESET}$$sep${GREEN}$$_->[1]${RESET}\n"; \
	}; \
	print "\n"; }
