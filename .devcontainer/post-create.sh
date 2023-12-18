#!/bin/bash

# Install global dependencies
NODE_MODULES="@angular/cli eslint"
npm install -g ${NODE_MODULES} && \
    npm cache clean --force

# install backend dependencies
cd backend
go mod download
go mod tidy

# install frontend dependencies
cd ../frontend
yarn install

ng config -g cli.analytics false && \
    ng config -g cli.completion.prompted true