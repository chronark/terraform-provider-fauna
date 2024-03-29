default: testacc

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m


build: fmt
	go build -o terraform-provider-fauna
	mkdir -p ~/.terraform.d/plugins/hashicorp.com/chronark/fauna/9000.1/linux_amd64
	mv terraform-provider-fauna ~/.terraform.d/plugins/hashicorp.com/chronark/fauna/9000.1/linux_amd64


fmt:
	go generate -v ./...
	golangci-lint run -v
	go fmt ./...


rm-state:
	rm -rf examples/e2e/terraform.tfstate*


init: build
	rm -rf examples/e2e/.terraform*
	terraform -chdir=examples/e2e init -upgrade
apply: 
	terraform -chdir=examples/e2e apply


release:
	git tag $$(svu next) && git push --tags