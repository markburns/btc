HD Wallet Genrration
----------

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


```json

{
  "hd": "generated with libbitcoin",

  "hd": {

    "seed": {
      "value": "this is the only non-deterministic step in these examples",
      "value": "bx seed > seed",
      "value": "e73fd9d3f95040b926c0d52cc220add0",

      "keywords": "bx mnemonic-encode < seed > words",
      "keywords": "whenever freely jealous caught mess twist attempt cigarette cell sneak during lady",
    },

    "private": "bx hd-new < seed > master-private",
    "private": "xprv9s21ZrQH143K46kpMYUr7t41j5EeZrzCaywXgMj3ZACUsuHQ1qb6C6PoWYHSEXHo7yLvaWiiWe32buNGp2VzK6SJPvtJv7FACjoumy9Mayk",

    "address": "bx hd-to-address > master-address",
    "address": "162EjaKLhesmsf9vJz4L2j5eQjLrHLrM2P",

    "wif": {
      "private": "bx hd-to-wif < master-private"
      "private": "L12b4umXNz2jrksaivxBA963tjoZky18sXq4WzoahDww4AFbs47n"
    },


    "ec": "bx hd-to-ecaddress < master-private",
    "ec": "71a0adb59fc17c29fb509d74c3d9824d290540da3b0fbfacf9c5789e4d3ee921",


    "keys": [
      {
      "private": "bx hd-private -i 0 < master-private > key-0/private",
      "private": "xprv9uJdLVz5WhzaAS2QhRaUNhnuZ6XR1KFUK6NV3k84SN2HiGFrqwRt7pCVcsoDp6R5htAnpZx4pmarwkAS5c8XNhEwh6tGRyb4qvjRbpxDGts",

      "public": "bx hd-public -i 0 < master-private > key-0/public",
      "public": "xpub68Hyk1WyM5YsNv6soT7Ujqje78MuQmyKgKJ5r8XfzhZGb4b1PUk8fcWyU93UWkXaA492iMp5q2XqpwjaWDS6ywFRD1z99o1ch18QVRTmteB",

      "address": "bx hd-to-address < key-0/public > key-0/address",
      "address": "1CCDBXV4vBCf5qxmcxLVtK1tGXzRiEaXpp",
    },

    {
      "private": "bx hd-private -i 1 < master-private > key-0/private",
      "private": "xprv9uJdLVz5WhzaFBMdh8hcjE6YQNwR1LHqxPXotF54K42xa9nx8oGTRwKk79GCR3bwpPWwk8dAgE1qEBqdXUJs8amzgNQ6SA1KhRHQv8fSc7L",

      "public": "bx hd-public -i 1 < master-private > key-0/public",
      "public": "xpub68Hyk1WyM5YsTfS6oAEd6N3GxQmuQo1hKcTQgdUfsPZwSx86gLahyjeDxTDh1PjsH9Ant1iGsVhk5VKgCX5qNBVkwk8ZjZYgPuySHRXS9Nv",

      "address": "bx hd-to-address < key-1/public > key-0/address",
      "address": "12nwNVK7iCAn4JnfwjsXjdEjLFHfZf9pjQ",
    },
    ]
  },
}

```
