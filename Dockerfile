FROM node:20

WORKDIR /app

COPY package*.json /

RUN npm install

COPY . .

EXPOSE 3000

COPY prisma ./prisma
RUN npx prisma generate

CMD ["npm", "run" ,"start"]