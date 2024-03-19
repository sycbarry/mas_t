# Use a base image with Maven and OpenJDK 17 installed
FROM maven:3.8.4-openjdk-17 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Maven project's source code into the container
COPY . .

# Build the Maven project
RUN mvn clean package

# Use a lightweight base image to run the application
FROM openjdk:17-jdk-alpine

# Set the working directory inside the container
WORKDIR /app

EXPOSE 8080

ENTRYPOINT ["mvnw" , "spring-boot:run"]
