FROM node:9

RUN mkdir -p /usr/gofcc
WORKDIR /usr/gofcc

ARG NODE_ENV
ENV NODE_ENV $NODE_ENV
COPY . /usr/gofcc/
RUN npm install && npm install -g @angular/cli --unsafe

EXPOSE 4200/tcp

CMD ["npm", "start", "--", "--host", "0.0.0.0", "--poll", "500"]