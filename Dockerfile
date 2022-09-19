FROM golang:1.17-alpine
RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/TyDataSci/Flowershop-GoBackend/cmd
RUN cd /build && git clone https://github.com/TyDataSci/Flowershop-GoBackend.git
RUN cd /build/Flowershop-GoBackend/cmd && go build

EXPOSE 8000

ENTRYPOINT [ "/bulid/Flowershop-GoBackend/cmd/main" ]