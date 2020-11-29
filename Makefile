tidy:
	$(foreach x,$(shell find . -name "go.mod"),$(call withNewLine ,pushd $(dir $(x)) && go mod tidy; popd))
.PHONY: tidy

define withNewLine
$(1)

endef
