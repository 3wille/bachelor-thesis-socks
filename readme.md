# IPv6 Address hopping - SOCKS5 proxy

This repository is part of my thesis "design and implementation of context-aware IPv6 address hopping".
The proxy behaves similar to the one delivered with Tor and The Tor Browser.

## Initialize Subrepository

The larger part of this is implemented in a fork of github.com/oov/socks5.
It is tweaked to support above functionality. To Initialize the subrepo run ``git submodule update --init`` or clone the whole repo with ``git clone --recursive <url>``.

## Build

To trigger a build of the go package run ``make``.
This yields an executable ``main``.
