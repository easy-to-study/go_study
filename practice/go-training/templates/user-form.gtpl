<!DOCTYPE html>
<html>
<head>
  <link rel="stylesheet" type="text/css" href="../asset/css/style.css">
</head>
<body>
  {{if .Id}}
  <h2>ユーザ情報の編集</h2>
  <form action="user-update" method="post">
  {{else}}
  <h2>ユーザ情報の登録</h2>
  <form action="user-confirm" method="post">  
  {{end}}
    <table>
      <tr>
        <td>アカウント名</td>
        <td><input type="text" name="account" value="{{.Account}}"></td>
      </tr>
      <tr>
        <td>お名前</td>
        <td><input type="text" name="name" value="{{.Name}}"></td>
      </tr>
      <tr>
        <td>パスワード</td>
        <td><input type="password" name="passwd" value="{{.Passwd}}"></td>
      </tr>
    </table>
    <input type="hidden" name="id" value="{{.Id}}">
    <input type="submit" class="btn-push btn-push-blue" value="確認画面へ">
  </form>
</body>
</html>
