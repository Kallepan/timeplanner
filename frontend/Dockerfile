### Build stage ###
FROM node:18-alpine as builder

WORKDIR /app

# Install dependencies
RUN --mount=type=bind,source=package.json,target=/app/package.json \
    --mount=type=bind,source=yarn.lock,target=/app/yarn.lock \
    --mount=type=cache,target=/root/.npm \
    yarn install --frozen-lockfile

# Build the application
RUN --mount=type=bind,source=package.json,target=/app/package.json \
    --mount=type=bind,source=yarn.lock,target=/app/yarn.lock \
    --mount=type=bind,source=src,target=/app/src \
    --mount=type=bind,source=angular.json,target=/app/angular.json \
    --mount=type=bind,source=tsconfig.json,target=/app/tsconfig.json \
    --mount=type=bind,source=tsconfig.app.json,target=/app/tsconfig.app.json \
    --mount=type=cache,target=/root/.npm \
    yarn run build --configuration=production

### Final stage ###
FROM nginx:alpine

# Copy the binary file from the builder stage
COPY --from=builder /app/dist/frontend/browser /usr/share/nginx/html

# Copy nginx configuration
COPY ./nginx.conf /etc/nginx/conf.d/default.conf

# Expose port 80
EXPOSE 80
