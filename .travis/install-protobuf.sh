#!/bin/bash
set -e

if [[ ! -x "${HOME}/protobuf-${PROTOBUF_VERSION}/bin/protoc" ]]; then
  curl -Lo /tmp/protobuf-${PROTOBUF_VERSION}.tar.gz https://github.com/google/protobuf/archive/v${PROTOBUF_VERSION}.tar.gz
  tar -C /tmp -xf /tmp/protobuf-${PROTOBUF_VERSION}.tar.gz
  pushd /tmp/protobuf-${PROTOBUF_VERSION}
  ./autogen.sh
  ./configure --prefix=$HOME/protobuf-${PROTOBUF_VERSION}
  make
  make install
  popd
else
  echo -e "Found ${HOME}/protobuf-${PROTOBUF_VERSION} using it from cache"
fi
