ADBS
==============================================

## 概要 ##

adbsは、adb(Android Debug Bridge)のためのシンプルなツールです。

テストやデバッグの目的で複数の端末をPC/Macに接続しているとき、特定の端末に対してadbコマンドでの操作をする場合はシリアル番号を指定する必要があります。
そのため、まず'adb devices'を実行して対象端末のシリアル番号を探し、コピー＆ペーストしてコマンドを作るという面倒な手順を踏むことになると思います。

adbsは、この手順をシンプルにします。
adbsではシリアル番号の最初の1文字を指定するだけで端末を特定して実行できるようになるので、正確なシリアル番号をコピー＆ペーストする必要がありません。

例えば、adb devicesの結果が以下のようなものだったとします。

    $ adb devices
    List of devices attached
    304D19E0D41F543F  device
    275501700028393   device

この場合は以下のように実行できます。

    $ adbs -s 3 shell
    shell@android:/ $

もしくは、'-s'オプションなしで実行すればadbsが聞いてきます。
例えば以下のようになります。

    $ adbs shell
    List of devices attached
    [3] 304D19E0D41F543F
    [2] 275501700028393
    1st character of the serial number you want to use: 2
    shell@android:/ $

## インストール ##

'adbs'ファイルをパスの通ったディレクトリ(/usr/local/binなど)に置くだけです。

## License ##

Copyright (c) 2012 Soichiro Kashima.

MITライセンスでライセンスされています。
詳しくは同梱のLICENSEファイルをご覧ください。
