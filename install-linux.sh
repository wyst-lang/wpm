sudo cp -rf . ~/wpm
cd ~/wpm
go build -o bin/wpm src/*.go
sudo cp bin/wpm /usr/local/bin/