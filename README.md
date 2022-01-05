light
=====

`light [amount]`

Print backlight brightness percentage to STDOUT. Optionally change the
brightness to either a set value or by adding or subtracting from the current
value.

`light` - Print backlight percent
`light 50` - Set backlight to 50%
`light +10` - Increase backlight brightness by 10
`light -10` - Decrease backlight brightness by 10

Build
------

Build dependencies  

 * golang
 * make

`make all`

Install
--------

Optionally configure `config.mk` to specify a different install location.  
Defaults to `/usr/local/`

`sudo make install`

Uninstall
----------

`sudo make uninstall`

License
--------

GPL3: See LICENSE for details.
Copyright (C) 2022 Dakota Walsh
