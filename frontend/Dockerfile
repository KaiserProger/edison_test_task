FROM node:18-alpine
WORKDIR /frontend
COPY package.json /frontend/package.json
RUN apk --no-cache add --virtual builds-deps build-base python3
RUN npm install
COPY . /frontend
COPY /public /frontend/public
RUN ls .
EXPOSE 80
USER node
CMD ["npm", "run", "start"]