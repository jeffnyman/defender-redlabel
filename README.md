<h1 align="center">

<img src="https://raw.githubusercontent.com/jeffnyman/defender-redlabel/master/assets/defender-title.jpg" alt="Defender"/>

</h1>

<p align="center"><strong>RED LABEL VERSION</strong></p>

<p align="center">

<img src="https://raw.githubusercontent.com/jeffnyman/defender-redlabel/master/assets/defender-red-label-chip.jpg" alt="Red Label Chip"/>

</p>

This repository will be a Go-based implementation of the extracted, reassembled and rebuilt ROM from my [Retro Defender](https://github.com/jeffnyman/defender-retro) project.

This is a Go project, so all you should have to do once you have [Go installed on your system](https://go.dev/dl/) is:

```sh
go build
```

That will generate a `defender-redlabel` executable file.

This is currently an incomplete implementation based on how the logic is being generated, although it is workable in the sense that you can fly the ship around and shoot aliens! The controls are as follows:

```
MOVE UP     = Up Arrow
MOVE DOWN   = Down Arrow
REVERSE     = Space Bar
THRUST      = Enter
HYPERSPACE  = H
FIRE LASER  = F
SMART BOMB  = B
```

The key bindings are currently set up in the `gl/gl.go` file if you want to change them.
