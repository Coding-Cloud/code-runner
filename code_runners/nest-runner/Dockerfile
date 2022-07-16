FROM nospy/pnpm:16
RUN pnpm config set store-dir=/root/.local/share/pnpm
RUN pnpm install -g @nestjs/cli

RUN mkdir /home/app
WORKDIR /home/app

COPY ./start.sh /home/start.sh
COPY ./install-dependencies.sh /home/install-dependencies.sh
RUN chmod +x /home/start.sh
RUN chmod +x /home/install-dependencies.sh

ENV SCRIPTS_PATH=/home

EXPOSE 3000
VOLUME /home/app
VOLUME /logs/$PROJECT_ID.log
VOLUME /root/.local/share/pnpm
ENTRYPOINT ["sh", "/home/nest-starter.sh"]
