# justatftpd
Just a tftp server that runs as a windows service.  
## Usage  
### Example:  
`justatftpd.exe --dir ./ --ro=false --conns :69`

### Flags:  
--dir   Path to directory you want to serve. Recommended to use unix style absolute paths. Default ./  
--ro    When true, the tftp server runs in read only mode. Default false  
--conns Connection string. Examples :69, 0.0.0.0:69, 192.168.1.1:69. Default: :69  

### Installing the service  
`justatftpd.exe --dir ./ --ro=true --conns :69 install`

### Uninstalling the service  
`justatftpd.exe --dir ./ --ro=true --conns :69 uninstall`  

### Logging  
justatftpd logs to windows event log when run as a service.  

## A note about PXE booting  
juatatftpd does not implement variable window sizing like the server bundled in Windows Deployment Services.
This means that you will not be able to transfer large WIM or ISO images using justatftpd.
You should use either pxelinux or iPXE to transfer large files over HTTP. Most modern NBPs support this feature.
