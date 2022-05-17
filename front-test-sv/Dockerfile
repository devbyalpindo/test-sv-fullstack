FROM node:16-alpine AS deps

WORKDIR /app

COPY . /app

RUN npm install

RUN npm install -g serve

RUN npm run build

EXPOSE 5000

# Run application
CMD ["serve", "-l", "5000", "-s" ,"build"]
# CMD serve -s build
