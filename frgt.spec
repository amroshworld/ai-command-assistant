Name:       frgt     
Version:    1.00    
Release:        1%{?dist}
Summary:    this ai tool help you generate command and save it in bash_history so can easily access it by arrow keys like normal terminal history simply by type frgt and your command     

License:    GPL-3.0-or-later      
URL:        https://github.com/amroshworld/ai-command-assistant    
Source0:    /home/amrosh/rpmbuild/SOURCES/frgt.tar.gz
     

%description
this ai tool help you generate command for any purpose  and save it in bash_history 
so can easily access it by arrow keys like normal terminal history simply by type 
frgt and your command description you want ai tool to generate for you 

%prep
%autosetup


%build
%configure
%make_build


%install
%make_install


%files
%license add-license-file-here
%doc add-docs-here



%changelog
* Thu Dec 14 2023 Amrosh
- 
