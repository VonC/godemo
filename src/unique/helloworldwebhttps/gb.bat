@SETLOCAL ENABLEEXTENSIONS ENABLEDELAYEDEXPANSION
@set GOROOT=%PRGS%\go\latest
@set GOPATH=%HOME%\docker\godemo
@set project=helloworldwebhttps

set PATH=C:\WINDOWS\system32;C:\WINDOWS;C:\WINDOWS\System32\Wbem;C:\WINDOWS\System32\WindowsPowerShell\v1.0\;C:\Windows\system32
set PATH=%PATH%;%GOROOT%\bin;%GOPATH%\bin;%PRGS%\git\latest\usr\bin

if not exist localhost.crt (
	cmd /v /c "set PATH=%~dp0;%PRGS%\git\latest\bin;%PRGS%\git\latest\usr\bin;%PATH% && git gen-cert"
)
@echo go install .
pause
go install .
copy /Y ..\..\..\bin\helloworldwebhttps.exe .
@rem set msg=doskey %project%=%GOPATH%\bin\%project%.exe $*
@rem doskey /macros:all|grep %project% 1>NUL || echo %msg% |clip && echo %msg%
