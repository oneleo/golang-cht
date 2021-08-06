:: Run as administrator through CMD.exe
@echo off
:: Install Chocolatey
@"%SystemRoot%\System32\WindowsPowerShell\v1.0\powershell.exe" -NoProfile -InputFormat None -ExecutionPolicy Bypass -Command " [System.Net.ServicePointManager]::SecurityProtocol = 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))" && SET "PATH=%PATH%;%ALLUSERSPROFILE%\chocolatey\bin"
"%ProgramData%\chocolatey\bin\choco.exe" install -y chocolateygui golang vscode vscode-go mingw git.install --params "/NoShellIntegration"

:: Set Visual Studio Code Environment for Golang
::"%ProgramFiles%\Microsoft VS Code\bin\code" --force --install-extension "golang.go"
"%ProgramFiles%\Go\bin\go.exe" get "github.com/acroca/go-symbols"
"%ProgramFiles%\Go\bin\go.exe" get "github.com/davidrjenni/reftools/cmd/fillstruct"
"%ProgramFiles%\Go\bin\go.exe" get "github.com/haya14busa/goplay/cmd/goplay"
:: Install two versions of gocode (one as gocode-gomod)
"%ProgramFiles%\Go\bin\go.exe" get "github.com/stamblerre/gocode"
move /y "%USERPROFILE%\go\bin\gocode.exe" "%USERPROFILE%\go\bin\gocode-gomod.exe"
"%ProgramFiles%\Go\bin\go.exe" get "github.com/mdempsky/gocode"
"%ProgramFiles%\Go\bin\go.exe" get "github.com/sqs/goreturns"
"%ProgramFiles%\Go\bin\go.exe" get "github.com/uudashr/gopkgs/v2/cmd/gopkgs"
"%ProgramFiles%\Go\bin\go.exe" get "github.com/zmb3/gogetdoc"
"%ProgramFiles%\Go\bin\go.exe" get "honnef.co/go/tools/..."
"%ProgramFiles%\Go\bin\go.exe" get "golang.org/x/tools/cmd/gorename"
:: https://github.com/golang/go/issues/45732
"%ProgramFiles%\Go\bin\go.exe" get "golang.org/x/tools/gopls"
"%ProgramFiles%\Go\bin\go.exe" get "github.com/cweill/gotests/..."
"%ProgramFiles%\Go\bin\go.exe" get "github.com/rogpeppe/godef"
"%ProgramFiles%\Go\bin\go.exe" get "github.com/ramya-rao-a/go-outline"
"%ProgramFiles%\Go\bin\go.exe" get "github.com/go-delve/delve/cmd/dlv@master"
copy /v /y "%USERPROFILE%\go\bin\dlv.exe" "%USERPROFILE%\go\bin\dlv-dap.exe"
"%ProgramFiles%\Go\bin\go.exe" get "github.com/go-delve/delve/cmd/dlv"

:: Golang Present for CHT
"%ProgramFiles%\Git\cmd\git.exe" clone "https://github.com/oneleo/golang-cht.git" "%USERPROFILE%\go\src\github.com\oneleo\golang-cht"
"%ProgramFiles%\Git\cmd\git.exe" clone "https://github.com/golang/tools.git" "%USERPROFILE%\go\src\golang.org\x\tools"
"%ProgramFiles%\Go\bin\go.exe" get "golang.org/x/tools/cmd/present"
cd /d "%USERPROFILE%\go\src\github.com\oneleo\golang-cht"
start microsoft-edge:http://127.0.0.1:3999
"%USERPROFILE%\go\bin\present.exe" -base "%USERPROFILE%\go\src\golang.org\x\tools\cmd\present" -notes
pause