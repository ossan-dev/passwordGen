# passwordGen
A simple command line password generator written in go.

go build .
./passwordGen.exe generate [options]

example

./passwordGen.exe generate -h

Show help

./passwordGen.exe generate -l 16 -d -s
 Generates a random password 16 chars long including digits and special chars.
