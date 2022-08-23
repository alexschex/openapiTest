java -jar ./openapi-generator-cli-5.3.1-SpectraFork.jar generate \
		-i ./test.yaml \
		-g go \
		-o pkg/lumosclient/lumosrestapiclient \
		-p enumClassPrefix=true \
		-p useOneOfDiscriminatorLookup=true \
		--additional-properties=generateInterfaces=true \
		--additional-properties=useOneOfDiscriminatorLookup=true
