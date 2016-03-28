@echo off
SETLOCAL ENABLEEXTENSIONS ENABLEDELAYEDEXPANSION
set GOROOT=%PRGS%\go\latest
set GOPATH=%HOME%\docker\godemo

set PATH=C:\WINDOWS\system32;C:\WINDOWS;C:\WINDOWS\System32\Wbem;C:\WINDOWS\System32\WindowsPowerShell\v1.0\;C:\Windows\system32
set PATH=%PATH%;%GOROOT%\bin;%GOPATH%\bin

echo go install .
pause
go install .
copy /Y ..\..\..\bin\concurrency.exe .
echo concurrency.exe
pause
concurrency.exe
pause
rem start "" http://divan.github.io/posts/go_concurrency_visualize/#ping-pong:cacee17578a8ad02e42a5a54daf9c476
start "" http://divan.github.io/demos/gifs/pingpong.gif
