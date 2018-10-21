package main

import "math"

type Vec3 struct {
	x, y, z float
}

func (v *Vec3) ToColor() *Color {
	c := &Color{}
	c.r = uint8(255.0 * math.Sqrt(v.x))
	c.g = uint8(255.0 * math.Sqrt(v.y))
	c.b = uint8(255.0 * math.Sqrt(v.z))
	return c
}

func (v *Vec3) ToGammaColor() *Color {
	c := &Color{}
	c.r = uint8(255.0 * v.x)
	c.g = uint8(255.0 * v.y)
	c.b = uint8(255.0 * v.z)
	return c
}

func Distance(v1, v2 *Vec3) float {
	d := Sub(v1, v2)
	return d.Length()
}

func Sum(vs ...*Vec3) *Vec3 {
	v := &Vec3{}
	for _, vi := range vs {
		v.x += vi.x
		v.y += vi.y
		v.z += vi.z
	}
	return v
}

// norm
func (v *Vec3) Length() float {
	return math.Sqrt(v.Dot(v))
}

func (v *Vec3) Norm() *Vec3 {
	return DivNum(v, v.Length())
}

func (v *Vec3) Normalize() *Vec3 {
	v.DivNum(v.Length())
	return v
}

func (v *Vec3) Reflect(n *Vec3) *Vec3 {
	r := Sub(v, MulNum(n, 2*Dot(v, n)/n.Length()))
	return r
}

func (v *Vec3) Refract(n *Vec3, ratio_n float) *Vec3 {
	uv := v.Norm()
	dt := Dot(uv, n)
	discrinant := 1.0 - ratio_n*ratio_n*(1-dt*dt)
	if discrinant > 0 {
		refracted := Sub(uv, MulNum(n, dt))
		refracted.MulNum(ratio_n)
		refracted.Sub(MulNum(n, math.Sqrt(discrinant)))
		return refracted
	} else {
		return nil
	}
}

func (v *Vec3) Copy() *Vec3 {
	r := &Vec3{}
	*r = *v
	return r
}

func (v *Vec3) Clear() {
	v.x = 0
	v.y = 0
	v.z = 0
}

func (v *Vec3) isZero() bool {
	return v.x == 0 && v.y == 0 && v.z == 0
}

// arithmetic
func (v *Vec3) Add(rhs *Vec3) *Vec3 {
	v.x += rhs.x
	v.y += rhs.y
	v.z += rhs.z
	return v
}

func (v *Vec3) AddNum(n float) *Vec3 {
	v.x += n
	v.y += n
	v.z += n
	return v
}

func (v *Vec3) Sub(rhs *Vec3) *Vec3 {
	v.x -= rhs.x
	v.y -= rhs.y
	v.z -= rhs.z
	return v
}

func (v *Vec3) SubNum(n float) *Vec3 {
	v.x -= n
	v.y -= n
	v.z -= n
	return v
}

func (v *Vec3) Mul(rhs *Vec3) *Vec3 {
	v.x *= rhs.x
	v.y *= rhs.y
	v.z *= rhs.z
	return v
}

func (v *Vec3) MulNum(n float) *Vec3 {
	v.x *= n
	v.y *= n
	v.z *= n
	return v
}

func (v *Vec3) Div(rhs *Vec3) *Vec3 {
	v.x /= rhs.x
	v.y /= rhs.y
	v.z /= rhs.z
	return v
}

func (v *Vec3) DivNum(n float) *Vec3 {
	v.x /= n
	v.y /= n
	v.z /= n
	return v
}

func (v *Vec3) Dot(rhs *Vec3) float {
	return Dot(v, rhs)
}

func (v *Vec3) Cross(rhs *Vec3) *Vec3 {
	return v.Cross(rhs)
}

// non instance method
func Dot(v1, v2 *Vec3) float {
	return v1.x*v2.x + v1.y*v2.y + v1.z*v2.z
}

func Cross(v1, v2 *Vec3) *Vec3 {
	v := &Vec3{}
	v.x = v1.y*v2.z - v1.z*v2.y
	v.y = v1.z*v2.x - v1.x*v2.z
	v.z = v1.x*v2.y - v1.y*v2.x
	return v
}

func Add(v1, v2 *Vec3) *Vec3 {
	v := &Vec3{}
	v.x = v1.x + v2.x
	v.y = v1.y + v2.y
	v.z = v1.z + v2.z
	return v
}

func Sub(v1, v2 *Vec3) *Vec3 {
	v := &Vec3{}
	v.x = v1.x - v2.x
	v.y = v1.y - v2.y
	v.z = v1.z - v2.z
	return v
}

func Mul(v1, v2 *Vec3) *Vec3 {
	v := &Vec3{}
	v.x = v1.x * v2.x
	v.y = v1.y * v2.y
	v.z = v1.z * v2.z
	return v
}

func Div(v1, v2 *Vec3) *Vec3 {
	v := &Vec3{}
	v.x = v1.x / v2.x
	v.y = v1.y / v2.y
	v.z = v1.z / v2.z
	return v
}

func AddNum(v *Vec3, n float) *Vec3 {
	r := &Vec3{}
	r.x = v.x + n
	r.y = v.y + n
	r.z = v.z + n
	return r
}

func SubNum(v *Vec3, n float) *Vec3 {
	r := &Vec3{}
	r.x = v.x - n
	r.y = v.y - n
	r.z = v.z - n
	return r
}

func MulNum(v *Vec3, n float) *Vec3 {
	r := &Vec3{}
	r.x = v.x * n
	r.y = v.y * n
	r.z = v.z * n
	return r
}

func DivNum(v *Vec3, n float) *Vec3 {
	r := &Vec3{}
	r.x = v.x / n
	r.y = v.y / n
	r.z = v.z / n
	return r
}
