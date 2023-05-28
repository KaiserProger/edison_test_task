FROM node:16-alpine

WORKDIR /app

RUN apk --no-cache add --virtual builds-deps build-base python3

COPY ./package.json ./
RUN npm install

RUN mkdir -p node_modules/.cache && chmod -R 777 node_modules/.cache

COPY . ./
EXPOSE 3000
USER node
CMD ["npm", "run", "start"]