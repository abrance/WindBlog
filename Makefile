
.PHONY: release
release:
	@echo "start release server"
	@echo "pkg/api/ : 保存 gate 的对象"
	@mkdir -p build/WindBlog/bin
	@go build -o build/WindBlog/bin/server ./cmd/server
	@cp util/log/log.json ./config.ini build/WindBlog/bin/
	@tar zcvf WindBlog.tar.gz build/

.PHONY: clean
clean:
	@echo "start clean"
	@rm build/* -r

.PHONY: test-run
test-run:
	@echo "test run"
	@mkdir -p build/WindBlog/bin
	@go build -o build/WindBlog/bin/server ./cmd/server
	@cp util/log/log.json ./config.ini build/WindBlog/bin/
	cd build/WindBlog/bin && bash -c ./server
