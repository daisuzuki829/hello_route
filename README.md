## gin を用いたルーティング処理

| 確認項目 | 期待値 | 結果 |  
| ---- | ---- | ---- |
| http://localhost:8080 にアクセス | http://localhost:8080/hello にリダイレクト <br> 「Hello Michael!」と表示 | ◎
| 適当なURL http://localhost:8080/abcd　など適当なURLにアクセス | http://localhost:8080/hello にリダイレクト <br> 「Hello!」と表示 | ◎
| http://localhost:8080/hello?name={name} にアクセス | 「Hello {name}!」と表示 | ◎ |
| http://localhost:8080/hello/{name}　にアクセス | 「Hello {name}!」と表示 | ◎ |
| http://localhost:8080/api/hello にアクセス  | 「{"name":"Michael"}」が返る | ◎ |
| http://localhost:8080/api/hello?name={name}　にアクセス  | 「{"name":"{name}"}」が返る | ◎ |
| http://localhost:8080/api/hello/{name}　にアクセス  | 「{"name":"{name}"}」が返る | ◎ |

