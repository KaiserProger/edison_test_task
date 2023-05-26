FROM node:18-alpine

WORKDIR /frontend

COPY ./package.json ./
COPY ./package.json /frontend/
COPY ./package-lock.json ./
COPY ./package-lock.json /frontend/
RUN apk --no-cache add --virtual builds-deps build-base python3
RUN npm install
COPY . ./
EXPOSE 80
USER node
CMD ["npm", "run", "start"]