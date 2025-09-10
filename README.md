# Microservice Deployment with GitHub Actions

This repository demonstrates the automation of a Golang microservice deployment using GitHub Actions.

## About the Project

This is a simple Go microservice built with the Gin framework that exposes two endpoints:
- `/ping` - Returns a "pong" message
- `/health` - Returns the health status of the service

## CI/CD Pipeline

The GitHub Actions workflow automates the following tasks:

1. Builds and tests the Go application
2. Builds a Docker image from the Dockerfile
3. Pushes the Docker image to Docker Hub

## Workflow Triggers

The workflow can be triggered in two ways:
1. Automatically on push to the `main` branch
2. Manually through the GitHub Actions interface (workflow_dispatch)

## Setup Instructions

### Prerequisites
- GitHub account
- Docker Hub account

### GitHub Secrets

To use this workflow, you need to add the following secret to your GitHub repository:

- `DOCKER_HUB_TOKEN` - Your Docker Hub access token

### Manual Workflow Run

1. Go to the "Actions" tab in your GitHub repository
2. Select the "Go CI" workflow
3. Click "Run workflow"
4. Enter your Docker image name or use the default

### Testing the Microservice

Once deployed, you can test the microservice using:

```bash
# Test ping endpoint
curl http://<your-service-url>/ping

# Test health endpoint
curl http://<your-service-url>/health
```

## Dockerfile

The service is containerized using a multi-stage Dockerfile for optimal image size and security.
