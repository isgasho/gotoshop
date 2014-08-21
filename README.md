# Gotoshop

Gotoshop is a image manipulation server purely written in golang. It provides a
set of basic image manipulation operations such as scaling, cropping and
rotating.

### Build & Run

Gotoshop uses [gom](https://github.com/mattn/gom) for dependencies management.
If you don't have gom, install it first.

```
$ gom install && gom build
$ ./gotoshop
```

### General Headers

Type of returning image is specified by `Accept` header.

```
Accept: image/(png|jpeg|gif)
```

### General Params

Image to manipulate is specified by `url` param. The `url` must be accessible.

```
{
    url: <string>,
}
```

### GET /blur

Blurs an image using a Gaussian blur. `std_dev` is the standard deviation of
the normal, higher is blurrier. `size` is the size of the kernel. If zero, it
is set to ceil(6 * `std_dev`).

```
{
    std_dev: <float>,
    size: <int>, // optional
}
```

Example

http://gotoshop.herokuapp.com/blur?url=http%3A%2F%2Fblog.golang.org%2Fgo-at-google-io-2011-videos_gopher.jpg&std_dev=3

### GET /crop

Crops an image. Does not scale. `(x0, y0)` is top left, `(x1, y1)` is bottom
right of a rectangle. Returns 406 if one ore more points are out of image
bounds.

```
{
    x0: <int>,
    y0: <int>,
    x1: <int>,
    y1: <int>,
}
```

Example

http://gotoshop.herokuapp.com/crop?url=http%3A%2F%2Fblog.golang.org%2Fgo-at-google-io-2011-videos_gopher.jpg&x0=100&y0=100&x1=200&y1=200

### GET /flip

Flips an image vertically or horizontally.


```
{
    axis: (x|y),
}
```

Example

http://gotoshop.herokuapp.com/flip?url=http%3A%2F%2Fblog.golang.org%2Fgo-at-google-io-2011-videos_gopher.jpg&axis=x

### GET /rotate

Rotates an image. `angle` is the angle, in degree, to rotate the image
clockwise.

```
{
    angle: <float>,
}
```

Example

http://gotoshop.herokuapp.com/rotate?url=http%3A%2F%2Fblog.golang.org%2Fgo-at-google-io-2011-videos_gopher.jpg&angle=1.0

### GET /scale

Scales an image. `ratio` must be larger than 0.

```
{
    ratio: <float>,
}
```

Example

http://gotoshop.herokuapp.com/scale?url=http%3A%2F%2Fblog.golang.org%2Fgo-at-google-io-2011-videos_gopher.jpg&ratio=0.5

### GET /thumbnail

Scales and crops source image to fit `width` and `height`.

```
{
    width: <int>,
    height: <int>,
}
```

Example

http://gotoshop.herokuapp.com/thumbnail?url=http%3A%2F%2Fblog.golang.org%2Fgo-at-google-io-2011-videos_gopher.jpg&width=100&height=100
