FROM golang:1.13.4 as builder

RUN mkdir -p /go/src/hardware
RUN mkdir -p /go/build

COPY . /go/src/hardware/

ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.cn"
RUN go env
RUN go mod verify
RUN (cd /go/src/hardware && go build -o /go/build/hardware)
RUN (cp -rf  /go/src/hardware/public /go/build/)
RUN (mv /go/build  /go/hardware-1.0)
RUN (cd /go &&  tar zcf hardware-1.0.tar.gz hardware-1.0)
RUN (tar tvf /go/hardware-1.0.tar.gz)

FROM centos:7

RUN yum -y install \
           gcc \
           tree \
           make \
           rpm-build

RUN mkdir -p /root/rpmbuild/{BUILD,BUILDROOT,RPMS,SOURCES,SPECS,SRPMS}

COPY .rpmmacros /root/.rpmmacros
COPY --from=builder /go/hardware-1.0.tar.gz /root/rpmbuild/SOURCES/

COPY hardware.spec /root/rpmbuild/SPECS/hardware.spec
RUN rpmbuild -ba /root/rpmbuild/SPECS/hardware.spec 
RUN rpm  -ivh /root/rpmbuild/RPMS/x86_64/hardware-1.0-1.el7.x86_64.rpm
RUN /opt/hardware/bin/hardware