# Website Health Checker

This is a Go application to check the health of a website by verifying its availability. You can provide a domain to check the website's status.

## Features
- Check the health of a website by providing its domain.
- Supports command-line arguments for domain and port.
- Includes Dockerfile and docker-compose for easy testing.

## Usage

### Command-Line Usage

To run the application, use the following command:

```bash
go run main.go --domain <website_domain> [--port <port>]
```

### Example
```bash
docker build -t website-health-checker .
docker run website-health-checker --domain <website_domain> --port <port>
docker run website-health-checker --domain google.com --port 80


