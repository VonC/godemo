set GOROOT=%PRGS%\go\latest
set GOPATH=%PROG%\git\godemo

set PATH=C:\WINDOWS\system32;C:\WINDOWS;C:\WINDOWS\System32\Wbem;C:\WINDOWS\System32\WindowsPowerShell\v1.0\;C:\Windows\system32
set PATH=%PATH%;%GOROOT%\bin;%GOPATH%\bin

pushd ..\..
godoc simple/godoc
godoc -http=:6060 || popd
popd
