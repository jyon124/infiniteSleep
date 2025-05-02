# Guidance to Terminate

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
