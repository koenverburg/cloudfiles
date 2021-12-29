FROM ubuntu:focal

SHELL ["/bin/bash", "-c"]

ENV DEBIAN_FRONTEND=noninteractive

LABEL Author = "Koen Verburg <creativekoen@gmail.com>"

RUN apt-get update && \
  apt-get install -y --no-install-recommends \
  apt-utils \
  software-properties-common

RUN apt-add-repository -y ppa:ansible/ansible

RUN apt-get install -y --no-install-recommends \
  # automake \
  build-essential \
  # pkg-config \
  # cargo \
  # cmake \
  # coreutils \
  # gcc \
  git \
  # golang \# add via ansible
  # libc6-dev \
  # libtool \
  # lua5.3-dev \# add via ansible
  make \
  # ninja-build \
  # nodejs \ # add via ansible
  # npm \
  python3 \
  # rustc \
  sudo \
  # gettext \
  # libtool-bin \
  # autoconf \
  # g++ \
  # unzip \
  # curl \
  # doxygen \
  # bat \
  # ripgrep \
  # tmux \ # add via ansible
  ansible

COPY . .

RUN sudo bash ./ubuntu.sh

CMD "/usr/bin/bash"
