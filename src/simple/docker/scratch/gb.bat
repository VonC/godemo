@echo off
SETLOCAL ENABLEEXTENSIONS ENABLEDELAYEDEXPANSION
set GOROOT=%PRGS%\go\latest
set GOPATH=%PROG%\git\godemo

set PATH=C:\WINDOWS\system32;C:\WINDOWS;C:\WINDOWS\System32\Wbem;C:\WINDOWS\System32\WindowsPowerShell\v1.0\;C:\Windows\system32
set PATH=%PATH%;%GOROOT%\bin;%GOPATH%\bin

@set go
@echo PATH=%PATH%
@echo.
@echo cmd /v /c "set GOBIN=&& set GOOS=linux&& set GOARCH=amd64&& go install simple/helloworld2"
@echo.
cmd /v /c "set GOBIN=&& set GOOS=linux&& set GOARCH=amd64&& go install simple/helloworld2"
pause
dir %PROG%\git\godemo\bin\linux_amd64
copy /Y %PROG%\git\godemo\bin\linux_amd64\helloworld2 .
