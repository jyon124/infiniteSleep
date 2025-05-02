# Guidance to Terminate
Go Running Server
```
ps aux | grep go
```

Look for something similar to below
```
username            54663   0.0  0.0 35364512   4956 s001  S    10:20PM   0:00.05 /var/folders/gj/1g_3sgtd4b/T/go-build2196433466/b001/exe/m
```

Then Kill
```
kill -9 54663
```

Docker Running server
```
ps aux | grep docker
```

```
username            64626   0.0  0.1 35415704  16184 s001  S+   10:54PM   0:00.06 docker run -p 8080:8080 infinite-sleep-server
```

```
kill -9 64626
```

# Guidance to build Dockerfile & Run
```
docker build -t infinite-sleep-server .
```

```
docker run -p 8080:8080 infinite-sleep-server
```
