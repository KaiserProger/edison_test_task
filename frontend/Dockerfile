FROM node:16-alpine

WORKDIR /app

RUN apk --no-cache add --virtual builds-deps build-base python3

COPY --chown=node:node ./package.json .
RUN npm install

COPY --chown=node:node . .
RUN chmod 777 /app/node_modules
EXPOSE 3000

USER node
CMD ["npm", "run", "start"]