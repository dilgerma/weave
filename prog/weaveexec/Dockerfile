FROM resin/rpi-raspbian:jessie 

MAINTAINER Weaveworks Inc <help@weave.works>
LABEL works.weave.role=system

WORKDIR /home/weave
VOLUME /w
VOLUME /w-noop
VOLUME /w-nomcast
ENTRYPOINT ["/home/weave/sigproxy", "/home/weave/weave"]
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
