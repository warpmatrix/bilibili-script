FROM ubuntu
COPY bin/main /usr/bin
ENTRYPOINT main
# CMD /bin/bash
