#Seed

this is the only non-deterministic step in these examples

```
bx seed > seed
e73fd9d3f95040b926c0d52cc220add0

bx mnemonic-encode < seed > words
whenever freely jealous caught mess twist attempt cigarette cell sneak during lady
```


#Master private key

```
bx hd-new < seed > master-private
xprv9s21ZrQH143K46kpMYUr7t41j5EeZrzCaywXgMj3ZACUsuHQ1qb6C6PoWYHSEXHo7yLvaWiiWe32buNGp2VzK6SJPvtJv7FACjoumy9Mayk

bx hd-to-address > master-address
162EjaKLhesmsf9vJz4L2j5eQjLrHLrM2P

bx hd-to-wif < master-private
L12b4umXNz2jrksaivxBA963tjoZky18sXq4WzoahDww4AFbs47n

bx hd-to-ecaddress < master-private
71a0adb59fc17c29fb509d74c3d9824d290540da3b0fbfacf9c5789e4d3ee921

```


#Child keys

##Private

``` 
bx hd-private -i 0 < master-private > key-0/private
xprv9uJdLVz5WhzaAS2QhRaUNhnuZ6XR1KFUK6NV3k84SN2HiGFrqwRt7pCVcsoDp6R5htAnpZx4pmarwkAS5c8XNhEwh6tGRyb4qvjRbpxDGts
```

##Public
```
bx hd-public -i 0 < master-private > key-0/public
xpub68Hyk1WyM5YsNv6soT7Ujqje78MuQmyKgKJ5r8XfzhZGb4b1PUk8fcWyU93UWkXaA492iMp5q2XqpwjaWDS6ywFRD1z99o1ch18QVRTmteB
```

##Address
```
bx hd-to-address < key-0/public > key-0/address
1CCDBXV4vBCf5qxmcxLVtK1tGXzRiEaXpp

```

private key KybnQU4c2T65RzZbGn5Mq22HNH7GBoda2dhp1W6m1r94zVno2mpT

            payload     0x4707a525e75850b2a6018d28dbe3cc9ed688c5b52cfd714c1045ffed0286bb5401
                  dec:  8224696317156884363404180639537314541359879904682317186878780991563904960910337
            version 128
            checksum 2552843364

public key  19YhkH2ZQHjcDQqMkMabtEEqY7ti4sv8es
            payload  0x5dbfea0c8332f17b5247fd836d82f2b6ece5b247
                 dec: 535215972552226239844686566138855564960439054919
            version 0
            checksum 3,450,854,292



