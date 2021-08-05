:: Run as administrator through CMD.exe
@echo off
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

:: Set GoMobile Environment for Android
setx ANDROID_NDK_HOME "%ANDROID_NDK_ROOT%"
"%ProgramFiles%\Go\bin\go.exe" get "golang.org/x/mobile/cmd/gomobile"
"%ProgramFiles%\Go\bin\go.exe" get "golang.org/x/mobile/cmd/gobind"
"%USERPROFILE%\go\bin\gomobile.exe" init
"%ProgramFiles%\Git\cmd\git.exe" clone "https://github.com/golang/mobile.git" "%USERPROFILE%\go\src\github.com\golang\mobile"

:: Set GoBot Environment for IoT Platforms
"%ProgramFiles%\Go\bin\go.exe" get -d -u "gobot.io/x/gobot"
"%ProgramFiles%\Git\cmd\git.exe" clone "https://github.com/hybridgroup/gobot.git" "%USERPROFILE%\go\src\github.com\hybridgroup\gobot"

:: Golang Present for CHT
"%ProgramFiles%\Git\cmd\git.exe" clone "https://github.com/oneleo/golang-cht.git" "%USERPROFILE%\go\src\github.com\oneleo\golang-cht"
"%ProgramFiles%\Go\bin\go.exe" get "golang.org/x/tools/cmd/present"
cd /d "%USERPROFILE%\go\src\github.com\oneleo\golang-cht"
start microsoft-edge:http://127.0.0.1:3999
"%USERPROFILE%\go\bin\present.exe" -notes
pause