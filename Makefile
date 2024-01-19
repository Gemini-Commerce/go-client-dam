version=1.0.0

generate:
	$(MAKE) -C ./modules/go-client-generator/ generate-go-client service=dam version=${version}
