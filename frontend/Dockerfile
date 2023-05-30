FROM node:18-alpine

WORKDIR /app
COPY ./package.json .
RUN apk --no-cache add --virtual builds-deps build-base python3
RUN npm install
COPY . .

ARG REACT_APP_SERVER_IP
ENV REACT_APP_SERVER_IP $REACT_APP_SERVER_IP

RUN chmod 777 /app/node_modules
EXPOSE 3000

USER node
CMD ["npm", "run", "start"]