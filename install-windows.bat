@echo off

where go >nul 2>&1
if %errorlevel% neq 0 (
    echo "please install go: https://go.dev/doc/install"
    exit /b 1
) else (
    xcopy . "%HOMEDRIVE%%HOMEPATH%\" /E /I /Y
    cd /d "%HOMEDRIVE%%HOMEPATH%\wyst-package-manager"
    go build -o bin/ src/*.go
)