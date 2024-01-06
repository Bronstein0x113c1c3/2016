FROM janpfeifer/gonb_jupyterlab:latest

USER root

RUN apt-get update
# RUN apt-get install gcc libgl1-mesa-dev xorg-dev -y

ENTRYPOINT ["jupyter","notebook","--allow-root"] 