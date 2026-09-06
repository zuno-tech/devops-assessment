# Zuno DevOps Assessment

## About

This repository contains a small project intended to assess an engineer's fundamental skills in application build and deployment workflows.

The repository includes everything required to build a Docker container containing the provided web application.

Detailed instructions outlining the assessment requirements and expected deliverables are provided separately in the accompanying PDF.

## Build Requirements

The following software is required to build and run the application:

* Docker

## Building and Running

### Building

To build the Docker image containing the web application, run:

```bash
docker build . -t webapp
```

### Running

To run the container locally, use:

```bash
docker run -p 8080:8080 webapp
```

Once the container is running, the application should be available at:

`http://localhost:8080`

## Restrictions

The following files must not be modified as part of the assessment:

* `main.go`
* `go.mod`
