# fuzzl

**fuzzl** is a Go package for experimenting with shader-inspired noise functions.  
It provides simple APIs for generating structured noise values from 3D vectors.

The first implementation included here is `DotNoise`, a port of Xor’s **dot_noise** function  
from [GM Shaders](https://mini.gmshaders.com/p/phi).

> GM Shaders is a reader-supported publication.  
> To support the author, consider becoming a free or paid subscriber:  
> [gmshaders.com](https://gmshaders.com)

---

## Installation

```bash
go get github.com/mdhender/fuzzl
````

---

## Usage

```go
package main

import (
	"fmt"

	"github.com/mdhender/fuzzl"
)

func main() {
	p := fuzzl.Vec3{1.0, 2.0, 3.0}
	value := fuzzl.DotNoise(p)

	fmt.Println("DotNoise at", p, "=", value)
}
```

Output (deterministic for given inputs):

```
DotNoise at [1 2 3] = -0.6843639602299883
```

---

## API

* **`type Vec3 [3]float64`**
  A minimal 3D vector with helper methods for dot products and scalar ops.

* **`func DotNoise(p Vec3) float64`**
  Returns a structured noise value in the approximate range `[-3, +3]`.

---

## Roadmap

* Additional noise generators (gyroid, Perlin-like, turbulence).
* Expanded vector/matrix utilities for compact math.
* Benchmarks and performance profiling.

---

## License

MIT
