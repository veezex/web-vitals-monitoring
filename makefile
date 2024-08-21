include .env

# Variables
GO_BUILD_CMD=go build -o serverapp cmd/main.go
CERTBOT_CMD=certbot certonly --standalone -d $(DOMAIN) --non-interactive --agree-tos -m admin@$(DOMAIN) --http-01-port=80
MKDIR_CERTS=mkdir -p /app/certs
CP_AND_CHMOD_CERTS=cp /etc/letsencrypt/live/$(DOMAIN)/* /app/certs && chmod -R 740 /app/certs

# Default target
all: build certs

# Build the Go application
build:
	@echo "Building the Go application..."
	$(GO_BUILD_CMD)

# Generate and setup certificates
certs:
	@echo "Generating certificates with Certbot..."
	$(CERTBOT_CMD)
	@echo "Setting up certificate directory and adjusting permissions..."
	$(MKDIR_CERTS)
	$(CP_AND_CHMOD_CERTS)

vet:
	go vet ./...

.PHONY: all vet build certs setup_user
