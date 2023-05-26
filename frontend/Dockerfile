FROM node:18-alpine
WORKDIR /frontend_app
COPY package.json /frontend_app/package.json
RUN apk --no-cache add --virtual builds-deps build-base python3.10
RUN npm install
COPY . /frontend_app
EXPOSE 80
USER node
CMD ["npm", "run", "start"]