# Disk-Cleanup
Simple utility for kicking off the Windows disk cleaning utility on remote computers when Powershell Remoting is not available. Utilizes PsExec to copy a small tool written in Go to the computers which sets the appropriate registry keys and then start cleanmgr.exe.
