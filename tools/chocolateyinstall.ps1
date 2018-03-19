
$ErrorActionPreference = 'Stop';
$toolsDir   = "$(Split-Path -parent $MyInvocation.MyCommand.Definition)"
$distDir = (get-item $toolsDir ).parent.FullName

$fileLocation = Join-Path $toolsDir 'x86/xavtool.exe'
$fileLocation64 = Join-Path $toolsDir 'x64/xavtool.exe'

Write-Host $fileLocation

$packageArgs = @{
  packageName   = $env:ChocolateyPackageName

  fileType      = 'exe'
  file           = $fileLocation
  file64      = $fileLocation64

  softwareName  = 'xavtool*'

  silentArgs    = "/qn /norestart /l*v `"$($env:TEMP)\$($packageName).$($env:chocolateyPackageVersion).MsiInstall.log`""
  validExitCodes= @(0, 3010, 1641)
}

Install-ChocolateyPackage @packageArgs










    








