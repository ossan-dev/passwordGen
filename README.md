# passwordGen
A simple command line password generator written in go.

## go build .

## ./passwordGen.exe generate [options]

# example

## ./passwordGen.exe generate -h

Shows help

## ./passwordGen.exe generate

The default option is to generate an 8 char password of mixed upper and lower case.

## ./passwordGen.exe generate -l 16 -d -s

 Generates a random password 16 chars long (-l 16) including digits (-d) and special chars (-s).
