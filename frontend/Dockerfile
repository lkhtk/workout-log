FROM node:lts-alpine AS build-stage
WORKDIR /workout-log-frontend
COPY ./workout-log-frontend/package*.json ./
RUN npm install --ignore-scripts
COPY ./workout-log-frontend/ ./
RUN npm run build --ignore-scripts

FROM nginxinc/nginx-unprivileged:alpine-perl AS production-stage
COPY --from=build-stage /workout-log-frontend/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]