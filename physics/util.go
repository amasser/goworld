package physics

import (
    "github.com/ianremmler/ode"
    mgl "github.com/go-gl/mathgl/mgl32"
    "math"
)

const (
    deg2rad = math.Pi / 180.0
    rad2deg = 1.0 / deg2rad
)

func FromOdeVec3(vec ode.Vector3) mgl.Vec3 {
    return mgl.Vec3 {
        float32(vec[0]),
        float32(vec[1]),
        float32(vec[2]),
    }
}

func ToOdeVec3(vec mgl.Vec3) ode.Vector3 {
    return ode.Vector3 {
        float64(vec[0]),
        float64(vec[1]),
        float64(vec[2]),
    }
}

func FromOdeRotation(mat3 ode.Matrix3) mgl.Vec3 {

    x := math.Atan2(mat3[2][1], mat3[2][2])
    y := math.Atan2(-mat3[2][0], math.Sqrt(math.Pow(mat3[2][1], 2) + math.Pow(mat3[2][2], 2)))
    z := math.Atan2(mat3[1][0], mat3[0][0])
    return mgl.Vec3 { float32(x * rad2deg), float32(y * rad2deg), float32(z * rad2deg) }
}