# ---- Stage 1: The Build Environment ----
# Use a Node.js image to build our Svelte app. 'alpine' is a small version.
FROM node:20-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json first to leverage Docker's cache.
# 'npm install' will only run again if these files change.
COPY package*.json ./
RUN npm install

# Copy the rest of the application source code
COPY . .

# Run the production build command defined in package.json
RUN npm run build

# ---- Stage 2: The Production Environment ----
# Now, start from a fresh, tiny Nginx web server image.
FROM nginx:stable-alpine

# Copy ONLY the built assets from the 'builder' stage.
# The result of 'npm run build' is in the '/app/build' directory.
# We copy it to the directory Nginx serves files from by default.
COPY --from=builder /app/build /usr/share/nginx/html

# Tell Docker that the container listens on port 80 (the default Nginx port)
EXPOSE 80

# The default command for the nginx image is to start the server.
# This command keeps it running in the foreground, which is what Docker wants.
CMD ["nginx", "-g", "daemon off;"]