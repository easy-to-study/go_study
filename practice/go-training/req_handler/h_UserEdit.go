package req_handler // 独自のHTTPリクエストハンドラパッケージ
 
import (
    "database/sql"
    "fmt"
    "html/template"
    "net/http"
    "strconv"
 	_ "github.com/go-sql-driver/mysql"
 	"github.com/easy-to-study/go_study/practice/go-training/conf"
	"github.com/easy-to-study/go_study/practice/go-training/query"
)
 
// ユーザ情報編集画面
func HandlerUserEdit(w http.ResponseWriter, req *http.Request) {
 
    // テンプレートをパースする
    tpl := template.Must(template.ParseFiles("templates/user-form.gtpl"))
 
    // 設定ファイルを読み込む
    confDB, err := conf.ReadConfDB()
    if err != nil {
        fmt.Println("設定ファイルの読み込みに失敗しました。")
    }
 
    // 設定値から接続文字列を生成
    conStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", confDB.User, confDB.Pass, confDB.Host, confDB.Port, confDB.DbName, confDB.Charset)
 
    // データベース接続
    db, err := sql.Open("mysql", conStr)
    if err != nil {
        fmt.Println("データベース接続に失敗しました。")
    }
    // deferで処理終了前に必ず接続をクローズする
    defer db.Close()
 
    // リクエストパラメータからIDを取得してint64型に変換
    reqid, _ := strconv.Atoi(req.FormValue("id"))
    trgid := int64(reqid)
 
    // ユーザマスタの単一行取得関数を実行
    user, err := query.SelectUserById(trgid, db)
 
    // マップを展開してテンプレートを出力する
    if err := tpl.ExecuteTemplate(w, "user-form.gtpl", user); err != nil {
        fmt.Println(err)
    }
}
