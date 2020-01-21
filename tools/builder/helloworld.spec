%define _prefix    /opt/helloworld
Name:		helloworld
Version:	1.0
Release:	1%{?dist}
Summary:	A hello world program

Group:		Applications/File
License:	GPLv3+
URL:	 	https://blog.packagecloud.io
Source0:	helloworld-1.0.tar.gz

BuildRequires:	info
Requires:	info

%description
A helloworld program from the packagecloud.io blog!

%prep

%setup -q

%build
make PREFIX=%{_prefix} %{?_smp_mflags}


%install
make PREFIX=%{_prefix} install DESTDIR=%{?buildroot}

%preun


%files
%{_bindir}/helloworld

%clean
rm -rf %{buildroot}
rm -rf %{buildrootdir}