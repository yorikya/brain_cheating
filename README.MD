# Cheat brain game
This is simple 2D game which cheat the brain clock.

## Rules: 
You need click on circle before time is up, in left side you have a timeout clock for a click.

## Cheat the brain clock
First start play with `-delay` flag then after try to play without `-delay` flag, when you are running without delay you should fill like the circle disapere before you click.

### Requirements 
https://github.com/faiface/pixel#requirements


### Installt
```console
go get github.com/yorikya/brain_cheating
go build .
./brain_cheating #Or with -delay flag
```

### Run
```console
go mod tidy
go run main.go 

#OR run with delay
go run main.go 
```


