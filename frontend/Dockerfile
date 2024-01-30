### Build stage ###
FROM node:18-alpine as builder

WORKDIR /app

# Install dependencies
COPY ./package.json ./yarn.lock ./angular.json ./
RUN yarn install --frozen-lockfile

# Copy the source code
COPY ./src ./src
COPY ./angular.json ./tsconfig.json tsconfig.app.json ./

# Build the application
RUN yarn run build --configuration=production

### Final stage ###
FROM nginx:alpine

# Copy the binary file from the builder stage
COPY --from=builder /app/dist/frontend/browser /usr/share/nginx/html

# Copy nginx configuration
COPY ./nginx.conf /etc/nginx/conf.d/default.conf

# Expose port 80
EXPOSE 80
