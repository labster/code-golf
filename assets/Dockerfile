FROM alpine:edge

RUN apk add --no-cache bash curl gcc make musl-dev nodejs-npm openjdk8-jre perl

# Brotli
RUN curl -L https://github.com/google/brotli/archive/v1.0.2.tar.gz \
  | tar xzf -                                                      \
 && cd brotli-1.0.2                                                \
 && make -j`nproc`                                                 \
 && mv bin/brotli /usr/bin

# Closure Compiler
RUN curl http://dl.google.com/closure-compiler/compiler-20180101.tar.gz \
  | tar -zxf - -C /                                                     \
 && mv closure-compiler-v20180101.jar closure-compiler.jar              \
 && chmod +r closure-compiler.jar

# CSSO & SVGO
RUN npm install -g csso-cli@1.1.0 csso@3.4.0 svgo@1.0.3

# Bashisms FTW.
RUN ln -snf /bin/bash /bin/sh

WORKDIR /work

CMD ["assets/build"]
