# Building the application

Make sure `golang` is installed on your system.

For building you need `main.go`, `go.mod` and the `Makefile` in the same directory.
Simply run make. 

```sh
$ ls
go.mod main.go Makefile

$ make
go build -tags 'netgo osusergo static_build' -o webserver-volume-access

$ ls 
go.mod main.go Makefile webserver-volume-access
```

# Running the application
Make sure you have a file as input lying around somewhere.

Run the application

```sh
echo "hello world" > input
./webserver-volume-access -i input
```

In another shell modify your input file and refresh your browser. 
Curl to the webserver and see what it delivers to you.

```sh
echo "I'm sitting on the moon" > input
curl localhost:8080
```