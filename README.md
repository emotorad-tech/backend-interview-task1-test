# backend-interview-task1-test
### Installing Go
This code is written in Go, needing Go installed on your machine. Please install Go in one of the following ways:
 - With Homebrew: brew install go
 - By installing the binaries at: https://go.dev/doc/install

### Running the code
Clone the repository to your machine and go to the directory. 

Once in the directory, run the script ``go mod tidy`` to fetch the necessary dependencies to run the code.

Once the dependencies are installed, run the script by typing ``go run main.go`` into the terminal.

This will immediately starts sending a random speed value in JSON format specified in the document every second. The hard coded MQTT broker host is assumed as ``localhost:1883``, the standard default unencrypted MQTT port.
