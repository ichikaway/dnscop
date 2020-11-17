# DNSCOP

DNS proxy(forwarder) daemon to block youtube video for kids.  

## Usage

```
go build dnscop.go  
sudo ./dnscop -listen="192.168.0.28:53"
```
### Options

```
sudo ./dnscop -listen="192.168.0.28:53" -resolver="1.1.1.1:53" -between="20:00-07:00"
```

Default value
- Listen IP:Port is "0.0.0.0:53"
- Resolver is "8.8.8.8:53"
- Between is "20:00-23:00"
