@echo off
SETLOCAL ENABLEEXTENSIONS ENABLEDELAYEDEXPANSION
set GOROOT=%PRGS%\go\latest
set GOPATH=%HOME%\docker\godemo

set PATH=C:\WINDOWS\system32;C:\WINDOWS;C:\WINDOWS\System32\Wbem;C:\WINDOWS\System32\WindowsPowerShell\v1.0\;C:\Windows\system32
set PATH=%PATH%;%GOROOT%\bin;%GOPATH%\bin

echo PATH=%PATH%
echo.
echo cmd /v /c "set GOBIN=&& set GOOS=linux&& set GOARCH=amd64&& go install unique/helloworldwebhttps"
echo.
cmd /v /c "set GOBIN=&& set GOOS=linux&& set GOARCH=amd64&& go install unique/helloworldwebhttps"
pause
dir %HOME%\docker\godemo\bin\linux_amd64
copy /Y %HOME%\docker\godemo\bin\linux_amd64\helloworldwebhttps .
