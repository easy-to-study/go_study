package config // 独自の設定ファイルパッケージ

import (
	"encoding/json"
	"io/ioutil"
)

// DB設定の構造体
type ConfDB struct {
	Host    string `json:"host"`    // ホスト名
	Port    int    `json:"port"`    // ポート番号
	DbName  string `json:"db-name"` // 接続先DB名
	Charset string `json:"charset"` // 文字コード
	User    string `json:"user"`    // 接続ユーザ名
	Pass    string `json:"pass"`    // 接続パスワード
}

// URL設定の構造体
func ReadConfDB() (*ConfDB, error) {

	// 設定ファイル名
	const conffile = "config/db.json"

	// 構造体を初期化
	conf := new(ConfDB)

	// 設定ファイルを読み込む
	cValue, err := ioutil.ReadFile(conffile)
	if err != nil {
		return conf, err
	}

	// 読み込んだjson文字列をデコードし構造体にセット
	err = json.Unmarshal([]byte(cValue), conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}
