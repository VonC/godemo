@SETLOCAL ENABLEEXTENSIONS ENABLEDELAYEDEXPANSION
@set GOROOT=%PRGS%\go\latest
@set GOPATH=%HOME%\docker\godemo

@set PATH=C:\WINDOWS\system32;C:\WINDOWS;C:\WINDOWS\System32\Wbem;C:\WINDOWS\System32\WindowsPowerShell\v1.0\;C:\Windows\system32
@set PATH=%PATH%;%GOROOT%\bin;%GOPATH%\bin

go install
dir ..\..\..\..\pkg\windows_amd64\simple\helloworld
pause
godoc -v simple/helloworld/hw
go test -v
godoc -http=:6060
