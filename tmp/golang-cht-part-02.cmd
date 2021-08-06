:: Run as administrator through CMD.exe
@echo off
:: Install Chocolatey
@"%SystemRoot%\System32\WindowsPowerShell\v1.0\powershell.exe" -NoProfile -InputFormat None -ExecutionPolicy Bypass -Command " [System.Net.ServicePointManager]::SecurityProtocol = 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))" && SET "PATH=%PATH%;%ALLUSERSPROFILE%\chocolatey\bin"
"%ProgramData%\chocolatey\bin\choco.exe" install -y golang openjdk android-sdk android-ndk openal cmake make git.install --params "/NoShellIntegration"

:: Set GoMobile Environment for Android
setx ANDROID_NDK_HOME "%%ANDROID_NDK_ROOT%%"
"%ProgramFiles%\Git\cmd\git.exe" clone "https://github.com/golang/mobile.git" "%USERPROFILE%\go\src\golang.org\x\mobile"
"%ProgramFiles%\Go\bin\go.exe" get "golang.org/x/mobile/cmd/gomobile"
"%ProgramFiles%\Go\bin\go.exe" get "golang.org/x/mobile/cmd/gobind"

::cd "%USERPROFILE%\go\src\golang.org\x\mobile\example\basic"
::"%USERPROFILE%\go\bin\gomobile.exe" init -openal "%ProgramFiles(x86)%\OpenAL"
::"%USERPROFILE%\go\bin\gomobile.exe" install "%USERPROFILE%\go\src\golang.org\x\mobile\example\basic"

:: Set GoBot Environment for IoT Platforms
"%ProgramFiles%\Git\cmd\git.exe" clone "https://github.com/hybridgroup/gobot.git" "%USERPROFILE%\go\src\gobot.io\x\gobot"
"%ProgramFiles%\Go\bin\go.exe" get -d -u "gobot.io/x/gobot"
pause