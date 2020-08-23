FROM node:14.7.0

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

COPY package.json /usr/src/app/
ENV NODE_ENV development
RUN npm install
