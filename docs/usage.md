# Usage
There are two ways to run the api, locally or inside a container.

## Calc API
`calc-api` exposes methods over gRPC. 

### Local Go
```
go run ./services/calc-api
```

### Docker
Running this target builds the service inside a container. Once the containter is built, it is tagged with the short hash of the current branch and a latest tag.
```
make build-all -B
```
Run the container exposing a port to localhost.
```
docker run -p 3000:3000 superdecimal/calc-api:latest
```


## CLI
The CLI calls the `calc-api` gRPC methods.

```
go run ./services/cli
GMicro CLI Tool
>>>
```

### Supported Commands
```
>>> help

Commands:
  calc       CalcAPI operations
  clear      clear the screen
  exit       exit the program
  help       display help

>>> calc

CalcAPI operations

Commands:
  add      Adds two numbers
  sum      Sums numbers until eof (type eof to stop the stream)

```

#### Add
Add calls the `calc-api` to add two numbers. 

```
>>> calc add
Number a:
1
Number b:
2
Done, Result:  3  time:  5.000492ms
```

#### Sum
Sum calls the `calc-api` and starts a stream. The user can send multiple numbers to the api, once they are done, they just type `eof` to stop.

```
>>> calc sum
Number:
1
Number:
3
Number:
4
Number:
5
Number:
6
Number:
eof
Done, Result:  19
```