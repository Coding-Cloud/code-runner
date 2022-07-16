FROM nospy/pnpm:16
RUN pnpm config set store-dir=/root/.local/share/pnpm

RUN mkdir /home/app

WORKDIR /home/app

COPY ./start.sh /home/start.sh
COPY ./install-dependencies.sh /home/install-dependencies.sh
RUN chmod +x /home/start.sh
RUN chmod +x /home/install-dependencies.sh

ENV HOST=0.0.0.0
ENV PROJECT_ID=""
ENV PUBLIC_URL="localhost"
ENV PORT=443
ENV SCRIPTS_PATH=/home

EXPOSE 443
VOLUME /home/app
VOLUME /logs/$PROJECT_ID.log
ENTRYPOINT ["sh", "/home/react-starter.sh"]
