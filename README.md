# APApp

## 概要
過去問道場で試験後に出力される CSV を読み取り、DB に保存し、  
試験単位での得点率や苦手分野の表示などをグラフで可視化するアプリです。

---

## 使用技術

### Backend
- 言語: Go  
- フレームワーク: Gin
- ORM: GORM（DB 操作用）  
- テスト: Testify（ユニットテストフレームワーク）  
- ロガー: Zap（構造化ロガー）

### Frontend
- 言語: TypeScript  
- フレームワーク: Next.js 15 (App Router)  
- スタイル: Tailwind CSS  
- UIコンポーネント: Shadcn/UI

### Infrastructure
- SQLite（ローカル開発環境）  
- PostgreSQL（本番環境予定）

---

## 開発状況
- バックエンド構築完了  
- 現在、ユニットテストを実施中（controllers層・services層）  
- 今後、フロントエンドの実装を予定（CSVアップロード、成績可視化UI）

---

## 今後の予定
- Dockerによる開発環境構築  
- CI/CD導入  
- グラフ描画機能（Recharts または Chart.js の導入）

---

## ディレクトリ構成（予定）
