FROM node:16-alpine 

WORKDIR /e-lookup-fe
COPY . .

#npm ci makes sure the exact versions in the lockfile gets installed
RUN npm ci 
RUN npm run build

ENV VITE_ELOOKUPBE_HOST=http://localhost:3000/
EXPOSE 3000

CMD ["npx","serve","dist"]