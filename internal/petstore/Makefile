SHELL = bash
GENVERS = 6.0.1
GO_POST_PROCESS_FILE = "/usr/bin/env gofmt -w"
SPECFILE = petstore.yaml

.PHONY: clean checkbash checkcurl checkgo checkjava

all: openapi-generator-cli.jar checkdeps petstore.yaml



openapi-generator-cli.jar:
	curl -sSL "https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/$(GENVERS)/openapi-generator-cli-$(GENVERS).jar" -o openapi-generator-cli.jar

checkdeps: checkbash checkjava checkgo checkcurl
	@echo 🎉 All build dependencies met 👌

checkbash:
	@(( $${BASH_VERSION:0:1} >= 4 )) && \
		printf "%-20s ✅\n" "Bash Version: $${BASH_VERSION:0:1}"

checkcurl:
	@vers=$$(curl -V 2>&1| head -1 | awk '{print $$2}'); \
		printf "%-20s ✅\n" "Curl Version: $$vers"

checkgo:
	@vers=$$(go version| awk '{print $$3}'); \
		printf "%-20s ✅\n" "Go Version: $$vers"

checkjava:
	@jvers=$$(java -version 2>&1 | head -1); \
		jvers=$${jvers#*version \"}; \
		jvers=$${jvers%%\"*}; \
		[[ $$jvers =~ ^(1\.[89]|9\.|[1-9][0-9]+) ]] && \
		printf "%-20s ✅\n" "Java Version: $$jvers"

validate: petstore.yaml
	java -jar openapi-generator-cli.jar validate -i $(SPECFILE) --recommend

go:
	java -jar openapi-generator-cli.jar generate -g go-gin-server -i "petstore.yaml"

clean:
	rm openapi-generator-cli.jar
