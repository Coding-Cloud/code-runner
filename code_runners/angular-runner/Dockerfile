FROM nospy/pnpm:16
RUN pnpm install -g @angular/cli

RUN mkdir /home/app
WORKDIR /home/app

COPY ./angular-starter.sh /home/angular-starter.sh
RUN chmod +x /home/angular-starter.sh

ENV HOST=0.0.0.0
ENV PUBLIC_URL="localhost"
ENV PROJECT_ID=""
ENV PORT=4200

EXPOSE 4200
VOLUME /home/app
VOLUME /logs/$PROJECT_ID.log
VOLUME /root/.local/share/pnpm
ENTRYPOINT ["sh", "/home/angular-starter.sh"]
