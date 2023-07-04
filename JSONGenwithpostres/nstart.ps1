$scriptPath = $PSScriptRoot
$folderPath = Join-Path -Path $scriptPath -ChildPath "nginx"
$nginxPath = Join-Path -Path $scriptPath -ChildPath "nginx/nginx.exe"

Set-Location -Path $folderPath

Write-Host "Starting NGINX server from path: $nginxPath"

Start-Process -FilePath $nginxPath -NoNewWindow

Write-Host "NGINX server started!"

Start-Sleep -Seconds 2

# Exit the script
Exit