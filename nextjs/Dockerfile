###################
# BUILD FOR LOCAL DEVELOPMENT
###################

FROM node:18-alpine As development

WORKDIR /var/www/nextjs-app

COPY package*.json ./

RUN npm ci
#RUN npm install

###################
# BUILD FOR PRODUCTION
###################

FROM node:18-alpine As build

WORKDIR /var/www/nextjs-app

COPY package*.json ./

COPY --from=development /var/www/nextjs-app/node_modules ./node_modules 

COPY . .

RUN npm run build

# if your NODE_ENV is set to production, then it avoids installing dev dependencies.
ENV NODE_ENV production

#RUN npm install --omit=dev && npm cache clean --force
RUN npm ci --omit=dev && npm cache clean --force
#RUN npm ci

###################
# PRODUCTION
###################

FROM node:18-alpine As production

WORKDIR /var/www/nextjs-app

# ENV NODE_ENV production

# Copy only used files to the production image
# COPY --from=build /var/www/nextjs-app/node_modules ./node_modules
# COPY --from=build /var/www/nextjs-app/package*.json ./
COPY --from=build /var/www/nextjs-app/.next ./.next

EXPOSE 3000

CMD [ "npm", "start" ]