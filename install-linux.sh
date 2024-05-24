cp -r . ~/wyst-package-manager

cd ~/wyst-package-manager

go build -o bin/ src/*.go

sudo cp bin/wpm /usr/bin/