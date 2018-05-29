#!/bin/bash

set -e

OS="darwin linux windows"
ARCH="amd64"

# clean up before building
rm -f bin/oneandone*

# get version from file
VERSION=$(cat -A version)

for GOOS in $OS; do
    for GOARCH in $ARCH; do
        export GOOS=$GOOS
        export GOARCH=$GOARCH
		PLATFORM="${GOOS}-${GOARCH}"
		if [[ $GOOS != windows ]]; then
			echo "Building binaries and scripts for $PLATFORM ..."
			go build -ldflags "-X main.AppVersion=${VERSION}" -o bin/oneandone
			tar --directory=bin --remove-files -cvf bin/oneandone-${PLATFORM}-${VERSION}.tar oneandone
			tar --directory=autocomplete -rvf bin/oneandone-${PLATFORM}-${VERSION}.tar bash_autocomplete
			gzip bin/oneandone-${PLATFORM}-${VERSION}.tar

			echo "Making installer script"
			(exec ./build/make-installer.sh "../bin/oneandone-${PLATFORM}-${VERSION}.tar.gz" "../bin/oneandone-${PLATFORM}-${VERSION}.sh")
		else
			# Restrict building windows binaries to windows OS only
			if [[ "$OSTYPE" == msys ]]; then
				echo "Building app binary and installer for $PLATFORM ..."
				go build -ldflags "-X main.AppVersion=${VERSION}" -o bin/oneandone.exe
				# Build windows installer. WiX tools (http://wixtoolset.org/) must be in $PATH
				# This works only on a Windows host.
				candle -nologo -dVersion=${VERSION} -arch x64 -out bin/oneandone.wixobj build/windows/Installer.wxs
				light.exe -nologo -spdb -ext WixUIExtension -out bin/oneandone-${PLATFORM}-${VERSION}.msi bin/oneandone.wixobj
			fi
		fi
    done
done

rm -f bin/*.exe bin/*.wixobj
