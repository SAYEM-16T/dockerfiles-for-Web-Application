# Use Node.js LTS version
FROM node:18

# Set working directory
WORKDIR /app

# Copy package.json and package-lock.json first (better caching)
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the files
COPY . .

# Expose the application port
EXPOSE 8085

# Start the server
CMD ["node", "server.js"]
