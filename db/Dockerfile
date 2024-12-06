FROM mongo:latest

# Set the environment variables for MongoDB
ENV MONGO_INITDB_ROOT_USERNAME=user
ENV MONGO_INITDB_ROOT_PASSWORD=password
ENV MONGO_INITDB_DATABASE=newDB

# Copy the initialization JS file into the container
COPY ./init/data.js /docker-entrypoint-initdb.d/data.js