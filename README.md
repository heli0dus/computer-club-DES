# computer-club-DES

## To launch the program locally 

Build your program and launch it with file as argument:

```
go build
./computer-club-DES <input file name>
```

## To launch inside docker container

Build docker image and run container matching local input to container space.

```
docker build -t computer-club .
docker run -v /local/path/to/input/file:/container/path/to/input/file computer-club /container/path/to/input/file 
```