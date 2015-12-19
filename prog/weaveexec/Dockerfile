#seems like it it not so easy to build this image for example for an apline linux
#the the built docker-binary (last line) is built for raspbian jessie.
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
    util-linux
#    conntrack-tools
#    bind-tools
ADD ./weave ./sigproxy ./weaveproxy ./symlink /home/weave/
ADD ./netcheck ./docker_tls_args /usr/bin/
ADD ./weavewait /w/w
ADD ./weavewait_noop /w-noop/w
ADD ./weavewait_nomcast /w-nomcast/w
ADD ./docker.tgz /
