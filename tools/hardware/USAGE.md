# Demo for RPM package

```bash
duanp-a01:hardware duanp$ docker build . -t hw
Sending build context to Docker daemon  19.46kB
Step 1/24 : FROM golang:1.13.4 as builder
 ---> a2e245db8bd3
Step 2/24 : RUN mkdir -p /go/src/hardware
 ---> Using cache
 ---> 0a59435dffa0
Step 3/24 : RUN mkdir -p /go/build
 ---> Using cache
 ---> 351fe3768f19
Step 4/24 : COPY . /go/src/hardware/
 ---> f828b1cc2f6b
Step 5/24 : ENV GOOS=linux
 ---> Running in a5ed47cdaeac
Removing intermediate container a5ed47cdaeac
 ---> 5407472718ea
Step 6/24 : ENV GOARCH=amd64
 ---> Running in 6bdff9c9e978
Removing intermediate container 6bdff9c9e978
 ---> 3ff268e923b8
Step 7/24 : ENV GO111MODULE=on
 ---> Running in bd91df97a0e5
Removing intermediate container bd91df97a0e5
 ---> 79c7c11624e9
Step 8/24 : ENV GOPROXY="https://goproxy.cn"
 ---> Running in 795d12096593
Removing intermediate container 795d12096593
 ---> 04f8ea83a895
Step 9/24 : RUN go env
 ---> Running in 26bf4d79c5cb
GO111MODULE="on"
GOARCH="amd64"
GOBIN=""
GOCACHE="/root/.cache/go-build"
GOENV="/root/.config/go/env"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GONOPROXY=""
GONOSUMDB=""
GOOS="linux"
GOPATH="/go"
GOPRIVATE=""
GOPROXY="https://goproxy.cn"
GOROOT="/usr/local/go"
GOSUMDB="sum.golang.org"
GOTMPDIR=""
GOTOOLDIR="/usr/local/go/pkg/tool/linux_amd64"
GCCGO="gccgo"
AR="ar"
CC="gcc"
CXX="g++"
CGO_ENABLED="1"
GOMOD="/dev/null"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build903003820=/tmp/go-build -gno-record-gcc-switches"
Removing intermediate container 26bf4d79c5cb
 ---> 14d679ee26a1
Step 10/24 : RUN go mod verify
 ---> Running in 55ef407ffcce
all modules verified
Removing intermediate container 55ef407ffcce
 ---> d81bf94b0485
Step 11/24 : RUN (cd /go/src/hardware && go build -o /go/build/hardware)
 ---> Running in 72e54fd8b912
go: finding github.com/shirou/gopsutil v2.19.12+incompatible
go: downloading github.com/shirou/gopsutil v2.19.12+incompatible
go: extracting github.com/shirou/gopsutil v2.19.12+incompatible
go: finding golang.org/x/sys latest
go: downloading golang.org/x/sys v0.0.0-20200116001909-b77594299b42
go: extracting golang.org/x/sys v0.0.0-20200116001909-b77594299b42
Removing intermediate container 72e54fd8b912
 ---> 8f11c9c20cea
Step 12/24 : RUN (cp -rf  /go/src/hardware/public /go/build/)
 ---> Running in 65b0f4cccbb4
Removing intermediate container 65b0f4cccbb4
 ---> 48a1b7e765cf
Step 13/24 : RUN (mv /go/build  /go/hardware-1.0)
 ---> Running in 965c544a4b84
Removing intermediate container 965c544a4b84
 ---> 007271dbeff1
Step 14/24 : RUN (cd /go &&  tar zcf hardware-1.0.tar.gz hardware-1.0)
 ---> Running in 36b99fd1d231
Removing intermediate container 36b99fd1d231
 ---> 69d93a299d94
Step 15/24 : RUN (tar tvf /go/build/hardware-1.0.tar.gz)
 ---> Running in 09f82764a1aa
tar: /go/build/hardware-1.0.tar.gz: Cannot open: No such file or directory
tar: Error is not recoverable: exiting now
The command '/bin/sh -c (tar tvf /go/build/hardware-1.0.tar.gz)' returned a non-zero code: 2
duanp-a01:hardware duanp$ docker build . -t hw
Sending build context to Docker daemon  19.46kB
Step 1/24 : FROM golang:1.13.4 as builder
 ---> a2e245db8bd3
Step 2/24 : RUN mkdir -p /go/src/hardware
 ---> Using cache
 ---> 0a59435dffa0
Step 3/24 : RUN mkdir -p /go/build
 ---> Using cache
 ---> 351fe3768f19
Step 4/24 : COPY . /go/src/hardware/
 ---> 8de6e686bd4b
Step 5/24 : ENV GOOS=linux
 ---> Running in d76e7fb48e16
Removing intermediate container d76e7fb48e16
 ---> ab47dfb741c9
Step 6/24 : ENV GOARCH=amd64
 ---> Running in ca7e31d689dc
Removing intermediate container ca7e31d689dc
 ---> 6fc341e43316
Step 7/24 : ENV GO111MODULE=on
 ---> Running in e1282870c8ba
Removing intermediate container e1282870c8ba
 ---> 13c214eeb301
Step 8/24 : ENV GOPROXY="https://goproxy.cn"
 ---> Running in c968c7f96d56
Removing intermediate container c968c7f96d56
 ---> 377c3ef9a5cc
Step 9/24 : RUN go env
 ---> Running in 0d6140486b8c
GO111MODULE="on"
GOARCH="amd64"
GOBIN=""
GOCACHE="/root/.cache/go-build"
GOENV="/root/.config/go/env"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GONOPROXY=""
GONOSUMDB=""
GOOS="linux"
GOPATH="/go"
GOPRIVATE=""
GOPROXY="https://goproxy.cn"
GOROOT="/usr/local/go"
GOSUMDB="sum.golang.org"
GOTMPDIR=""
GOTOOLDIR="/usr/local/go/pkg/tool/linux_amd64"
GCCGO="gccgo"
AR="ar"
CC="gcc"
CXX="g++"
CGO_ENABLED="1"
GOMOD="/dev/null"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build971521368=/tmp/go-build -gno-record-gcc-switches"
Removing intermediate container 0d6140486b8c
 ---> 613071eb5b4e
Step 10/24 : RUN go mod verify
 ---> Running in 97e87e73c12e
all modules verified
Removing intermediate container 97e87e73c12e
 ---> f71c48193185
Step 11/24 : RUN (cd /go/src/hardware && go build -o /go/build/hardware)
 ---> Running in 26b120f65db3
go: finding github.com/shirou/gopsutil v2.19.12+incompatible
go: downloading github.com/shirou/gopsutil v2.19.12+incompatible
go: extracting github.com/shirou/gopsutil v2.19.12+incompatible
go: finding golang.org/x/sys latest
go: downloading golang.org/x/sys v0.0.0-20200116001909-b77594299b42
go: extracting golang.org/x/sys v0.0.0-20200116001909-b77594299b42
Removing intermediate container 26b120f65db3
 ---> 663e43ced94b
Step 12/24 : RUN (cp -rf  /go/src/hardware/public /go/build/)
 ---> Running in 598877630bc6
Removing intermediate container 598877630bc6
 ---> 61801339adf1
Step 13/24 : RUN (mv /go/build  /go/hardware-1.0)
 ---> Running in 9110886c796a
Removing intermediate container 9110886c796a
 ---> 298bf3a0dbb0
Step 14/24 : RUN (cd /go &&  tar zcf hardware-1.0.tar.gz hardware-1.0)
 ---> Running in b3500e74767f
Removing intermediate container b3500e74767f
 ---> 6c4e3946bd36
Step 15/24 : RUN (tar tvf /go/hardware-1.0.tar.gz)
 ---> Running in 3f7a029ac4cc
drwxr-xr-x root/root         0 2020-01-16 15:06 hardware-1.0/
-rwxr-xr-x root/root   2905567 2020-01-16 15:06 hardware-1.0/hardware
drwxr-xr-x root/root         0 2020-01-16 15:06 hardware-1.0/public/
-rw-r--r-- root/root      9266 2020-01-16 15:06 hardware-1.0/public/home.html
Removing intermediate container 3f7a029ac4cc
 ---> df42b8d92aa1
Step 16/24 : FROM centos:7
 ---> 5e35e350aded
Step 17/24 : RUN yum -y install            gcc            tree            make            rpm-build
 ---> Using cache
 ---> a387cb7d2305
Step 18/24 : RUN mkdir -p /root/rpmbuild/{BUILD,BUILDROOT,RPMS,SOURCES,SPECS,SRPMS}
 ---> Using cache
 ---> 4bffdea96839
Step 19/24 : COPY .rpmmacros /root/.rpmmacros
 ---> Using cache
 ---> 99799c9b5149
Step 20/24 : COPY --from=builder /go/hardware-1.0.tar.gz /root/rpmbuild/SOURCES/
 ---> 9438a6dd2f53
Step 21/24 : COPY hardware.spec /root/rpmbuild/SPECS/hardware.spec
 ---> 5e10c00023bf
Step 22/24 : RUN rpmbuild -ba /root/rpmbuild/SPECS/hardware.spec
 ---> Running in 634a9eac5487
Executing(%prep): /bin/sh -e /var/tmp/rpm-tmp.FyVQLv
+ umask 022
+ cd /root/rpmbuild/BUILD
+ cd /root/rpmbuild/BUILD
+ rm -rf hardware-1.0
+ /usr/bin/gzip -dc /root/rpmbuild/SOURCES/hardware-1.0.tar.gz
+ /usr/bin/tar -xf -
+ STATUS=0
+ '[' 0 -ne 0 ']'
+ cd hardware-1.0
+ /usr/bin/chmod -Rf a+rX,u+w,g-w,o-w .
+ exit 0
Executing(%install): /bin/sh -e /var/tmp/rpm-tmp.rt5DiM
+ umask 022
+ cd /root/rpmbuild/BUILD
+ '[' /root/rpmbuild/BUILDROOT/hardware-1.0-1.el7.x86_64 '!=' / ']'
+ rm -rf /root/rpmbuild/BUILDROOT/hardware-1.0-1.el7.x86_64
++ dirname /root/rpmbuild/BUILDROOT/hardware-1.0-1.el7.x86_64
+ mkdir -p /root/rpmbuild/BUILDROOT
+ mkdir /root/rpmbuild/BUILDROOT/hardware-1.0-1.el7.x86_64
+ cd hardware-1.0
+ rm -rf /root/rpmbuild/BUILDROOT/hardware-1.0-1.el7.x86_64
+ mkdir -p /root/rpmbuild/BUILDROOT/hardware-1.0-1.el7.x86_64//opt/hardware/bin
+ cp hardware /root/rpmbuild/BUILDROOT/hardware-1.0-1.el7.x86_64//opt/hardware/bin/
+ mkdir -p /root/rpmbuild/BUILDROOT/hardware-1.0-1.el7.x86_64//opt/hardware/public
+ cp public/home.html /root/rpmbuild/BUILDROOT/hardware-1.0-1.el7.x86_64//opt/hardware/public/
+ /usr/lib/rpm/check-buildroot
+ /usr/lib/rpm/redhat/brp-compress
+ /usr/lib/rpm/redhat/brp-strip /usr/bin/strip
+ /usr/lib/rpm/redhat/brp-strip-comment-note /usr/bin/strip /usr/bin/objdump
+ /usr/lib/rpm/redhat/brp-strip-static-archive /usr/bin/strip
+ /usr/lib/rpm/brp-python-bytecompile /usr/bin/python 1
+ /usr/lib/rpm/redhat/brp-python-hardlink
+ /usr/lib/rpm/redhat/brp-java-repack-jars
Processing files: hardware-1.0-1.el7.x86_64
Provides: hardware = 1.0-1.el7 hardware(x86-64) = 1.0-1.el7
Requires(rpmlib): rpmlib(CompressedFileNames) <= 3.0.4-1 rpmlib(FileDigests) <= 4.6.0-1 rpmlib(PayloadFilesHavePrefix) <= 4.0-1
Checking for unpackaged file(s): /usr/lib/rpm/check-files /root/rpmbuild/BUILDROOT/hardware-1.0-1.el7.x86_64
Wrote: /root/rpmbuild/SRPMS/hardware-1.0-1.el7.src.rpm
Wrote: /root/rpmbuild/RPMS/x86_64/hardware-1.0-1.el7.x86_64.rpm
Executing(%clean): /bin/sh -e /var/tmp/rpm-tmp.8fv6sC
+ umask 022
+ cd /root/rpmbuild/BUILD
+ cd hardware-1.0
+ rm -rf /root/rpmbuild/BUILDROOT/hardware-1.0-1.el7.x86_64exit 0
Removing intermediate container 634a9eac5487
 ---> 5e93a3eea5b6
Step 23/24 : RUN rpm  -ivh /root/rpmbuild/RPMS/x86_64/hardware-1.0-1.el7.x86_64.rpm
 ---> Running in f89fd856769d
Preparing...                          ########################################
Updating / installing...
hardware-1.0-1.el7                    ########################################
Removing intermediate container f89fd856769d
 ---> 80e9ac0f06ae
Step 24/24 : RUN /opt/hardware/bin/hardware
 ---> Running in 972d900b2ada
CPU information
numbers 4
|processor |cores |cpu MHz |cache size(MB) |stepping |vendor_id    |model name                                |
|0         |1     |2900    |8              |9        |GenuineIntel |Intel(R) Core(TM) i7-7820HQ CPU @ 2.90GHz |
|1         |1     |2900    |8              |9        |GenuineIntel |Intel(R) Core(TM) i7-7820HQ CPU @ 2.90GHz |
|2         |1     |2900    |8              |9        |GenuineIntel |Intel(R) Core(TM) i7-7820HQ CPU @ 2.90GHz |
|3         |1     |2900    |8              |9        |GenuineIntel |Intel(R) Core(TM) i7-7820HQ CPU @ 2.90GHz |


Memory information
       |total   |used   |free    |shared |cache   |buff   |available |used percent |
memory |1998 MB |312 MB |505 MB  |0 MB   |1024 MB |156 MB |1514 MB   |15.611546%   |
swap   |1023 MB |0 MB   |1023 MB |       |        |       |          |0.000000%    |
Removing intermediate container 972d900b2ada
 ---> 9e5006c0a475
Successfully built 9e5006c0a475
```
