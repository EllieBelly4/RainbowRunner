### Notes

#### Auth Messages

linAC [server] -> [client] ?

linAQ [client] -> [server] ?

Receive Data Raw: 0A4A19B0

Min Message length of 5 Length is Little Endian

Message structure:

| Len | Data|
|---|---|
|06 00| 01 01 01 01 01 01| 

#### Encryption
Login info before encryption:

```
0019B350  7F 00 00 01 00 00 00 00  00 00 00 00 68 65 6C 6C  ............hell
0019B360  6F 00 00 00 00 00 00 00  00 00 31 00 00 00 00 00  o.........1.....
0019B370  00 00 00 00 00 00 00 00  00 00 19 00 56 FD BA 11  ............Výº.
```

`0019B264` Keys?

https://paginas.fe.up.pt/~ei10109/ca/des.html

Encrypted data passed in with length 0x1E
Limited to 0x18 due to 8 byte block size

ESI input plain text 0019B35C