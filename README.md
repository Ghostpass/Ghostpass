# ghostpass

Secrets manager cryptosystem that enables users to distribute sensitive secrets in plainsight with plausible deniability

> WIP!

## Introduction

__Ghostpass__ is a secrets management cryptography scheme that can hide encrypted secrets in cleartext for the purpose of distribution across public mediums. It ensures that your secrets, whether in the form of authentication credentials or actual documentation, can appear in plainsight as normal cleartexts, and can be confidently transmitted even in a public medium with potential malicious actors.

Ghostpass is _novel_ in the sense that it applies modern symmetric cryptography towards digital consumer privacy through the use of textual steganography and plausible deniability. However, it should also be considered _novelty_ because it is __NOT__ a full replacement for current cryptographic software, but an ongoing effort to bridge together cryptography and privacy research. I encourage users to criticize, audit and expose shortcomings in order to better understand how these types of implementations can better be harnessed in the space of digital privacy.

## Features

> __Secure__ - all fields in credential stores are properly
> __Cryptographically Secure__ -
> __Privacy-First__ -

## Use Cases

### Who Can Use Ghostpass?

> __Journalists__ - hide secrets amongst corpuses of actual documents, and use plausible deniability to return bogus articles if ever interrogated
> __Lawyers__ - protect confidentiality of clients and sensitive anecdotes before trial by encrypting i

### Why Ghostpass?

There's quite a bit of password managers that are out there today, so why even bother with Ghostpass? In order to answer this question, let's take a look at different password managers "models" that already exist, and the problems that plague them:

* __Web-based password managers__ are centralized, and data breaches are possible since they are central points of failure.
* __Offline clients__ ( i.e KeePass) don't enable for fast distribution across hosts and portability. They share a unified filetype that require compliant client implementations to work

## Contributing

## License

[MIT License](https://codemuch.tech/license.txt)
