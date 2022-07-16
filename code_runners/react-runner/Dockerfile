FROM nospy/pnpm:16

RUN mkdir /home/app

WORKDIR /home/app

COPY ./react-starter.sh /home/react-starter.sh
RUN chmod +x /home/react-starter.sh

ENV HOST=0.0.0.0
ENV PROJECT_ID=""
ENV PUBLIC_URL="localhost"
ENV PORT=443

EXPOSE 443
VOLUME /home/app
VOLUME /logs/$PROJECT_ID.log
ENTRYPOINT ["sh", "/home/react-starter.sh"]
