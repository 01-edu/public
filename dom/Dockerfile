FROM alpine:latest

WORKDIR /app
COPY . .
COPY subjects ./subjects
#RUN npm i
RUN apk add --no-cache \
    chromium \
    nss \
    freetype \
    harfbuzz \
    ca-certificates \
    ttf-freefont \
    nodejs \
    yarn

# Tell Puppeteer to skip installing Chrome. We'll be using the installed package.
ENV PUPPETEER_SKIP_CHROMIUM_DOWNLOAD=true \
    PUPPETEER_EXECUTABLE_PATH=/usr/bin/chromium-browser
RUN apk add --update nodejs npm
RUN npm i
RUN npm i puppeteer
#RUN yarn add puppeteer@15.3.2
ENTRYPOINT ["/app/entrypoint.sh"]
