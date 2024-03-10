# Golang-Web-API

🎞️画像検索WebAPI

## 🏠URL

https://pixfinder.aki158-website.blog/

## ✨デモ

### Webアプリケーション

#### サインアップ(ユーザー登録)

![signup](https://github.com/Recursion-teamB-create-webAPI/Golang-Web-API/assets/119317071/641a986f-4dc0-4bb2-80b1-203979b02fab)

#### サインイン(ログイン)

![signin](https://github.com/Recursion-teamB-create-webAPI/Golang-Web-API/assets/119317071/7e0303f9-7829-455d-8ddf-13bb9a06ab62)

#### 画像検索

![pixfinder](https://github.com/Recursion-teamB-create-webAPI/Golang-Web-API/assets/119317071/32caf6de-4dfa-4bda-b94d-05d560d4f8b7)

### WebAPI

![api](https://github.com/Recursion-teamB-create-webAPI/Golang-Web-API/assets/119317071/172a843a-e3bb-4195-a904-90678f427828)

## 📝説明

このAPIは、画像検索WebAPIです。

特定のキーワード(例. coffee)の画像URLをまとめて自動で取得したいと思ったことはありませんか?

このAPIには、画像検索に特化したユーザーのニーズを満たす4つの機能が備わっています。

APIの開発時には、以下のような利用シナリオとメリットを想定して作成しました。

### 利用シナリオ

| 項目 | 内容 |
| ------- | ------- |
| WebサイトやWebアプリケーションでのコンテンツ生成 | ユーザーが特定のトピックに関する情報を求めているとき、このAPIを利用して関連画像を表示し、ユーザー体験を向上させることができます。 |
| デザインとインスピレーション | デザイナーやアーティストがプロジェクトのアイデアを探求する際に、様々な画像を簡単に検索し、インスピレーションを得るために使用できます。 |
| 商品リサーチ | Webサイトで商品を販売する際、類似の商品や競合他社の商品画像を調査するために使用できます。 |

### メリット

| 項目 | 内容 |
| ------- | ------- |
| 効率性 | 手動で画像を検索した場合は、選択する手間が発生しますが、このAPIを利用することで効率的に目的の画像を取得できます。 |
| 時間の節約 | 自動化された検索により時間を節約し、他の作業により多くの時間を割くことができます。 |
| 幅広い用途 | ブログ記事、学術研究、プレゼンテーション、マーケティング資料など、さまざまな目的で利用できます。 |

基本的な機能として、 検索機能/詳細表示/リスト表示/データの集計ができます。

機能の詳細については、[機能一覧](#機能一覧)を確認してください。

また、このAPIのデモWebアプリケーションも開発しました。

APIを利用するのに迷われている方は、このアプリケーションにアクセスしてみてください。

[URL](#URL)にアクセスしサインアップすることで、各機能について簡単に体験することができます。

### 注意事項

APIの検索機能には、`Google Custom Search API`を利用し検索結果を取得しています。

このAPIの利用者は、ライトユーザー(利用する頻度が少ないユーザー)を想定して作成しています。

データベースを利用することで、`Google Custom Search API`の検索上限(1日あたりの検索数:最大100件)以上の検索ができるような工夫をしていますが、過度な利用はオススメしていません。

## 🚀使用方法と使用例

イメージは[デモ](#デモ)を参考にしてください。

また、使用する前に[注意事項](#注意事項)を確認してください。

### Webアプリケーション

#### サインアップ(ユーザー登録)

1. [URL](#URL)にアクセスする
2. Sign Out ボタンをクリックする
3. アカウントを作る ボタンをクリックする
4. Username と Password を入力する
5. Sign Up ボタンをクリックする
6. 画像検索を利用する

#### サインイン(ログイン)

1. [URL](#URL)にアクセスする
2. Sign Out ボタンをクリックする
3. Username と Password を入力する
4. Sign In ボタンをクリックする
5. 画像検索を利用する

#### 画像検索

1. サインアップ または サインインで 画像検索ページに移動する
2. 検索欄に検索ワードを入力する<br>(検索候補の表示 ボタンをクリックして検索候補から選択することもできます)
3. Search ボタンをクリックする
4. 検索結果一覧に表示された検索結果を確認する<br>`Previout` : 前の画像へ<br>`Next` : 次の画像へ<br>`Go to description` : 画像の詳細表示ページへ移動する
5. アプリケーションを終了したい場合は、Sign Out ボタンをクリックする

### WebAPI

#### 検索機能

```
curl https://pixfinder.aki158-website.blog/api/search?keyword={検索ワード}
```

##### 使用例

```
curl https://pixfinder.aki158-website.blog/api/search?keyword=coffee
```

レスポンス

```
{
    "imageData":
    {
        "images":
        [
            "https://neurosciencenews.com/files/2023/06/coffee-brain-caffeine-neuroscincces.jpg",
            "https://upload.wikimedia.org/wikipedia/commons/e/e4/Latte_and_dark_coffee.jpg",
            "https://www.eatright.org/-/media/images/eatright-articles/eatright-article-feature-images/benefitsofcoffee_600x450.jpg?as=0\u0026w=967\u0026rev=6c8a9cd4a94d4cac8af8543054fd7b07\u0026hash=4C95EA3A031A783C6DFA3ED04AB4FED4",
            "https://www.rush.edu/sites/default/files/media-images/Coffee_OpenGraph.png",
            "https://i0.wp.com/images-prod.healthline.com/hlcmsresource/images/AN_images/butter-coffee-1296x728-feature.jpg?w=1155\u0026h=1528",
            "https://www.bhg.com/thmb/Dw9Sxcivh_bczUpo91Egapy-lGc=/7952x0/filters:no_upscale():strip_icc()/feshly-brewed--latte-coffee-on-a-white-table-1434303312-4d24a51e3bc748d68553a836499d0100.jpg",
            "https://upload.wikimedia.org/wikipedia/commons/thumb/e/e4/Latte_and_dark_coffee.jpg/640px-Latte_and_dark_coffee.jpg",
            "https://dynaimage.cdn.cnn.com/cnn/c_fill,g_auto,w_1200,h_675,ar_16:9/https%3A%2F%2Fcdn.cnn.com%2Fcnnnext%2Fdam%2Fassets%2F150929101049-black-coffee-stock.jpg",
            "https://static.scientificamerican.com/sciam/cache/file/4A9B64B5-4625-4635-848AF1CD534EBC1A_source.jpg?w=1200",
            "https://upload.wikimedia.org/wikipedia/commons/thumb/c/c5/Roasted_coffee_beans.jpg/640px-Roasted_coffee_beans.jpg"
        ]
    },
    "status":"success",
    "cause":""
}
```

#### 詳細表示

```
curl https://pixfinder.aki158-website.blog/api/description?keyword={検索ワード}
```

##### 使用例

```
curl https://pixfinder.aki158-website.blog/api/description?keyword=coffee
```

レスポンス

```
{
    "description":
    {
        "id":85,
        "item":"coffee",
        "imageData":
        {
            "images":
            [
                "https://neurosciencenews.com/files/2023/06/coffee-brain-caffeine-neuroscincces.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/e/e4/Latte_and_dark_coffee.jpg",
                "https://www.eatright.org/-/media/images/eatright-articles/eatright-article-feature-images/benefitsofcoffee_600x450.jpg?as=0\u0026w=967\u0026rev=6c8a9cd4a94d4cac8af8543054fd7b07\u0026hash=4C95EA3A031A783C6DFA3ED04AB4FED4",
                "https://www.rush.edu/sites/default/files/media-images/Coffee_OpenGraph.png",
                "https://i0.wp.com/images-prod.healthline.com/hlcmsresource/images/AN_images/butter-coffee-1296x728-feature.jpg?w=1155\u0026h=1528",
                "https://www.bhg.com/thmb/Dw9Sxcivh_bczUpo91Egapy-lGc=/7952x0/filters:no_upscale():strip_icc()/feshly-brewed--latte-coffee-on-a-white-table-1434303312-4d24a51e3bc748d68553a836499d0100.jpg",
                "https://upload.wikimedia.org/wikipedia/commons/thumb/e/e4/Latte_and_dark_coffee.jpg/640px-Latte_and_dark_coffee.jpg",
                "https://dynaimage.cdn.cnn.com/cnn/c_fill,g_auto,w_1200,h_675,ar_16:9/https%3A%2F%2Fcdn.cnn.com%2Fcnnnext%2Fdam%2Fassets%2F150929101049-black-coffee-stock.jpg",
                "https://static.scientificamerican.com/sciam/cache/file/4A9B64B5-4625-4635-848AF1CD534EBC1A_source.jpg?w=1200",
                "https://upload.wikimedia.org/wikipedia/commons/thumb/c/c5/Roasted_coffee_beans.jpg/640px-Roasted_coffee_beans.jpg"
            ]
        },
        "search_count":0,
        "created_at":"2024-03-06 02:20:13",
        "updated_at":"2024-03-06 02:20:13"
    },
    "status":"success",
    "cause":""
}
```

#### リスト表示

```
curl https://pixfinder.aki158-website.blog/api/list?keyword={検索ワード(部分一致)}
```

##### 使用例

```
curl https://pixfinder.aki158-website.blog/api/list?keyword=c
```

レスポンス

```
{
    "list":
    [
        "cat",
        "coffee"
    ],
    "status":"success",
    "cause":""
}
```

##### その他の使用例

クエリパラメータを設定しない(データベースに登録されているアイテムをリストとして取得する)

```
curl https://pixfinder.aki158-website.blog/api/list
```

#### データの集計

```
curl https://pixfinder.aki158-website.blog/api/total_result?keyword={検索ワード}&page={ページ番号}&perpage={ページ数}&order={count または newest}
```

##### 使用例

データベースに登録されているアイテム(cを文字列に含む)から検索数の多い順で上位5件を取得する

```
curl https://pixfinder.aki158-website.blog/api/total_result?keyword=c&page=1&perpage=5&order=count
```

レスポンス

```
{
    "totalResult":
    [
        {
            "item":"cat",
            "search_count":5,
            "updated_at":"2024-03-09 14:12:50"
        },
        {
            "item":"coffee",
            "search_count":2,
            "updated_at":"2024-03-09 14:13:38"
        }
    ],
    "status":"success",
    "cause":""
}
```

##### その他の使用例

- クエリパラメータを設定しない(データベースに登録されているアイテムから検索数の多い順で上位5件を取得する)

```
curl https://pixfinder.aki158-website.blog/api/total_result
```

- データベースに登録されているアイテム(cを文字列に含む)から日付が新しい順に上位5件を取得する

```
curl https://pixfinder.aki158-website.blog/api/total_result?keyword=c&order=newest
```

- データベースに登録されているアイテム(biを文字列に含む)から検索数の多い順で、1ページあたり10件とした場合の2ページ目を取得する

```
curl https://pixfinder.aki158-website.blog/api/total_result?keyword=bi&page=2&perpage=10
```

#### 補足

設定していないクエリパラメータには、デフォルト値が入ります。

詳細は、[機能一覧](#機能一覧)を確認してください。

## 💾使用技術

<table>
<tr>
  <th>カテゴリ</th>
  <th>技術スタック</th>
</tr>
<tr>
  <td rowspan=5>フロントエンド</td>
  <td>HTML</td>
</tr>
<tr>
  <td>CSS</td>
</tr>
<tr>
  <td>TypeScript</td>
</tr>
<tr>
  <td>ライブラリ : React</td>
</tr>
<tr>
  <td>フレームワーク : Tailwind CSS</td>
</tr>
<tr>
  <td rowspan=1>バックエンド</td>
  <td>Go</td>
</tr>
<tr>
  <td rowspan=4>インフラ</td>
  <td>Amazon EC2</td>
</tr>
<tr>
  <td>Nginx</td>
</tr>
<tr>
  <td>Ubuntu</td>
</tr>
<tr>
  <td>VirtualBox</td>
</tr>
<tr>
  <td>データベース</td>
  <td>MySQL</td>
</tr>
<tr>
  <td rowspan=6>その他</td>
  <td>Git</td>
</tr>
<tr>
  <td>Github</td>
</tr>
<tr>
  <td>SSL/TLS証明書取得、更新、暗号化 : Certbot</td>
</tr>
<tr>
  <td>画像取得 : Google Custom Search API</td>
</tr>
<tr>
  <td>ユーザー認証 : JWT</td>
</tr>
<tr>
  <td>デーモン(常駐プログラム) : Supervisor</td>
</tr>
</table>

## 🗄️ER図

### Imagesテーブル

![er](https://github.com/Recursion-teamB-create-webAPI/Golang-Web-API/assets/119317071/91134f3b-b203-4ba2-9715-6efb35c4f3db)

### Usersテーブル

![db](https://github.com/Recursion-teamB-create-webAPI/Golang-Web-API/assets/119317071/b4cb33e6-1155-47f7-9dfc-814baa650bf5)

## 👀機能一覧

### Webアプリケーション

| 項目 | 内容 | 
| ------- | ------- |
| サインアップ | [URL](#URL)に初めてアクセスするユーザーは、ユーザー登録をしてからアプリケーションを利用することができます。<br>入力には、Username と Password が必要です。 |
| サインイン | サインアップでユーザー登録をしたことがあるユーザーは、サインインからアプリケーションを利用することができます。<br>入力には、Username と Password が必要です。 |
| サインアウト | アプリケーションの利用を終了したい場合は、サインアウトをすることができます。 |
| 画像検索 | WebAPIの4つの機能を集約した、画像検索を体験することができます。<br>入力は、検索欄に直接入力する　または 検索候補から選択することができます。<br>Search ボタンをクリックすることで、検索ワードの画像を見ることができます。 |

### WebAPI

#### 検索機能

検索結果の画像のリストを取得することができます。

| 入力 | 必須 or 任意 | 内容 | 
| ------- | ------- | ------- |
| `/api/search` | 必須 | 対象のエンドポイント |
| `keyword={検索ワード}` | 必須 | クエリパラメータの値で検索します。<br>クエリパラメータを設定していない場合は、データが取得できません。<br>まずは、データベースに検索ワードのアイテムが存在するか確認します。<br>・存在する場合は、データベースからデータを取得します。<br>・存在しない場合は、`Google Custom Search API`で検索した結果をデータとして取得します。<br>`Google Custom Search API`は、1日あたりの検索数が100件まで無料です。<br>`PixFinder`は、無料の範囲内でデータを提供できます。<br>より多くのユーザーが利用できるように一度ユーザーが検索したワードは、データベースにアイテムとして登録され、次回以降同じワードで検索された場合は、データベースを参照するようにしています。<br>これにより、効率的で`Google Custom Search API`に依存しないデータ提供を実現しています。 |

#### 詳細表示

データベースに登録されているデータの詳細を確認することができます。

| 入力 | 必須 or 任意 | 内容 | 
| ------- | ------- | ------- |
| `/api/description` | 必須 | 対象のエンドポイント |
| `keyword={検索ワード}` | 必須 | クエリパラメータの値で検索します。<br>データベースに検索ワードのアイテムが存在するか確認します。<br>・存在する場合は、データベースからデータを取得します。<br>・存在しない場合 または クエリパラメータを設定していない場合は、データが取得できません。 |

#### リスト表示

データベースにどのようなデータがされて登録されているのかを確認することができます。

ほしいデータがデータベースにない場合は、検索機能を利用しデータベースにデータを登録できます。

より詳しいデータを取得したい場合は、詳細表示を利用してください。

| 入力 | 必須 or 任意 | 内容 | 
| ------- | ------- | ------- |
| `/api/list` | 必須 | 対象のエンドポイント |
| `keyword={検索ワード}` | 任意 | クエリパラメータの値で検索します。<br>検索は、検索したいワードの一部分のみの文字列でも検索ができます(部分一致)。<br>例えば、`keyword=bi`の場合、["bird","big"]というデータを返します。<br>(返されるデータは、データベースに登録されているデータによります)<br>データベースに検索ワードのアイテムが存在するか確認します。<br>・存在する場合は、データベースからデータをリストとして取得します。<br>・存在しない場合は、データが取得できません。<br>また、クエリパラメータを設定していない場合は、データベースから全てのアイテムをリストとして取得します。 | 

#### データの集計

データベースに登録されているデータを集計された形で取得することができます。

| 入力 | 必須 or 任意 | 内容 | 
| ------- | ------- | ------- |
| `/api/total_result` | 必須 | 対象のエンドポイント |
| `keyword={検索ワード}` | 任意 | クエリパラメータの値で検索します。<br>データベースに検索ワードのアイテムが存在するか確認します。<br>・存在する場合は、データベースからデータを取得します。<br>・存在しない場合は、データが取得できません。<br>また、クエリパラメータを設定していない場合は、データベースに登録されている全てのアイテムがデータ取得の対象となります。 |
| `page={ページ番号}` | 任意 | 取得する結果のページ番号です。<br>指定できるページ番号は、1〜10になります。<br>クエリパラメータを設定していない場合は、1ページ目を取得します。 |
| `perpage={ページ数}` | 任意 | 1ページあたりに表示するアイテム数です。<br>指定できるページ番号は、1〜10になります。<br>クエリパラメータを設定していない場合は、1ページあたり5件で表示します。 |
| `order={count または newest}` | 任意 | クエリパラメータの値で表示する順番が変化します。<br>・`count` : 検索件数の多い順で表示します。<br>・`newest` : 検索が新しい順で表示します。<br>クエリパラメータを設定していない または 上記の文字列以外をクエリパラメータの値として設定している場合は、`count`が設定されます。 |

#### エラーハンドリング

| 項目 | 内容 | 
| ------- | ------- |
| status | データの取得に失敗すると、`failed`が返されます。 |
| cause | statusが`failed`の場合には、失敗した原因がメッセージとして表示されます。<br>メッセージには、以下のような種類があり、確認後に設定を見直す必要があります。<br>・`Data could not be retrieved because query parameters were not set`・`Google Custom Search API daily usage limit reached or due to server misconfiguration,The image could not be retrieved.`<br>・`Keyword data could not be displayed because it does not exist in the database`<br>・`The item containing the characters set in the query parameter did not exist in the database`<br>・`Query parameters are not set correctly`<br>・`Data could not be retrieved with the set query parameters` |

## 📜作成の経緯

[Recursion](https://recursionist.io/)というサービスで、コンピュータサイエンスを学びました。

次の段階として、学んだ内容をアウトプットとして形にしたいという思いのユーザーが集まり、チーム開発を実施することとなりました。

チーム開発を開始する前に、メンバー間で話し合いをした結果、モダンで多くの豊富なライブラリを提供しているGo言語をメインで使用してWebAPIを作成しようという話になりました。

WebAPIには、様々なものがGithubで公開されており、テーマを決めるのに苦労しました。

最終的に、画像を提供できるAPIに決定し開発がスタートしました。

## ⭐こだわった点

### アクティビティ図

バックエンドの実装を本格的に開始する前に、設計資料としてアクティビティ図を作成しました。

作成した経緯としては、要件について各メンバー間で認識のずれがあると考えたからです。

アクティビティ図は、プロセスの理解やコミュニケーションを助けるツールとして、よく利用されます。

今回は、以下のような5つのアクティビティ図を作成しました。

これにより、チーム内で要件についての共通認識を持つことができ、コミュニケーションを円滑に進めることができました。

また、アクティビティ図をもとに実装したファイルは、以下のリンクから確認することができます。

- [main.go](https://github.com/Recursion-teamB-create-webAPI/Golang-Web-API/blob/main/Backend/main.go)
- [searchHandler.go](https://github.com/Recursion-teamB-create-webAPI/Golang-Web-API/blob/main/Backend/pkg/handlers/searchHandler.go)
- [descriptionHandler.go](https://github.com/Recursion-teamB-create-webAPI/Golang-Web-API/blob/main/Backend/pkg/handlers/descriptionHandler.go)
- [listHandler.go](https://github.com/Recursion-teamB-create-webAPI/Golang-Web-API/blob/main/Backend/pkg/handlers/listHandler.go)
- [totalResultHandler.go](https://github.com/Recursion-teamB-create-webAPI/Golang-Web-API/blob/main/Backend/pkg/handlers/totalResultHandler.go)

#### main関数

![image](https://github.com/Recursion-teamB-create-webAPI/Golang-Web-API/assets/119317071/82ad707a-2005-4805-b806-ff1eb7bfaa49)

#### 検索機能

![image](https://github.com/Recursion-teamB-create-webAPI/Golang-Web-API/assets/119317071/f538c522-53e5-443f-bd3a-bed8dc27e8dd)

#### 詳細表示

![image](https://github.com/Recursion-teamB-create-webAPI/Golang-Web-API/assets/119317071/376a6df8-975f-490e-ba7e-f4d6b964cc66)

#### リスト表示

![image](https://github.com/Recursion-teamB-create-webAPI/Golang-Web-API/assets/119317071/00be581e-726c-4166-82cb-60a70b81775d)

#### データの集計

![image](https://github.com/Recursion-teamB-create-webAPI/Golang-Web-API/assets/119317071/683495e0-67e7-46f6-9aec-cbf46df1b54a)

### データアクセス層

データアクセス層は、ビジネスロジックを扱うアプリケーション層とデータ層（MySQLのようなSQL RDBMSなど）の間の橋渡しをします。

今回は、データアクセス層として[dao.go](https://github.com/Recursion-teamB-create-webAPI/Golang-Web-API/blob/main/Backend/pkg/dao/dao.go)と[dao_users.go](https://github.com/Recursion-teamB-create-webAPI/Golang-Web-API/blob/main/Backend/pkg/dao/dao_users.go)を作成しました。

これは、MySQLへのCRUD操作や接続などに使用されます。

これを利用することにより、ビジネスロジックのコードに影響を与えることなく、データ層との連携を担うコードの多くを容易に交換することができます。

例えば、今回は、データ層にMySQLを使用していますが、MongoDBのような他のDBに交換することもできます。

これにより、コードの可読性やメンテナンスが容易になります。

### バックエンドの単体テスト

バックエンドには、Go言語を使用しており、各関数毎に単体テストを行いコードの信頼性を担保しました。

Go言語は、豊富なライブラリを提供しており、テストコードには、`net/http/httptest`と`github.com/DATA-DOG/go-sqlmock`を利用しました。

用途は、以下の通りです。

- `net/http/httptest` : ハンドラーに関わる関数に利用し、リクエスト/レスポンスの検証を行いました。
- `github.com/DATA-DOG/go-sqlmock`　: データ層に関わる関数に利用し、クエリの検証を行いました。

### チーム開発

チーム開発では、各メンバー間でのコミュニケーションは不可欠です。

コミュニケーションツールとして、以下のようなアプリケーションを利用しました。

| アプリケーション | 目的 |
| ------- | ------- |
| Discord | チャットでの連絡ややりとり |
| GitHubのissue | 課題管理 |
| Ovice | ミーティング |

ミーティングでは、チーム開発が順調か、困りごとや改善点がないかなどを確認しました。

ミーティング実施後は、アクションの確認や議論の内容を見返すために議事録を作成しました。

議事録には、3つの項目について定期的にチームでミーティングを開催し議論した内容を記載しています。

| 項目 | 内容 |
| ------- | ------- |
| やったこと | 前回のミーティング以降に各メンバーで実施した内容を共有します。<br>計画通りに進められているか確認します。<br>計画から遅れている場合は、各メンバーの進捗と合わせて検討し、マイルストーンを後ろにずらすか他メンバーで作業を巻き取るかなどを議論します。 |
| 直面した問題 | 仕様の確認や設計、実装時にでた不明点や問題、困りごとを共有します。<br>共有することで、再発防止やメンバー同士の共通認識の確認、改善点などがわかります。 |
| 今後の課題 | 次回のミーティングまでに、各メンバーが実施する内容を記載しています。<br>ここで決めた内容が、次回のミーティングまでの各メンバーのアクションになります。 |

議事録は、以下リンクから確認してください。

[dev-log.md](https://github.com/Recursion-teamB-create-webAPI/team-b-devlog/blob/main/dev-log.md)

## 📑参考文献
### 公式ドキュメント
- [MySQL](https://dev.mysql.com/doc/refman/8.0/en/innodb-online-ddl-operations.html)
- [Custom Search JSON API](https://developers.google.com/custom-search/v1/overview?hl=ja)

### 参考にしたサイト
- [A Tour of Go](https://go-tour-jp.appspot.com/list)
- [Recursion_Go入門](https://recursionist.io/learn/languages/go)

