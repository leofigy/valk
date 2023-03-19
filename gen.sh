# instructions are
```
fyne package -os darwin -icon icon.png
```

# quick test when using certs do i.e 
```
echo -n "test out the server" | openssl s_client -connect localhost:4001
```

# without self-sign certs do
```
echo -n "test out the server" | nc localhost 4001
```