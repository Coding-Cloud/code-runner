FROM nospy/pnpm:16
RUN pnpm install -g @nestjs/cli

RUN mkdir /home/app
WORKDIR /home/app

COPY ./nest-starter.sh /home/nest-starter.sh
RUN chmod +x /home/nest-starter.sh

EXPOSE 3000
VOLUME /home/app
VOLUME /logs/$PROJECT_ID.log
VOLUME /root/.local/share/pnpm
ENTRYPOINT ["sh", "/home/nest-starter.sh"]
