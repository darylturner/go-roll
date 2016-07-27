# go-roll
CLI dice rolls written in Go. Dungeons and Dragons style expressions.

## Usage

Takes multiple dice rolls at time plus static modifiers. Modifiers and dice can be specified in arbitrary order. 

```sh
> ./roll 3d10+1d6+8    
[5 7 5] + [5] + 8 = 30
```

A more convulted roll.
```sh
> ./roll 10+2d20+3+8d6                   
10 + [19 18] + 3 + [6 4 3 3 4 1 6 5] = 82
```

Negative modifiers at the moment must be specified by adding a negative number.

```sh
> ./roll 3d10+1d6+8+-3     
[7 3 8] + [3] + 8 + -3 = 26
```

## Build
```sh
go build roll.go
```
