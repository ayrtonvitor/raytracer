package utils

import (
	"errors"
	"math"
)

type Tuple struct {
    x, y, z, w float64
}

func (t *Tuple) AsArray() [4]float64 {
    if t.w != 0.0 && t.w != 1.0 {
        panic("There is a tuple that is not a point or vector (w != 0 && w != 1)")
    }
    return [...]float64{t.x, t.y, t.z, t.w}
}

func (t *Tuple) Equals(u Tuple) bool {
     return (FloatEqual(t.x, u.x) &&
         FloatEqual(t.y, u.y) &&
         FloatEqual(t.z, u.z) &&
         FloatEqual(t.w, u.w))
}

func (t *Tuple) SetComp(i int, val float64) {
    switch i {
    case 0:
        t.x = val
    case 1:
        t.y = val
    case 2:
        t.z = val
    }
}

func (t *Tuple) SetComps(x float64, y float64, z float64) {
    t.x = x
    t.y = y
    t.z = z
}

func (t1 *Tuple) Add(t2 Tuple) (Tuple, error) {
    if t1.w == 1 && t2.w == 1 {
        return *t1, errors.New("Can't add two points")
    }
    t := Tuple{t1.x + t2.x,
               t1.y + t2.y,
               t1.z + t2.z,
               t1.w + t2.w}
    return t, nil
}

func (t1 *Tuple) Minus(t2 Tuple) (Tuple, error) {
    res := Tuple{t1.x - t2.x,
                 t1.y - t2.y,
                 t1.z - t2.z,
                 t1.w - t2.w}
    if res.w != 0.0 && res.w != 1.0 {
        return *t1, errors.New("Cannot subtract points from vectors.")
    }

    return res, nil
}

func (t *Tuple) Opposite() Tuple {
    if t.w == 1 {
        return *t
    }
    return Tuple{-t.x, -t.y, -t.z, 0}
}

func (t *Tuple) ScalarMult(c float64) Tuple {
    if t.w == 1 {
        return *t
    }
    return Vector(c * t.x, c * t.y, c * t.z)
}

func (t *Tuple) Div(c float64) Tuple {
    if t.w == 1 {
        return *t
    }
    return Vector(t.x / c, t.y / c, t.z / c)
}

func (t *Tuple) Norm() float64 {
    return math.Sqrt(t.x * t.x + t.y * t.y + t.z * t.z)
}

func (t *Tuple) Normalize() Tuple {
    return t.Div(t.Norm())
}

func Point(x float64, y float64, z float64) Tuple {
    t := Tuple{x, y, z, 1.0}
    return t
}

func Vector(x float64, y float64, z float64) Tuple {
    t := Tuple{x, y, z, 0.0}
    return t
}
