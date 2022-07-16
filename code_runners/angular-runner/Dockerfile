FROM nospy/pnpm:16
RUN pnpm config set store-dir=/root/.local/share/pnpm
RUN pnpm install -g @angular/cli

RUN mkdir /home/app
WORKDIR /home/app

COPY ./start.sh /home/start.sh
COPY ./install-dependencies.sh /home/install-dependencies.sh
RUN chmod +x /home/start.sh
RUN chmod +x /home/install-dependencies.sh

ENV HOST=0.0.0.0
ENV PUBLIC_URL="localhost"
ENV PROJECT_ID=""
ENV PORT=4200
ENV SCRIPTS_PATH=/home

EXPOSE 4200
VOLUME /home/app
VOLUME /logs/$PROJECT_ID.log
VOLUME /root/.local/share/pnpm
ENTRYPOINT ["sh", "/home/angular-starter.sh"]
