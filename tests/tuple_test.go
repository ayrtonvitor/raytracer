package utils

import (
    "testing"
    "github.com/ayrtonvitor/raytracer/utils"
)


func TestTupleWithw1IsPoint(t *testing.T) {
    p := utils.Point(4.3, -4.2, 3.1)
    valArr := p.AsArray()
    if valArr[0] != 4.3 {
        t.Errorf("x is %v. Expected 4.3", valArr[0])
    }
    if valArr[1] != -4.2 {
        t.Errorf("y is %v. Expected -4.2", valArr[1])
    }
    if valArr[2] != 3.1 {
        t.Errorf("z is %v. Expected 3.1", valArr[2])
    }
    if valArr[3] != 1.0 {
        t.Errorf("Tuple is not a point. Got w != 1.0")
    }
    if p == utils.Vector(4.3, -4.2, 3.1) {
        t.Errorf("p is a also a vector")
    }
}

func TestTupleWithw0IsVector(t *testing.T) {
    p := utils.Vector(4.3, -4.2, 3.1)
    valArr := p.AsArray()
    if valArr[0] != 4.3 {
        t.Errorf("x is %v. Expected 4.3", valArr[0])
    }
    if valArr[1] != -4.2 {
        t.Errorf("y is %v. Expected -4.2", valArr[1])
    }
    if valArr[2] != 3.1 {
        t.Errorf("z is %v. Expected 3.1", valArr[2])
    }
    if valArr[3] != 0.0 {
        t.Errorf("Tuple is not a point. Got w != 0.0")
    }
    if p == utils.Point(4.3, -4.2, 3.1) {
        t.Errorf("p is a also a point")
    }
}
