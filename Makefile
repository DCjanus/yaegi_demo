SRC_FILES := $(shell find . -name '*.go' -not -path "./vendor/*" -not -path "./output/*" -not -path "./rule/rule.go")

internal/symbol/.genkeep: $(SRC_FILES)
	go generate ./...
	@echo "generated at `date`" > internal/symbol/.genkeep

generate: internal/symbol/.genkeep

output/yaegi_demo: generate $(SRC_FILES)
	go build -mod mod -o output/yaegi_demo cmd/main.go

build: output/yaegi_demo

run: build
	./output/yaegi_demo

clean:
	rm -rf output
	find internal/symbol -type f -not -name 'symbol.go' -delete

.PHONY: generate build run clean