FROM node:20-alpine AS builder

WORKDIR /app

ARG PUBLIC_BACKEND_URL 
ENV PUBLIC_BACKEND_URL=$PUBLIC_BACKEND_URL

COPY package*.json .
RUN npm install --verbose

COPY .npmrc .
COPY postcss.config.js .
COPY svelte.config.js .
COPY tailwind.config.ts .
COPY tsconfig.json .
COPY vite.config.ts .
COPY .svelte-kit .

COPY ./static /app/static
COPY ./src /app/src

RUN npm run build

CMD [ "node", "build" ]
