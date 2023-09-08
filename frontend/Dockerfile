FROM node
WORKDIR /app
COPY package.json .
RUN npm i
COPY . .
## EXPOSE Port in vite.config file
EXPOSE 5173
CMD ["npm", "run", "dev"]