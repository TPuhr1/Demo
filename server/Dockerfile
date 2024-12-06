# FROM node:18.13.0

# # Installing server modules
# WORKDIR /usr/src/app/server
# # COPY package.json ./
# # RUN npm install

# # Command to start the server
# # USER root
# # CMD ["go", "run", "main.go"]

# Use an official Go runtime as a base image
FROM golang:1.20-alpine

# Set up the working directory
WORKDIR /usr/src/app/server

# Copy the Go source code into the container
COPY . .

# Expose the port the server runs on
EXPOSE 12000

# Command to run the Go application
# CMD ["go", "run", "main.go"]