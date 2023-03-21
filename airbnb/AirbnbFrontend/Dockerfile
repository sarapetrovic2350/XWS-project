#stage 1
FROM node:latest as builder

# Setting up the work directory
WORKDIR /angular-app

# Add the source code to app
COPY ./ /angular-app

# Install all the dependecies
RUN npm install

# Building our application
RUN npm run build

#stage 2
# Fetching the latest nginx image
FROM nginx:latest

# Copying our nginx.conf
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Copying built assets from builder to replace the default nginx contents
COPY --from=builder /angular-app/dist/airbnb-frontend /usr/share/nginx/html/airbnb-frontend

# Expose port
EXPOSE 4200
