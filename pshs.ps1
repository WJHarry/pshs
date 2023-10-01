<#
.SYNOPSIS
Easy to search and apply history commands in Powershell.

.Description
https://github.com/WJHarry/pshs
Easy to search and apply history commands in Powershell.
The "history" command in PowerShell is session-level. pshs provides global-level history browsing and supports fuzzy search. Find your history and just "Enter" to execute it again.

#>

$exePath = $PSScriptRoot + "\pshs.exe"
$nextCommand = & $exePath

if ($nextCommand) {
   Write-Host "> " -NoNewline
   Write-Host $nextCommand  -ForegroundColor Yellow
   Invoke-Expression $nextCommand
}
