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

func TestTupleEquality(t *testing.T) {
    v1 := utils.Vector(4.3, -4.2, 3.1)
    v2 := utils.Vector(4.3, -4.2, 3.1)
    HelperTestTupleEquality(t,v1, v2)

    p1 := utils.Point(4.3, -4.2, 3.1)
    p2 := utils.Point(4.3, -4.2, 3.1)
    HelperTestTupleEquality(t,p1, p2)
}


func HelperTestTupleEquality(t *testing.T, t1 utils.Tuple, t2 utils.Tuple) {
    if !t1.Equals(t2) {
        t.Errorf("%v is close enough to %v. Got different.", t1, t2)
    }

func TestSetTupleComponents(t *testing.T) {
    v := utils.Vector(1.5, 2.6, -7.3)
    p := utils.Point(7.6, -2.9, -17.0)

    v.SetComp(0, 2)
    v.SetComp(1, 9.2)
    v.SetComp(2, 4)
    vArr := v.AsArray()
    if (vArr[0] != 2.0 ||
        vArr[1] != 9.2 ||
        vArr[2] != 4.0 ||
        vArr[3] != 0.0) {

        t.Errorf("Unsuccessful component setting. Got %v", vArr)
    }

    p.SetComp(0, 3)
    p.SetComp(1, 9)
    p.SetComp(2, 4.0)
    pArr := p.AsArray()
    if (pArr[0] != 3.0 ||
        pArr[1] != 9.0 ||
        pArr[2] != 4.0 ||
        pArr[3] != 1.0) {

        t.Errorf("Unsuccessful component setting. Got %v", pArr)
    }

    v.SetComps(1, 2, 3.1415)
    vArr = v.AsArray()
    if (vArr[0] != 1.0 ||
        vArr[1] != 2.0 ||
        vArr[2] != 3.1415 ||
        vArr[3] != 0.0) {

        t.Errorf("Unsuccessful component setting. Got %v", vArr)
    }

    p.SetComps(5, 8, 13.21)
    pArr = p.AsArray()
    if (pArr[0] != 5.0 ||
        pArr[1] != 8.0 ||
        pArr[2] != 13.21 ||
        pArr[3] != 1.0) {

        t.Errorf("Unsuccessful component setting. Got %v", pArr)
    }
}
