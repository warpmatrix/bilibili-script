FROM ubuntu

ARG app=main
ARG dir=bin
ARG duration="20 0 * * *"
ARG configFile=config/config.yaml

RUN apt-get update && \
 apt-get install -y cron && \
 rm -rf /var/lib/apt/lists/* && \
 apt-get clean
COPY ${dir}/${app} /usr/bin/
COPY init.sh /init.sh
COPY ${configFile} /config/config.yaml
RUN echo "${duration} ${app} >> /tmp.log" | crontab -

ENTRYPOINT [ "./init.sh" ]
