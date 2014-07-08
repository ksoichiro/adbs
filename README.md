# adbs

[![Build Status](https://drone.io/github.com/ksoichiro/adbs/status.png)](https://drone.io/github.com/ksoichiro/adbs/latest)

`adbs` is a simple tool for adb(Android Debug Bridge).  
Makes it easy to select one of the multiple attached devices when using `adb`.

[Japanese document](README.ja.md)

## Install

[Get the latest binary](https://github.com/ksoichiro/adbs/releases/latest) and locate it to somewhere in your `PATH` or install by go:

```sh
go get github.com/ksoichiro/adbs
```

## Introduction

When you are attaching multiple devices to the PC/Mac
for testing or debugging, you must specify serial number of the target device
to use `adb` command.
So you must use `adb devices` to look for the device's serial number at first,
then copy it, and make the complete command with pasting the number...

`adbs` makes this process simple.
You just specify the first character of the serial number, so you do not
have to copy and paste the correct serial number.

For example, if the result of the `adb devices` is following:

```sh
$ adb devices
List of devices attached
304D19E0D41F543F  device
275501700028393   device
```

Normally, you will see the error like this:

```sh
$ adb shell
error: more than one device and emulator
# Oops.. I should specify serial number..
```

With `adbs`, you can avoid this problem:

```sh
$ adbs -s 3 shell
adbs: serial: 304D19E0D41F543F
shell@android:/ $
```

or you can execute without `-s` option. `adbs` will ask it to you.
For example:

```sh
$ adbs shell
[1] 304D19E0D41F543F
[2] 275501700028393
Device to execute command: 1
Specified: 304D19E0D41F543F
shell@android:/ $
```

## License

Copyright © 2012 Soichiro Kashima.

Licensed under MIT License.
See the bundled LICENSE file for details.
