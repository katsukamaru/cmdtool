##　aws.Stringを使う理由
- [StringValue used instead of String type](https://github.com/aws/aws-sdk-go/issues/114)

## 参考
- [GoでAWS SDKを叩くCLIツールを作ってリリースするまでの流れ](https://qiita.com/minamijoyo/items/bb21a111882cb81213ab)
- [GolangでAPI Clientを実装する](http://deeeet.com/writing/2016/11/01/go-api-client/)


## Cobraの使い方
- 基本的には[GoでAWS SDKを叩くCLIツールを作ってリリースするまでの流れ](https://qiita.com/minamijoyo/items/bb21a111882cb81213ab)に書いているものを踏襲する
```
## CobraでScafoldを使う
$ cobra init

$ vim $HOME/.cobra.yaml
---
author: Katsukamaru
license: MIT
---

$ go install ## $GOPATH/binにビルド成果物がインストールされる。PATHを通しておくと良い。

## サブコマンドを作る
$ cobra add ec2

## サブサブコマンド
$ cobra add ec2ls -p 'ec2Cmd'

$ go run main.go ec2 ec2ls


$ go run main.go ec2 ls
```

## aws-sdk-goについて
github.com/awslabs/aws-sdk-go は古い実装
github.com/aws/aws-sdk-goを使うべき


