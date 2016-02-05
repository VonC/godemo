@echo off
SETLOCAL ENABLEEXTENSIONS ENABLEDELAYEDEXPANSION
set GOROOT=%PRGS%\go\latest
set GOPATH=%PROG%\git\godemo
set project=helloworldwebhttps

set PATH=C:\WINDOWS\system32;C:\WINDOWS;C:\WINDOWS\System32\Wbem;C:\WINDOWS\System32\WindowsPowerShell\v1.0\;C:\Windows\system32
set PATH=%PATH%;%GOROOT%\bin;%GOPATH%\bin;%PRGS%\git\latest\usr\bin

if not exist localhost.crt (
	cmd /v /c "set PATH=%~dp0;%PRGS%\git\latest\bin;%PRGS%\git\latest\usr\bin;%PATH% && git gen-cert"
)
@cd
@echo go install .
pause
go install .
set msg=doskey %project%=%GOPATH%\bin\%project%.exe $*
doskey /macros:all|grep %project% 1>NUL || echo %msg% |clip && echo %msg%
