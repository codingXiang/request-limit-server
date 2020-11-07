# 快取
FROM    codingxiang/go_vc AS build_base
ENV     RUN_PATH=/server PROJ_PATH=/build
RUN     mkdir -p $RUN_PATH
WORKDIR $RUN_PATH
ENV     GO111MODULE=on
COPY    go.mod .
COPY    go.sum .
RUN     go mod download

# Build 專案
FROM    build_base AS builder
LABEL   maintainer="賴念翔"
USER    root
ADD     . ${PROJ_PATH}
WORKDIR ${PROJ_PATH}
RUN     make build pack \
        && tar -zxf app-v*.tar.gz -C ${RUN_PATH} \
        && rm -rf ${PROJ_PATH}

# 打包 Image
FROM    alpine
LABEL   maintainer="賴念翔"
USER    root
ENV     RUN_PATH=/server
RUN     mkdir -p $RUN_PATH && apk add --no-cache ca-certificates bash
COPY    --from=builder ${RUN_PATH} ${RUN_PATH}
WORKDIR ${RUN_PATH}
ENTRYPOINT ["./app"]
