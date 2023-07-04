param(
    [string]$arg0,
    [string]$arg1,
    [string]$arg2
)

$scriptPath = $PSScriptRoot
$gatewayPath = Join-Path -Path $scriptPath -ChildPath "gateway/gateway.exe"

Write-Host "Starting gateway server at: $arg0"

Start-Process -FilePath $gatewayPath -ArgumentList $arg0 -NoNewWindow

Start-Sleep -Seconds 2

Write-Host "Starting gateway server at: $arg1"

Start-Process -FilePath $gatewayPath -ArgumentList $arg1 -NoNewWindow

Start-Sleep -Seconds 2

Write-Host "Starting gateway server at: $arg2"

Start-Process -FilePath $gatewayPath -ArgumentList $arg2 -NoNewWindow

Start-Sleep -Seconds 2

EXIT

#powershell -ExecutionPolicy Unrestricted -File serverstart.ps1 -arg0 "0.0.0.0:8888"-arg1 "0.0.0.0:8889" -arg2 "0.0.0.0:8890"              