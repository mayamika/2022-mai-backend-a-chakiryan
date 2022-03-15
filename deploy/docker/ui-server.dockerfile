ARG NODE_VERSION=17-bullseye-slim
ARG NGINX_VERSION=1.21.6

FROM docker.io/node:${NODE_VERSION} as builder

COPY ui /ui/
WORKDIR /ui/
RUN npm install
RUN npm run build

FROM docker.io/nginx:${NGINX_VERSION}

# Clean existing configuration.
RUN rm /etc/nginx/conf.d/*

COPY --from=builder /ui/build /opt/app/
COPY /deploy/nginx/templates /etc/nginx/templates

EXPOSE 80
