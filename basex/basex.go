package basex

// Base57 is a radix 57 encoding/decoding scheme
var Base57 = NewEncoding("23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

// Base64 is a radix 64 encoding/decoding scheme
var Base64 = NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
