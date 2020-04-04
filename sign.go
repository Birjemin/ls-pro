package main

import (
    "crypto/md5"
    "encoding/hex"
)

func Md5Encode(s string) string {
    md5Ctx := md5.New()
    md5Ctx.Write([]byte(s))
    return hex.EncodeToString(md5Ctx.Sum(nil))
}
