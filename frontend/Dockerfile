FROM node:lts-alpine AS build-stage

ARG NODE_ENV
ENV NODE_ENV $NODE_ENV
ARG VUE_APP_GOOGLE_CLIENT_ID
ENV VUE_APP_GOOGLE_CLIENT_ID $VUE_APP_GOOGLE_CLIENT_ID
ADD ./workout-log-frontend/package.json /tmp/package.json
RUN cd /tmp && npm install --ignore-scripts
RUN mkdir -p /opt/app && cp -a /tmp/node_modules /opt/app/
WORKDIR /opt/app
ADD ./workout-log-frontend /opt/app
RUN npm run build --ignore-scripts

FROM nginxinc/nginx-unprivileged:alpine-perl AS production-stage
COPY --from=build-stage /opt/app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]