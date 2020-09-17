# rpn

Get's a Reverse Polish Notation (RPN) string from an HTTP GET request and returns the computation.

See https://en.wikipedia.org/wiki/Reverse_Polish_notation for more RPN goodness.

# Testing
```console
myComputer:~$ go test
```

# Building & Running
```console
myComputer:~$ go build
myComputer:~$ ./rpn
```

# Usage

Typical usage is as follows:
```console
myComputer:~$ curl localhost:5000/calculate --data "14 2 / 1 + 3 3 + +"
14%
```

Errors (Note the f which is not a mathematical operator, and operations on too little data) are reported back:
```console
myComputer:~$ curl localhost:5000/calculate -d "14 f 2 / 1 + 3 3 + + + +"
Bad request see https://en.wikipedia.org/wiki/Reverse_Polish_notation strconv.Atoi: parsing "f": invalid syntax bad math operation only addition +, subtraction -, division /, and multiplication * supported
myComputer:~$
myComputer:~$  
myComputer:~$ curl localhost:5000/calculate --data "14 2 / 1 + 3 3 + + + +" 
Bad request see https://en.wikipedia.org/wiki/Reverse_Polish_notation strconv.Atoi: parsing "+": invalid syntax not enough values to operate on

```

  

