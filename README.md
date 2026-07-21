
# muxt

a tool to provide a declarative tmux experience.

## build

just run:
```sh
make
```
or with a custom prefix:
```sh
make PREFIX=~/some/local
```

and if you want to install the binary run: 
```sh 
make install # it needs root
```

## usage

first you need to create a layout, you can it with:
```sh
muxt new <layout-name>
```
this will create a new file inside $XDG_CONFIG_HOME/muxt/layouts called `<layout-name>.kdl`.
a base start template will be write to the file.

you can edit this layout later with:
```sh
muxt edit <layout-name>
```

finally, to run the layout and create a session with it, use:
```sh
muxt start <layout-name>
```

### writing layouts
muxt uses the [kdl language](https://kdl.dev) to declare the layouts. see the [examples directory](./examples) for more.

