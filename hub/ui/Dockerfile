# Use an official Node.js runtime as a parent image
FROM node:16.18
WORKDIR /app
COPY package.json ./
WORKDIR /app
COPY . . 
EXPOSE 3000
CMD ["npm", "run", "watch"]

