#!/bin/bash

# Install global dependencies
NODE_MODULES="@angular/cli"
npm install -g ${NODE_MODULES} && \
    npm cache clean --force

# install auth backend dependencies
cd auth-backend
go mod download
go mod tidy

# install planner backend dependencies
cd ../planner-backend
go mod download
go mod tidy

# install frontend dependencies
cd ../frontend
yarn install

ng config -g cli.analytics false && \
    ng config -g cli.completion.prompted true

# install go tools
go install github.com/google/wire/cmd/wire@latest