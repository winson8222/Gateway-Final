$scriptPath = $PSScriptRoot
$nginxPath = Join-Path -Path $scriptPath -ChildPath "nginx"

Set-Location -Path $nginxPath

Write-Host "Starting NGINX server from folder: $nginxPath"
& ./nginx -s quit

Write-Host "NGINX server stopped"
