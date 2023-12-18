package utils

import (
	"math"
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
    t2Arr := t2.AsArray()

    t2.SetComps(t2Arr[0] + 0.0000001,
                t2Arr[1] + 0.0000001,
                t2Arr[2] + 0.000001)
    if !t1.Equals(t2) {
        t.Errorf("%v is close enough to %v. Got different.", t1, t2)
    }
    t2.SetComps(t2Arr[0] - 0.0000001,
                t2Arr[1] - 0.0000001,
                t2Arr[2] - 0.000001)

    t2.SetComp(0, t2Arr[0] + 0.01)
    if t1.Equals(t2) {
        t.Errorf("%v is not close enough to %v. Got equal.", t1, t2)
    }

    t2.SetComp(0, t2Arr[0] - 0.01)
    t2.SetComp(1, t2Arr[1] - 0.01)
    if t1.Equals(t2) {
        t.Errorf("%v is not close enough to %v. Got equal.", t1, t2)
    }

    t2.SetComp(1, t2Arr[0] + 0.01)
    t2.SetComp(2, t2Arr[2] + 0.01)
    if t1.Equals(t2) {
        t.Errorf("%v is not close enough to %v. Got equal.", t1, t2)
    }
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

func TestTupleAddition(t *testing.T) {
    p1 := utils.Point(3, -2, 5)
    v := utils.Vector(-2, 3, 1)

    p2, err := p1.Add(v)
    if !p2.Equals(utils.Point(1, 1, 6)) {
        t.Errorf("Error adding tuples. %v + %v != %v", p1, v, p2)
    } else if err != nil {
        t.Errorf("%v", err)
    }

    p3 := utils.Point(5, 0, -1)
    p2, err = p1.Add(p3)
    if err == nil {
        t.Errorf("Two pointers were added.\n%v\n%v", p1, p3)
    }
}

func TestTupleSubtraction(t *testing.T) {
    p1 := utils.Point(3, 2, 1)
    p2 := utils.Point(5, 6, 7)
    v1 := utils.Vector(8, 9, 0)
    v2 := utils.Vector(3, 4, 5)

    pmp, err := p1.Minus(p2)
    pmv, err := p1.Minus(v1)
    vmv, err := v1.Minus(v2)

    HelperTestTupleSubtraction(t, pmp, err, -2, -4, -6, 0)
    HelperTestTupleSubtraction(t, pmv, err, -5, -7, 1, 1)
    HelperTestTupleSubtraction(t, vmv, err, 5, 5, -5, 0)

    vmp, err := v1.Minus(p1)
    if err == nil {
        t.Errorf("%v - %v is not %v. Undefined behavior", v1, p1, vmp)
    }
}

func HelperTestTupleSubtraction(t *testing.T, tp utils.Tuple, err error, x float64,
    y float64, z float64, w float64) {

    if tp.AsArray() != [...]float64{x, y, z, w} {
        t.Errorf("got %v, expected (%v, %v, %v, %v)", tp, x, y, z, w)
        return
    }
 
    if err != nil {
        t.Errorf("%v", err)
    }
}

func TestNegateTuples(t *testing.T) {
    v := utils.Vector(-1, 2, 3)
    v = v.Opposite()

    if v != utils.Vector(1, -2, -3) {
        t.Errorf("Got %v. Expected (-1, 2, 3, 0)", v)
    }

    p := utils.Point(1, 3, 5)
    p = p.Opposite()
    if p != utils.Point(1, 3, 5) {
        t.Errorf("Got %v. Expected (1, 3, 5, 1)", p)
    }
}

func TestScalarMultiplicationTest(t *testing.T) {
    v := utils.Vector(1, -2, 3)

    if v.ScalarMult(3.5) != utils.Vector(3.5, -7, 10.5) {
        t.Errorf("Got %v. Expected (3.5, -7, 10.5, 0)", v)
    }
}

func TestScalarDivision(t *testing.T) {
    v := utils.Vector(1, -2, 3)

    if v.Div(2.0) != utils.Vector(0.5, -1, 1.5) {
        t.Errorf("Got %v. Expected (0.5, -1, 1.5, 0)", v)
    }
}

func TestVectorNorm(t *testing.T) {
    v1 := utils.Vector(1, 0, 0)
    v2 := utils.Vector(0, 1, 0)
    v3 := utils.Vector(0, 0, 1)
    v4 := utils.Vector(1, 2, 3)
    v5 := utils.Vector(-1, -2, -3)

    n1 := v1.Norm()
    n2 := v2.Norm()
    n3 := v3.Norm()
    n4 := v4.Norm()
    n5 := v5.Norm()

    if n1 != 1 {
        t.Errorf("Got %v. Expected 1", n1)
    }
    if n2 != 1 {
        t.Errorf("Got %v. Expected 1", n2)
    }
    if n3 != 1 {
        t.Errorf("Got %v. Expected 1", n3)
    }
    if !utils.FloatEqual(n4, math.Sqrt(14)) {
        t.Errorf("Got %v. Expected %v", n4, math.Sqrt(14))
    }
    if !utils.FloatEqual(n5, math.Sqrt(14)) {
        t.Errorf("Got %v. Expected %v", n5, math.Sqrt(14))
    }
}
