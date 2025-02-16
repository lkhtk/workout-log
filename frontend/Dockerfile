FROM node:18-alpine AS build-stage

ARG NODE_ENV
ENV NODE_ENV $NODE_ENV
ARG VUE_APP_GOOGLE_CLIENT_ID
ENV VUE_APP_GOOGLE_CLIENT_ID $VUE_APP_GOOGLE_CLIENT_ID
WORKDIR /opt/app
COPY ./workout-log-frontend/package.json .
RUN npm install
COPY ./workout-log-frontend .
RUN npm run build

FROM nginxinc/nginx-unprivileged:alpine-perl AS production-stage
COPY --from=build-stage /opt/app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]