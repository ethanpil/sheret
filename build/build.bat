go build ../src/sheret.go
upx --ultra-brute sheret.exe
move sheret.exe ../bin/