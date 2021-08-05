:: Run as administrator through CMD.exe
@echo off
:: Install Chocolatey
@"%SystemRoot%\System32\WindowsPowerShell\v1.0\powershell.exe" -NoProfile -InputFormat None -ExecutionPolicy Bypass -Command " [System.Net.ServicePointManager]::SecurityProtocol = 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))" && SET "PATH=%PATH%;%ALLUSERSPROFILE%\chocolatey\bin"
"%ProgramData%\chocolatey\bin\choco.exe" install -y chocolateygui golang vscode vscode-go mingw git.install --params "/NoShellIntegration"
"%ProgramData%\chocolatey\bin\choco.exe" install -y openjdk android-sdk android-ndk cmake make
pause