param (
    [Parameter(Mandatory = $true)]
    [int]$PortNumber
)

$netTCPConnections = Get-NetTCPConnection -LocalPort $PortNumber -ErrorAction SilentlyContinue

if ($netTCPConnections) {
    $process = Get-Process -Id $netTCPConnections.OwningProcess | Select-Object -First 1

    if ($process) {
        Write-Host "Terminating process with ID $($process.Id) and name $($process.Name)"
        Stop-Process -Id $process.Id -Force
    }
    else {
        Write-Host "No process found running on port $PortNumber."
    }
}
else {
    Write-Host "No TCP connections found on port $PortNumber."
}
