package utility // 独自のユーティリティパッケージ
 
import (
    "crypto/md5"
    "crypto/sha1"
    "crypto/sha256"
    "crypto/sha512"
    "encoding/hex"
)
 
/* ****************************************
 ハッシュ化関数
 # md5,sha1,sha256,sha512に対応
 # デフォルトはsha256
**************************************** */
func HashStr(trg, alg string) string {
 
    // ハッシュ化文字列の初期値
    hashed := ""
 
    // ハッシュ対象文字列をバイト型スライスに変換
    b := []byte(trg)
 
    // バイト型スライスを指定されたアルゴリズムでハッシュ化し文字列へエンコード
    switch alg {
    case "md5":
        // md5
        md5 := md5.Sum(b)
        hashed = hex.EncodeToString(md5[:])
    case "sha1":
        // sha1
        sha1 := sha1.Sum(b)
        hashed = hex.EncodeToString(sha1[:])
    case "sha512":
        //sha512
        sha512 := sha512.Sum512(b)
        hashed = hex.EncodeToString(sha512[:])
    default:
        // 上記以外が指定された場合または未指定の場合はsha256
        sha256 := sha256.Sum256(b)
        hashed = hex.EncodeToString(sha256[:])
    }
 
    // ハッシュ化した文字列をリターンする
    return hashed
 
}
