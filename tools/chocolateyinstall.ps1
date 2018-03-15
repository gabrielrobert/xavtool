
$ErrorActionPreference = 'Stop';
$toolsDir   = "$(Split-Path -parent $MyInvocation.MyCommand.Definition)"
$distDir = (get-item $toolsDir ).parent.FullName

$fileLocation = Join-Path $distDir 'dist/windows_386/xavtool.exe'
$fileLocation64 = Join-Path $distDir 'dist/windows_amd64/xavtool.exe'

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










    








