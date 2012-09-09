ADBS
==============================================

## Introduction ##

The adbs is a simple tool for adb(Android Debug Bridge).

When you are attaching several devices to the PC/Mac for testing or debugging
you must specify serial number of the target device
to use adb command.
So you must use 'adb devices' to look for the device's serial number at first,
then copy it, and make the complete command with pasting the number...

The adbs makes this process simple.
You just specify the first character of the serial number, so you do not
have to copy and paste the correct serial number.

For example, if the result of the adb devices is following:

    $ adb devices
    List of devices attached
    304D19E0D41F543F  device
    275501700028393   device

then you can do like this:

    $ adbs -s 3 shell
    shell@android:/ $

or you can execute without '-s' option. The adbs will ask it to you.
For example:

    $ adbs shell
    List of devices attached
    [3] 304D19E0D41F543F
    [2] 275501700028393
    1st character of the serial number you want to use: 2
    shell@android:/ $

## Installing adbs ##

Just put the 'adbs' file to the directory like /usr/local/bin.

## License ##

Copyright © 2012 Soichiro Kashima.

Licensed under MIT License.
See the bundled LICENSE file for details.
