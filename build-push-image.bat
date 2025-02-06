@echo off
set IMAGE_NAME=dms-be
set TAG=latest
set GHCR_USERNAME=ahmadzeinalwafi
set REPO=ghcr.io/%GHCR_USERNAME%/%IMAGE_NAME%

echo Building Docker image...
docker build -t %REPO%:%TAG% .

if %ERRORLEVEL% NEQ 0 (
    echo Build failed. Exiting...
    exit /b %ERRORLEVEL%
)

echo Pushing Docker image...
docker push %REPO%:%TAG%

if %ERRORLEVEL% NEQ 0 (
    echo Push failed. Exiting...
    exit /b %ERRORLEVEL%
)

echo Done!
