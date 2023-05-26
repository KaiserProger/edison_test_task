FROM node:18-alpine as builder
WORKDIR /frontend_app
COPY package.json /frontend_app/package.json
RUN apk --no-cache add --virtual builds-deps build-base python
RUN npm install
COPY . /frontend_app
EXPOSE 80
USER node
CMD ["npm", "run", "start"]