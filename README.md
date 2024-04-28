## Command

```bash
curl "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"Name":"John"}'
```

# Clean Architecture

![clean_arch.png](docs%2Fimg%2Fclean_arch.png)

| レイヤー名                      | 含む概念                                          |
|----------------------------|-----------------------------------------------|
| Enterprise Business Rules  | Entities                                      |
| Application Business Rules | Use Cases / Interactor                        |
| Interface Adapters         | Controllers / Presenters / Gateways           |
| Frameworks & Drivers       | Web / UI / External Interfaces / Devices / DB |

## EnterPrise Business Rules

ビジネスロジックを定義する

| ディレクトリ     | 意味                        |
|------------|---------------------------|
| domain     | ビジネスモデル                   |
| repository | domainを永続化するためのインターフェース定義 |

## Application Business Rules

Enterprise Business Rulesのクライアントとしてエンティティを操作する

- usecase
- interactor

## Interface Adapters

Frameworks & DriversとApplication Business Rulesとの間でデータの相互変換を行う

| ディレクトリ     | 意味                                   |
|------------|--------------------------------------|
| controller | 外部からの入力を受け取りInputDataに変換してusecaseに渡す |
| handler    | Frameworks&Driversからのデータを抽象化する       |

## Frameworks & Drivers

フレームワークやDBなど、詳細な技術を配置する

- infra

## handler
## DI
- registry