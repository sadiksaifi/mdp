$ErrorActionPreference = "Stop"
Set-StrictMode -Version 2.0

$installerPath = Join-Path $PSScriptRoot "install.ps1"
$tokens = $null
$parseErrors = $null
$ast = [System.Management.Automation.Language.Parser]::ParseFile(
    $installerPath,
    [ref]$tokens,
    [ref]$parseErrors
)
if ($parseErrors.Count -gt 0) {
    throw "Could not parse ${installerPath}: $($parseErrors -join '; ')"
}

$architectureFunction = $ast.Find({
    param($node)
    $node -is [System.Management.Automation.Language.FunctionDefinitionAst] -and
        $node.Name -eq "Get-MdpArchitecture"
}, $true)
if (-not $architectureFunction) {
    throw "Get-MdpArchitecture was not found in $installerPath"
}
Invoke-Expression $architectureFunction.Extent.Text

$originalArchitecture = $env:PROCESSOR_ARCHITECTURE
$originalArchitectureW6432 = $env:PROCESSOR_ARCHITEW6432
try {
    # A PowerShell x64 process on Windows ARM64 reports AMD64 as its process
    # architecture and ARM64 as the native machine architecture.
    $env:PROCESSOR_ARCHITECTURE = "AMD64"
    $env:PROCESSOR_ARCHITEW6432 = "ARM64"

    $actual = Get-MdpArchitecture
    if ($actual -ne "arm64") {
        throw "Get-MdpArchitecture returned '$actual' under x64 emulation on ARM64; expected 'arm64'."
    }
}
finally {
    $env:PROCESSOR_ARCHITECTURE = $originalArchitecture
    $env:PROCESSOR_ARCHITEW6432 = $originalArchitectureW6432
}

Write-Host "PowerShell installer architecture tests passed."
