# Two ways to test hello world
# 1
C:\Users\kkot\golang-helloworld>go run main.go

# Above step will create executable and print the message and deleted executable automatically once the program is closed
# Because of this some people say that Go can be used as scripting language as compared to scripting language 
# Because end result has removed all the traces of an executable file

# 2 
# Second way is to build executable file and reuse, or share with my friend who can run it

C:\Users\kkot\golang-helloworld>go build main.go

# We have 2 files now, main.go and main.exe
\golang-helloworld>main.exe

# Mac or Ubuntu users
/golang-helloworld> ./main

# File is still present there and not deleted
# Rename file and run it and it should still print output

# Another way to generate executable
go build -o new-file-name


## Quick Summary

Run golang program

```bash
go run main.go
```

Testing

```bash
go test
```

Build binary

```bash
go build
```

Install binary

```bash
go install
```
