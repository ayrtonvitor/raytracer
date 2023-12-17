package utils

import "math"

func FloatEqualWithPrecision(x float64, y float64, epsilon float64) bool {

    if math.Abs(x - y) < epsilon {
        return true
    } else {
        return false
    }
}

func FloatEqual(x float64, y float64) bool {
    epsilon := 0.00001
    return FloatEqualWithPrecision(x, y, epsilon)
}

func FloatXLessThanYWithPrecision(x float64, y float64, epsilon float64) bool {
    if (!FloatEqualWithPrecision(x, y, epsilon) && x < y) {
        return true
    }
    return false
}

func FloatXLessThanY(x float64, y float64) bool {
    epsilon := 0.00001
    return FloatXLessThanYWithPrecision(x, y, epsilon)
}

func FloatXGreaterThanYWithPrecision(x float64, y float64, epsilon float64) bool {
    if (!FloatEqualWithPrecision(x, y, epsilon) && x > y) {
        return true
    }
    return false
}

func FloatXGreaterThanY(x float64, y float64) bool {
    epsilon := 0.00001
    return FloatXGreaterThanYWithPrecision(x, y, epsilon)
}

func FloatXLEQYWithPrecision(x float64, y float64, epsilon float64) bool {
    if (x < y || !FloatEqualWithPrecision(x, y, epsilon)) {
        return true
    }
    return false
}

func FloatXLEQY(x float64, y float64) bool {
    epsilon := 0.00001
    return FloatXLessThanYWithPrecision(x, y, epsilon)
}

func FloatXGEQYWithPrecision(x float64, y float64, epsilon float64) bool {
    if (x > y || !FloatEqualWithPrecision(x, y, epsilon)) {
        return true
    }
    return false
}

func FloatXGEQY(x float64, y float64) bool {
    epsilon := 0.00001
    return FloatXGreaterThanYWithPrecision(x, y, epsilon)
}
