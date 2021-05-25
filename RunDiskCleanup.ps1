#Requires -RunAsAdministrator

<#
    Simple script that is used for sending the diskcleanup utility to a list of computers.
    Usage:
        .\RunDiskCleanup.ps1 -ProfileID 555 V1370 V1484 V1533

        This will run PsExec, copying diskcleanup.exe to V1370, V1484, and V1533 and then executing
        it remotely. It will run the process on each computer one after the other, not waiting for it
        to finish on each one.
#>

param (
    [Parameter(Mandatory=$true, Position=0)]
    [ValidateRange(0, 9999)]
    [int] $ProfileID = 123,
    [Parameter(Mandatory=$true, ValueFromRemainingArguments=$true)]
    [string[]] $ComputerNames
)

$PSEXEC = "$PSScriptRoot\bin\PsExec.exe"
$DISKCLEANUP = "$PSScriptRoot\bin\diskcleanup.exe"
$VMS = $ComputerNames -join ","

<# PsExec Parameters:
   First parameter is the virtual machine names, comma separated and the first one starting with '\\'
   -c : Copy the file to the remote computer
   -f : Copy the file even if it exists on the computer anyway (overwriting it)
   -d : Don't wait for the process to terminate
   Second to last parameter is the diskcleanup executable
   Last parameter is the diskcleanup arguments
 diskcleanup Parameters:
   -profile : The profile ID for the disk cleanup tool #>
$ARG_STRING = "\\$VMS -f -d -c $DISKCLEANUP `"-profile=$ProfileID`""
Start-Process -FilePath $PSEXEC -ArgumentList $ARG_STRING -NoNewWindow
Write-Host "PsExec.exe $ARG_STRING"