param(
    [Parameter(Position = 0, ValueFromRemainingArguments = $true)]
    [string[]]$remainingArgs
)

$scriptPath = $PSScriptRoot
$gatewayPath = Join-Path -Path $scriptPath -ChildPath "gateway/gateway.exe"

foreach ($arg in $remainingArgs) {
    Write-Host "Starting gateway server at: $arg"
    Start-Process -FilePath $gatewayPath -ArgumentList $arg -NoNewWindow
    Start-Sleep -Seconds 2
}

EXIT
#powershell -ExecutionPolicy Unrestricted -File serverstart.ps1 -arg0 "8888"-arg1 "8889" -arg2 "8890"              