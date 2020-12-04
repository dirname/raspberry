# Touch Sensor

> Use touch to wake up the screen, which is more accurate than human infrared sensor wake up

Use a trigger type touch sensor to turn off the screen with a single tap Double tap to turn on the screen

![demo](demo.gif)

Due to the policy of `xauth`, you need add entry:

+ run `xauth.sh` as user `pi`

```shell
sh xauth.sh
```

+ run in initialization function:
```python
display = os.popen("sudo -u pi xauth list :0.0").read()
os.system("sudo -S xauth add " + display)
```

+ manually set
```shell
xauth list :0.0
sudo -S xauth add $value
```

`:0.0` is a normal default monitor if you have multi-monitor to change this value.


`run.py` ~~is a script what can make the buzzer sound.~~ (But not required.)