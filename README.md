# Disk-Cleanup
Simple utility for kicking off the Windows disk cleaning utility on remote computers when Powershell Remoting is not available. Utilizes PsExec to copy a small tool written in Go to the computers which sets the appropriate registry keys and then start cleanmgr.exe.

https://docs.microsoft.com/en-us/sysinternals/downloads/psexec
https://docs.microsoft.com/en-us/troubleshoot/windows-server/backup-and-storage/automating-disk-cleanup-tool
