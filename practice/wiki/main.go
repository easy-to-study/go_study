package main

import (
	"errors"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp" //正規表現のパッケージ
	"strings"
)

//wikiのデータ構造
type Page struct {
	Title string //タイトル
	Body  []byte //タイトルの中身
}

//パスのアドレスを設定して文字の長さを定数として持つ
const lenPath = len("/view/")

//テンプレートファイルの配列を作成
var templates = make(map[string]*template.Template)

//正規表現でURLを生成できる大文字小文字の英字と数字を判別する
//^は1文字目のチェック []その中の文字の種類 $最後の文字にマッチ +1文字以上
var titleValidator = regexp.MustCompile("^[a-zA-Z0-9]+$")

//.txt
const expend_string = ".txt"

//初期化関数
func init() {
	for _, tmpl := range []string{"edit", "view"} {
		//エラーの場合Panicを起こすためエラー処理はなし
		t := template.Must(template.ParseFiles(tmpl + ".html"))
		templates[tmpl] = t
	}
}

//タイトルのチェックを行う
func getTitle(w http.ResponseWriter, r *http.Request) (title string, err error) {
	title = r.URL.Path[lenPath:]
	if !titleValidator.MatchString(title) {
		http.NotFound(w, r)
		err = errors.New("Invalid Page Title")
		log.Print(err)
	}
	return
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		//editHandlerのURLに飛ばすことで編集ページに飛ばすことができます。
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func topHandler(w http.ResponseWriter, r *http.Request) {
	//main.goがいる階層のディレクトリにある.txtデータを取得する
	files, err := ioutil.ReadDir("./")
	if err != nil {
		err = errors.New("所定のディレクトリ内にテキストファイルがありません")
		log.Print(err)
		return
	}

	var paths []string    //テキストデータの名前
	var fileName []string //でキスとデータのファイル名
	for _, file := range files {
		//対象となる.txtデータのみを取得
		if strings.HasSuffix(file.Name(), expend_string) {
			//テキストデータの .txtで名前をスライスしたものをfileNameに入れる
			fileName = strings.Split(string(file.Name()), expend_string)
			paths = append(paths, fileName[0])
		}
	}
	//ファイルパスがなかった場合
	if paths == nil {
		err = errors.New("テキストファイルが存在しません")
		log.Print(err)
	}

	t := template.Must(template.ParseFiles("top.html"))
	err = t.Execute(w, paths)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Requestからページタイトルを取り出して、fnを呼び出す
		title := r.URL.Path[lenPath:]
		if !titleValidator.MatchString(title) {
			http.NotFound(w, r)
			err := errors.New("Invalid Page Title")
			log.Print(err)
			return
		}
		fn(w, r, title)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	//htmlの中にTitleやBodyを入れれるようにする
	err := templates[tmpl].Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//テキストファイルの保存メソッド
func (p *Page) save() error {
	//タイトルの名前でテキストファイルを作成して保存します。
	filename := p.Title + ".txt"
	//0600は、テキストデータを書き込んだり読み込んだりする権限を設定しています。
	return ioutil.WriteFile(filename, p.Body, 0600)
}

//titleからファイル名を読み込んで新しいPageのポインタを返す
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	//errに値が入ったらエラーとしてbodyの値をnilにして返す
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/top/", topHandler)
	http.ListenAndServe(":8080", nil)
}
