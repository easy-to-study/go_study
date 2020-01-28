package query // 独自のクエリパッケージ
 
import (
    "database/sql"
 
    _ "github.com/go-sql-driver/mysql"
)
 
// マスタからSELECTしたデータをマッピングする構造体
type M_USER struct {
    Id      string `db:"ID"`       // ID
    Account string `db:"ACCOUNT"`  // アカウント名
    Name    string `db:"NAME"`     // ユーザ名称
    Passwd  string `db:"PASSWORD"` // パスワード
    Created string `db:"CREATED"`  // 登録日
}

// 全行SELECT用の構造体配列
type UserList []M_USER
 
// データ登録関数
func InsertUser(acc, name, pw string, db *sql.DB) (id int64, err error) {
 
    // プリペアードステートメント
    stmt, err := db.Prepare("INSERT INTO M_USER(ACCOUNT,NAME,PASSWORD,CREATED) VALUES(?,?,?,now())")
    if err != nil {
        return 0, err
    }
    defer stmt.Close()
 
    // クエリ実行
    result, err := stmt.Exec(acc, name, pw)
    if err != nil {
        return 0, err
    }
 
    // オートインクリメントのIDを取得
    insertedId, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }
 
    // INSERTされたIDを返す
    return insertedId, nil
}
 
// 単一行データ取得関数
func SelectUserById(id int64, db *sql.DB) (userinfo M_USER, err error) {
 
    // 構造体M_USER型の変数userを宣言
    var user M_USER
 
    // プリペアードステートメント
    stmt, err := db.Prepare("SELECT ID,ACCOUNT,NAME,PASSWORD,CREATED FROM M_USER WHERE ID = ?")
    if err != nil {
        return user, err
    }
 
    // クエリ実行
    rows, err := stmt.Query(id)
    if err != nil {
        return user, err
    }
    defer rows.Close()
 
    // SELECTした結果を構造体にマップ
    rows.Next()
    err = rows.Scan(&user.Id, &user.Account, &user.Name, &user.Passwd, &user.Created)
    if err != nil {
        return user, err
    }
 
    // 取得データをマッピングしたM_USER構造体を返す
    return user, nil
}

// 全行データ取得関数
func SelectUserAll(db *sql.DB) (userlist UserList, err error) {
 
    // 配列宣言
    var ul UserList
 
    // プリペアードステートメント
    stmt, err := db.Prepare("SELECT ID,ACCOUNT,NAME,PASSWORD,CREATED FROM M_USER")
    if err != nil {
        return ul, err
    }
 
    // クエリ実行
    rows, err := stmt.Query()
    if err != nil {
        return ul, err
    }
    defer rows.Close()
 
    // SELECTした結果を構造体にマップ
    for rows.Next() {
        // 構造体宣言
        var user M_USER
        err = rows.Scan(&user.Id, &user.Account, &user.Name, &user.Passwd, &user.Created)
        // 配列にScan結果を追加
        ul = append(ul, user)
    }
 
    // 取得データをマッピングしたM_USER構造体配列を返す
    return ul, nil
}

// データ更新関数
func UpdateUser(acc, name, pw, id string, db *sql.DB) (err error) {
 
    // プリペアードステートメント
    stmt, err := db.Prepare("UPDATE M_USER SET ACCOUNT = ? ,NAME = ? , PASSWORD = ? WHERE ID = ?")
    if err != nil {
        return err
    }
    defer stmt.Close()
 
    // クエリ実行
    _, err = stmt.Exec(acc, name, pw, id)
    if err != nil {
        return err
    }
 
    return nil
}
