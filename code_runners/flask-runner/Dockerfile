FROM python:3.9.13-slim
RUN pip install Flask --target=/libs

RUN mkdir /home/app
WORKDIR /home/app

COPY ./flask-starter.sh /home/flask-starter.sh
RUN chmod +x /home/flask-starter.sh

ENV FLASK_APP main.py
ENV FLASK_ENV development
ENV PYTHONPATH /libs
ENV PORT 5000
ENV HOST 0.0.0.0

EXPOSE 5000
VOLUME /home/app
VOLUME /logs/$PROJECT_ID.log
VOLUME /libs
ENTRYPOINT ["sh", "/home/flask-starter.sh"]
