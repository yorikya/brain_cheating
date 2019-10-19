# Cheat brain game
This is simple 2D game based [pixel library](https://github.com/faiface/pixel) which cheat the brain clock.

## Rules: 
You need click on the circle before time is up, in the left side you have a timeout clock for a click.

## Cheat the brain clock
First start play with delay (click on `delay` button) what happen behind is you have a small delay from when you click on circle to time when circle is disapeare. After play without delay just click `start` button, in this mode circle disapeare right after click, as result when you are running delay game then you run regular game (`start`) in regular game you should fill like the circle disapere before you click on the circle, this way you are cheat the brain sync.

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
go build .
./brain_cheating 

#Or with delay flag
./brain_cheating -delay
```


### Develope
```console
go run main.go 

#OR run with delay
go run main.go 
```

### Test

