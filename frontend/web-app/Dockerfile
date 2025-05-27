FROM node:22 AS build

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm install

COPY . .
RUN npm run build --prod

FROM nginx:stable-alpine
COPY --from=build /app/dist/web-app/browser /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

#RUN rm -rf /usr/share/nginx/html/*
#
#COPY --from=build /app/dist/web-app/browser /usr/share/nginx/html
#
#EXPOSE 80
#CMD ["nginx", "-g", "daemon off;"]
