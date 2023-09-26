$exePath = $PSScriptRoot + "\pshs-core.exe"
$nextCommand = & $exePath
Write-Host "> " -NoNewline
Write-Host $nextCommand  -ForegroundColor Yellow
Invoke-Expression $nextCommand
