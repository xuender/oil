/*
Package basex 提供对数字进行编码的工具.

默认支持2中格式 Base57 和 Base64.

并可根据需求自行定义编码格式.

*/
package basex

// Base57 is a radix 57 encoding/decoding scheme
var Base57 = NewEncoding("23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

// Base64 is a radix 64 encoding/decoding scheme
var Base64 = NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_")
