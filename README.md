# passwordGen
A simple command line password generator written in go.
Inspired by a YouTube video by Melkey which is a tutorial for using Cobra and Cobra CLI

https://www.youtube.com/watch?v=yybzcix10XI

### To build the execuatable

go build .

### Or to install passwordGen in your $GOPATH/bin

go install

## Useage

### passwordGen generate [options]

## Examples

## passwordGen generate -h

Displays the help section

## passwordGen generate

The default option with no arguments is to generate an 8 char password of mixed upper and lower case chars.
It is highly reccomeneded to use a large password length -l and to use the -d and -s options to increase the complexity of your password.

##  passwordGen generate -l 16

The -l option specifies the length. In this case a length of 16 mixed upper and lower case chars.
To specify additional additional chars such as numbers and special chars see below.

## passwordGen generate -l 16 -d -s

 Generates a random password 16 chars long (-l 16) including digits (-d) and special chars (-s).
 In this case a password length (-l) of 16 chars (-d) for digits "0123456789" and -s for special chars
 "!@#$%^&*()_+{}[]|;:,.<>?-="

