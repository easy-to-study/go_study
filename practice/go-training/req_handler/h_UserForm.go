package req_handler // 独自のHTTPリクエストハンドラパッケージ
 
import (
    "fmt"
    "html/template"
    "net/http"
)
 
/// 入力フォーム画面
func HandlerUserForm(w http.ResponseWriter, r *http.Request) {
    // テンプレートをパースする
    tpl := template.Must(template.ParseFiles("templates/user-form.gtpl"))
 
    // テンプレートに出力する値をマップにセット
    values := map[string]string{}
 
    // マップを展開してテンプレートを出力する
    if err := tpl.ExecuteTemplate(w, "user-form.gtpl", values); err != nil {
        fmt.Println(err)
    }
}
