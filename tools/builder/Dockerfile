FROM centos:7 as builder-test

RUN yum -y install \
           gcc \
           tree \
           make \
           rpm-build

RUN mkdir -p /root/rpmbuild/{BUILD,BUILDROOT,RPMS,SOURCES,SPECS,SRPMS}

COPY helloworld-1.0.tar.gz /home/rpmbuild/SOURCES
COPY .rpmmacros /root/.rpmmacros

COPY helloworld-1.0.tar.gz /root/rpmbuild/SOURCES/
COPY helloworld.spec /root/rpmbuild/SPECS/helloworld.spec
RUN rpmbuild -ba /root/rpmbuild/SPECS/helloworld.spec 
RUN rpm  -ivh /root/rpmbuild/RPMS/x86_64/helloworld-1.0-1.el7.x86_64.rpm
RUN /opt/helloworld/bin/helloworld