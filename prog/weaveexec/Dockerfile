FROM resin/rpi-raspbian:jessie

MAINTAINER Weaveworks Inc <help@weave.works>
LABEL works.weave.role=system

WORKDIR /home/weave
ENTRYPOINT ["/home/weave/sigproxy", "/home/weave/weave"]

RUN echo "deb http://archive.raspbian.org/raspbian stretch main contrib non-free" >> /etc/apt/sources.list
RUN apt-get update 
RUN apt-get install -y \
    curl \
    ethtool \
    iptables \
    iproute2 \
    util-linux \ 
#    conntrack-tools \
#    bind-tools \
  && rm -rf /var/cache/apt/*

ADD ./weave ./sigproxy ./weaveproxy ./symlink /home/weave/
ADD ./weaveutil /usr/bin/
ADD ./weavewait /w/w
ADD ./weavewait_noop /w-noop/w
ADD ./weavewait_nomcast /w-nomcast/w
ADD ./docker.tgz /
