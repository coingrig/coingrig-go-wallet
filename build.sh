#!/bin/sh
gomobile bind -ldflags "-s -w" -target android -o android/CGWallet.aar
gomobile bind -ldflags "-s -w" -target ios -o ios/CGWallet.xcframework