FROM node:16-alpine

WORKDIR /app

RUN apk --no-cache add --virtual builds-deps build-base python3

COPY --chown=node:node ./package.json ./
RUN npm install

COPY --chown=node:node . ./
USER node
EXPOSE 3000

CMD ["npm", "run", "start"]