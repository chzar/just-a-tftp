# justatftpd
Just a tftp server that runs as a windows service.  

I put this server together to replace the default tftp server bundled with Windows Deployment Service. It is suitable for serving PXE boot files on windows or as a general purpose tftp server.  

Those looking for a linux solution should use hpa-tftpd.
  
Based on:  
github.com/kardianos/service  
github.com/pin/tftp  
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
justatftpd logs to the windows event log when run as a service.  

## A note about PXE booting  
juatatftpd does not implement variable window sizing like the server bundled in Windows Deployment Services.
This means that you will not be able to transfer large WIM or ISO images using justatftpd.
You should use either pxelinux or iPXE to transfer large files over HTTP. Most modern NBPs support this feature.
