# ASCII Art Web Dockerize

A containerized version of the ASCII Art Web application.

## Description

This project extends the ASCII Art Web application by providing Docker containerization. It packages the web server and all its dependencies into a Docker container for easy deployment and consistent environment.

## Features

- All functionality of the ASCII Art Web application
- Docker containerization
- Easy deployment across different platforms
- Environment isolation

## Usage

### Building the Docker Image

```bash
docker build -t ascii-art-web .
```

### Running the Container

```bash
docker run -p 8080:8080 ascii-art-web
```

Then navigate to `http://localhost:8080` in your web browser.

## Docker Configuration

- Base image: `golang:alpine` for a minimal footprint
- Multi-stage build process for optimized image size
- Exposed port: 8080
- Volume support for banner files

## Implementation Details

- Dockerfile with best practices
- Container health checks
- Environment variable configuration
- Proper signal handling for graceful shutdown

## Requirements

- Docker
- Web browser with JavaScript enabled
