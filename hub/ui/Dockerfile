# Use an official Node.js runtime as a parent image
FROM node:14-alpine as build

# Set the working directory in the container
WORKDIR /app

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the entire project to the working directory
COPY . .

# Build the React app
RUN npm run build

# Use Node.js HTTP server to serve the static files
FROM node:14-alpine

# Set the working directory in the container
WORKDIR /app

# Copy the build output from the previous stage to the working directory
COPY --from=build /app/build .

# Expose port 3000 to the outside world
EXPOSE 3000

# Start the Node.js server when the container starts
CMD ["npx", "http-server", "-p", "3000"]

