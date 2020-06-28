%define _prefix    /opt/hardware
Name:		hardware
Version:	1.0
Release:	1%{?dist}
Summary:	A hardware program

Group:		Applications/File
License:	GPLv3+
URL:	 	https://blog.packagecloud.io
Source0:	hardware-1.0.tar.gz

BuildRequires:	info
Requires:	info

%description
A helloworld program from the packagecloud.io blog!

%prep

%setup -q

%install
rm -rf $RPM_BUILD_ROOT
# Support mkdir and install two method
#mkdir -p $RPM_BUILD_ROOT/%{_prefix}
#mkdir -p $RPM_BUILD_ROOT/%{_prefix}/bin
#cp hardware $RPM_BUILD_ROOT/%{_prefix}/bin/
#mkdir -p $RPM_BUILD_ROOT/%{_prefix}/public
#cp public/* $RPM_BUILD_ROOT/%{_prefix}/public/

install -d  $RPM_BUILD_ROOT/%{_prefix}/bin
install hardware $RPM_BUILD_ROOT/%{_prefix}/bin/hardware
install -d $RPM_BUILD_ROOT/%{_prefix}/public
install public/* $RPM_BUILD_ROOT/%{_prefix}/public/*

%files
%{_bindir}/hardware
%{_prefix}/public

%clean
rm -rf %{buildroot}