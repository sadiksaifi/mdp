# install.ps1 - mdp installer for Windows
# Usage: irm https://raw.githubusercontent.com/sadiksaifi/mdp/main/scripts/install.ps1 | iex

$ErrorActionPreference = "Stop"
Set-StrictMode -Version 2.0

$Repository = "sadiksaifi/mdp"
$InstallDirectory = Join-Path ([Environment]::GetFolderPath("LocalApplicationData")) "Programs\mdp"

function Write-Info([string]$Message) {
    Write-Host "==> $Message" -ForegroundColor Cyan
}

function Write-Success([string]$Message) {
    Write-Host "==> $Message" -ForegroundColor Green
}

function Get-MdpArchitecture {
    if ($env:PROCESSOR_ARCHITEW6432) {
        # Under x64 emulation on Windows ARM64, older .NET runtimes report X64.
        # PROCESSOR_ARCHITEW6432 identifies the native machine architecture.
        $architecture = $env:PROCESSOR_ARCHITEW6432
    }
    else {
        try {
            $architecture = [System.Runtime.InteropServices.RuntimeInformation]::OSArchitecture.ToString()
        }
        catch {
            $architecture = $env:PROCESSOR_ARCHITECTURE
        }
    }

    switch ($architecture.ToUpperInvariant()) {
        { $_ -in @("X64", "AMD64") } { return "amd64" }
        "ARM64" { return "arm64" }
        default { throw "Unsupported Windows architecture: $architecture. mdp supports AMD64 and ARM64." }
    }
}

function Add-MdpToUserPath([string]$Directory) {
    $userPath = [Environment]::GetEnvironmentVariable("Path", "User")
    $pathEntries = if ($userPath) { $userPath -split ";" } else { @() }
    $alreadyConfigured = $pathEntries | Where-Object {
        $_.TrimEnd("\") -ieq $Directory.TrimEnd("\")
    }

    if (-not $alreadyConfigured) {
        $newUserPath = if ($userPath) { "$userPath;$Directory" } else { $Directory }
        [Environment]::SetEnvironmentVariable("Path", $newUserPath, "User")
        Write-Success "Added $Directory to your user PATH"
    }

    if (($env:Path -split ";") -notcontains $Directory) {
        $env:Path = "$Directory;$env:Path"
    }
}

function Install-Mdp {
    if ([Environment]::OSVersion.Platform -ne [PlatformID]::Win32NT) {
        throw "This installer only supports Windows."
    }

    $architecture = Get-MdpArchitecture
    $assetName = "mdp-windows-$architecture.zip"

    Write-Info "Fetching the latest mdp release..."
    [Net.ServicePointManager]::SecurityProtocol = `
        [Net.ServicePointManager]::SecurityProtocol -bor [Net.SecurityProtocolType]::Tls12
    $headers = @{
        Accept = "application/vnd.github+json"
        "User-Agent" = "mdp-installer"
    }
    $release = Invoke-RestMethod `
        -Uri "https://api.github.com/repos/$Repository/releases/latest" `
        -Headers $headers

    $asset = $release.assets | Where-Object { $_.name -eq $assetName } | Select-Object -First 1
    if (-not $asset) {
        throw "Release asset $assetName was not found in release $($release.tag_name)."
    }

    $temporaryDirectory = Join-Path ([IO.Path]::GetTempPath()) ("mdp-install-" + [Guid]::NewGuid())
    $archivePath = Join-Path $temporaryDirectory $assetName

    try {
        New-Item -ItemType Directory -Path $temporaryDirectory -Force | Out-Null

        Write-Info "Downloading mdp $($release.tag_name) for Windows/$architecture..."
        Invoke-WebRequest -Uri $asset.browser_download_url -OutFile $archivePath -UseBasicParsing

        Write-Info "Extracting $assetName..."
        Expand-Archive -LiteralPath $archivePath -DestinationPath $temporaryDirectory -Force

        $binaryPath = Join-Path $temporaryDirectory "mdp.exe"
        if (-not (Test-Path -LiteralPath $binaryPath -PathType Leaf)) {
            throw "$assetName does not contain mdp.exe."
        }

        Write-Info "Installing to $InstallDirectory..."
        New-Item -ItemType Directory -Path $InstallDirectory -Force | Out-Null
        Copy-Item -LiteralPath $binaryPath -Destination (Join-Path $InstallDirectory "mdp.exe") -Force

        Add-MdpToUserPath $InstallDirectory

        $installedBinary = Join-Path $InstallDirectory "mdp.exe"
        $installedVersion = & $installedBinary --version
        if ($LASTEXITCODE -ne 0) {
            throw "Installation verification failed."
        }

        Write-Success "$installedVersion installed successfully!"
        Write-Host ""
        Write-Host "Get started with:"
        Write-Host "  mdp README.md"
        Write-Host ""
        Write-Host "Restart your terminal if mdp is not immediately available."
    }
    finally {
        if (Test-Path -LiteralPath $temporaryDirectory) {
            Remove-Item -LiteralPath $temporaryDirectory -Recurse -Force
        }
    }
}

Install-Mdp
