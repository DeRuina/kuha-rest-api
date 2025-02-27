include .envrc

.PHONY: gen-docs
gen-docs:
	@swag init -g ./api/main.go -d cmd,internal,docs/swagger && swag fmt