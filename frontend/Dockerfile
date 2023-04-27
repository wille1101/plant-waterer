FROM node:19-alpine

WORKDIR /

ENV PATH /app/node_modules/.bin:$PATH
ENV PORT 3000

COPY package*.json ./

RUN npm install --silent
RUN npm install -g serve

COPY . ./

RUN npm run build

EXPOSE 3000

CMD ["serve", "-s", "build"]
