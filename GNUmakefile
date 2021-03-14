default: testacc

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m


build: 
	go build -o terraform-provider-fauna
	mkdir -p ~/.terraform.d/plugins/hashicorp.com/chronark/fauna/0.2/linux_amd64
	mv terraform-provider-fauna ~/.terraform.d/plugins/hashicorp.com/chronark/fauna/0.2/linux_amd64