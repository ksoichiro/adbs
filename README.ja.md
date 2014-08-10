# adbs

[![Build Status](https://drone.io/github.com/ksoichiro/adbs/status.png)](https://drone.io/github.com/ksoichiro/adbs/latest)

`adbs`は、adb(Android Debug Bridge)のためのシンプルなツールです。  
接続されている複数端末から1つを選んで`adb`を実行するのが簡単になります。

## インストール

以下のような方法でインストールできます。

### バイナリ

[最新のバイナリをダウンロード](https://github.com/ksoichiro/adbs/releases/latest)してパスの通ったディレクトリに配置します。

### Go

```sh
$ go get github.com/ksoichiro/adbs
```

### [Homebrew](http://brew.sh/)

```sh
$ brew tap ksoichiro/adbs
$ brew install adbs
```

## 概要

テストやデバッグの目的で複数の端末をPC/Macに接続しているとき、特定の端末に対して`adb`コマンドでの操作をする場合はシリアル番号を指定する必要があります。
そのため、まず`adb devices`を実行して対象端末のシリアル番号を探し、コピー＆ペーストしてコマンドを作るという面倒な手順を踏むことになると思います。

`adbs`は、この手順をシンプルにします。
`adbs`ではシリアル番号の最初の1文字を指定するだけで端末を特定して実行できるようになるので、正確なシリアル番号をコピー＆ペーストする必要がありません。

例えば、`adb devices`の結果が以下のようなものだったとします。

```sh
$ adb devices
List of devices attached
304D19E0D41F543F  device
275501700028393   device
```

この状況では、以下のようなことが起きるはずです。

```sh
$ adb shell
error: more than one device and emulator
# しまった、シリアル番号を入力しないと実行できない...
```

`adbs`を使うとこの問題を回避できます。

```sh
$ adbs -s 3 shell
adbs: serial: 304D19E0D41F543F
shell@android:/ $
```

もしくは、'-s'オプションなしで実行すれば`adbs`が聞いてきます。  
例えば以下のようになります。

```sh
$ adbs shell
[1] 304D19E0D41F543F
[2] 275501700028393
Device to execute command: 1
Specified: 304D19E0D41F543F
shell@android:/ $
```

## License ##

Copyright (c) 2012 Soichiro Kashima.

MITライセンスでライセンスされています。
詳しくは同梱のLICENSEファイルをご覧ください。
