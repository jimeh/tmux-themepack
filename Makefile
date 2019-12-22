BUILDER := bin/build-theme
THEME_SRC := $(shell find src -name '*.tmuxtheme')
INCLUDES := $(shell find src -name '*.tmuxsh')
THEMES := $(patsubst src/%,%,$(THEME_SRC))
TESTS := $(addsuffix .test,$(THEMES))

.PHONY: build
build: $(THEMES)

.PHONY: clean
clean:
	rm $(shell find * -name "*.tmuxtheme" -not -path "src/*")

.PHONY: lint
lint:
	cd test && golangci-lint run -v

.PHONY: test
test: needs-build
	cd test && go test -count=1 -v ./...

.PHONY: needs-build
needs-build:
	$(foreach file,$(THEMES), \
		$(BUILDER) "src/$(file)" | diff -q "$(file)" - && \
	) true

$(THEMES): %.tmuxtheme: src/%.tmuxtheme $(INCLUDES)
	$(BUILDER) "src/$@" "$@"

$(TESTS): %.test: src/%.test
