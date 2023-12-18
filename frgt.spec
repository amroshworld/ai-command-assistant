Name:           frgt
Version:        0.0.1
Release:        1%{?dist}
Summary:        ai commandline tool that help to write any command for any task 

License:      GPL     
Source0:      %{name}-%{version}.tar.gz   

Requires:     bash 

%description
ai commandLine tool that help to write any command for any task 
using frgt and task you want then the out put will save in your history 
and you can get it by using arrow keys like normal command history 

%global debug_package %{nil}

%prep
%autosetup 

%build
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
go build -ldflags="-B  0x${RPM_BUILD_ID}" -gcflags="-N -l" -o %{name}

%install
rm -rf $RPM_BUILD_ROOT
mkdir -p $RPM_BUILD_ROOT/%{_bindir}
cp %{name} $RPM_BUILD_ROOT/%{_bindir}

%undefine _missing_build_ids_terminate_build


%files
%{_bindir}/%{name}

%changelog
* Sun Dec  12 2023 Amr Taha <amrosh.world@gmail.com> - 0.0.1
- First version being packaged


