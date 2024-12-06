FROM node:18.13.0

# Installing server modules
WORKDIR /usr/src/app
USER root 
# RUN chown -R 1000:1000 "/root/.npm"

COPY package*.json ./
RUN npm install

# RUN chown -R node ./node_module
# RUN chown -R 1000:1000 /usr/src/app/node_module


USER root
# CMD ["npm", "run", "dev"]