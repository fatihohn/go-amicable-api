# Use an official Go runtime as a parent image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY ./app /app
COPY .env /app/.env

# TODO. replace schema.sql env variables

# Install any dependencies you might need
# For example, if you need PostgreSQL client tools
RUN apt-get update && apt-get install -y postgresql-client

# Build the Go application (replace with your actual build command)
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./scripts/docker-entrypoint.sh"]
