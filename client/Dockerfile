# directory ./client/Dockerfile
FROM node:18-alpine

WORKDIR /app

COPY . .

RUN npm ci

RUN npm run build

EXPOSE 4173

CMD ["npm", "run", "preview"]