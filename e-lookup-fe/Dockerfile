FROM node:16-alpine 

WORKDIR /e-lookup-fe
COPY . .

ARG elookupbe_host=http://localhost:6002/
RUN echo VITE_ELOOKUPBE_HOST=$elookupbe_host > .env

#npm ci makes sure the exact versions in the lockfile gets installed
RUN npm ci 
RUN npm run build

EXPOSE 3000

CMD ["npx","serve","dist"]