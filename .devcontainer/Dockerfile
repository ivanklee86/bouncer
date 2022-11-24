# See here for image contents: https://github.com/microsoft/vscode-dev-containers/blob/main/containers/go/.devcontainer/Dockerfile
ARG VARIANT="1"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# Install dependencies.
USER vscode
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1 && \
    go install github.com/go-task/task/v3/cmd/task@v3.18.0 && \
    go install github.com/jstemmer/go-junit-report/v2@latest