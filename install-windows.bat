@echo off
cd %~dp0
where go >nul 2>&1
if %errorlevel% neq 0 (
    echo please install go: https://go.dev/doc/install
    start https://go.dev/doc/install
    exit /b 1
) else (
    @echo off
    xcopy /s /e /i /h /y . "%USERPROFILE%\wpm"
    cd /d "%USERPROFILE%\wpm\src"
    go build -o ..\wpm.exe
    copy ..\wpm.exe C:\Windows\System32
)

pause
