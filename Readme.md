# My HNG 9 Task 2

The task was completed using golang, it has been containerized for ease of use.


# Without Docker
To run, first build bin using the command below
```bash
$ go build .
```
Then run the executable
```bash
$ ./hng-task-1
```
# With Docker
To run, first build the image using the command below
```bash
$ docker build -t hng-task-1 .
```
Then run the image
```bash
$ docker run -p 80:80 hng-task-1
```

# With Docker compose
```bash
$ docker-compose up --build -d
```

Then visit http://localhost to see the output