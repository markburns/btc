Base58 Check Prefixes
---------------

Hex          | (Dec)        | base58 | details
----------------------------------------------------
0x000        | (  0)        |      1 | bitcoin address
0x005        | (  5)        |      3 | pay to script hash address
0x06F        | (111)        |   m, n | test net address
0x080        | (128)        |  5,K,L | private key
0x142        | (322)        |  6P    | BIP38 Extended private key
0x0488B21E0  | (76,067,358) |  xpub  | BIP32 extended public key


Math
  Add(Wifable, Wifable)
  Sub(Wifable, Wifable)
  Mod(Wifable, Wifable)
  Div(Wifable, Wifable)
  Exp(Wifable, Wifable)


Wif
 Wifable
    big.Int

  NewWif(Wifable) Wif
    base58
      Encode(Wifable) Wif

Key
  Public
    Wif //Compressed
    Full       Wif
  Private
    Wif //Compressed

Address
  Wif
