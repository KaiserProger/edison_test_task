FROM node:18-alpine as builder
WORKDIR /frontend_app
COPY package.json /frontend_app/package.json
RUN apk --no-cache add --virtual builds-deps build-base python
RUN npm install
COPY . /frontend_app
RUN npm run build
FROM node:18-alpine
WORKDIR /frontend_app
COPY --from=builder /app/dist /frontend_app
COPY package.json /frontend_app/package.json
RUN apk --no-cache add --virtual builds-deps build-base python
RUN npm install --only=prod
EXPOSE 80
USER node
CMD ["node", "index.js"]