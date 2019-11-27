GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
GOTEST=$(GOCMD) test
BINARY_NAME=local-weather

all: test build
build:
	$(GOBUILD) -o $(BINARY_NAME) cmd/local-weather/main.go
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run:
	$(GOBUILD) -o $(BINARY_NAME) cmd/local-weather/main.go 
	./$(BINARY_NAME) --zipcode 10007 
deps:
	$(GOGET) github.com/olekukonko/tablewriter
	$(GOGET) github.com/justinrsmith/local-weather-cli-go/pkg/fetchweather
