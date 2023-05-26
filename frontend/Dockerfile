FROM node:18-alpine

WORKDIR /frontend

RUN apk --no-cache add --virtual builds-deps build-base python3

COPY ./package.json .
RUN npm install
COPY . ./
COPY ./package.json .
RUN ls -la .
EXPOSE 3000
USER node
CMD ["npm", "start"]