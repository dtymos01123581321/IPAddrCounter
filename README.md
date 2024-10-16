# Counter of IP addresses

You have a simple text file with IPv4 addresses. One line is one address, line by line:

145.67.23.4
8.34.5.23
89.54.3.124
89.54.3.124
3.45.71.5
...
The file is unlimited in size and can occupy tens and hundreds of gigabytes.

You should calculate the number of unique addresses in this file using as little memory and time as possible. 
There is a "naive" algorithm for solving this problem (read line by line, put lines into HashSet). 
It's better if your implementation is more complicated and faster than this naive algorithm.

```
GOROOT=/Users/dimasik/sdk/go1.23.2 #gosetup
GOPATH= #gosetup
/Users/dimasik/sdk/go1.23.2/bin/go build -o /Users/dimasik/Library/Caches/JetBrains/IntelliJIdea2023.1/tmp/GoLand/___1go_build_test1_go /Users/dimasik/work/studione/inw7000/code/GoLand/IPAddrCounter/test1.go #gosetup
/Users/dimasik/Library/Caches/JetBrains/IntelliJIdea2023.1/tmp/GoLand/___1go_build_test1_go
Number of unique IPs: 1000000000
Time taken: 13m48.096369333s
```
