# CA-techtrain-mission

## 説明
これはTechTrainのミッション[「オンライン版　CA Tech Dojo サーバサイド (Go)編」](https://techbowl.co.jp/techtrain/missions/12)を行うために作成したリポジトリです。
API定義はミッションを参考に、実装は全て自分で行いました
メンターさんからフィードバックをもらい改良をしています

## 起動方法
dockerが必須です
```
make compose/build && make compose/up
```

## e2eテスト
go+dockerが必須です
```
make test/e2e
```