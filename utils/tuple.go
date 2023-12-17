package utils

type Tuple struct {
    x, y, z, w float64
}

func (t *Tuple) AsArray() [4]float64 {
    if t.w != 0.0 && t.w != 1.0 {
        panic("There is a tuple that is not a point or vector (w != 0 && w != 1)")
    }
    return [...]float64{t.x, t.y, t.z, t.w}
}

func Point(x float64, y float64, z float64) Tuple {
    t := Tuple{x, y, z, 1.0}
    return t
}


func Vector(x float64, y float64, z float64) Tuple {
    t := Tuple{x, y, z, 0.0}
    return t
}
