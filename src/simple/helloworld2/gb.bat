@SETLOCAL ENABLEEXTENSIONS ENABLEDELAYEDEXPANSION
@set GOROOT=%PRGS%\go\latest
@set GOPATH=%HOME%\docker\godemo

@set PATH=C:\WINDOWS\system32;C:\WINDOWS;C:\WINDOWS\System32\Wbem;C:\WINDOWS\System32\WindowsPowerShell\v1.0\;C:\Windows\system32
@set PATH=%PATH%;%GOROOT%\bin;%GOPATH%\bin

@set go
go install .
@pause
@helloworld2 "in go FLASHY"
