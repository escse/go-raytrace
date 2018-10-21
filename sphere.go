package main

import (
	"math"
	"math/rand"
)

type Sphere struct {
	center   *Vec3
	radius   float
	material Scatterable
}

func (s *Sphere) Distance(r *Ray) float {
	oc := Sub(r.Origin, s.center)
	b := math.Abs(Dot(r.Direction.Norm(), oc))
	c := oc.Length()
	d := math.Sqrt(c*c - b*b)
	return d
}

func (s *Sphere) hitRay(r *Ray) bool {
	return s.hit(r, 0.0, 0.0, nil)
}

func (s *Sphere) hit(r *Ray, tmin, tmax float, record *HitRecord) bool {
	oc := Sub(r.Origin, s.center)
	a := Dot(r.Direction, r.Direction)
	b := Dot(oc, r.Direction)
	c := Dot(oc, oc) - s.radius*s.radius
	discriminant := b*b - a*c
	if record == nil {
		return discriminant > 0
	}
	if discriminant > 0 {
		var temp float
		delta := math.Sqrt(discriminant)
		temp = (-b - delta) / a
		if temp < tmax && temp > tmin {
			record.t = temp
			record.p = r.PointAt(temp)
			record.normal = Sub(record.p, s.center).DivNum(s.radius)
			record.material = s.material
			return true
		}
		temp = (-b + delta) / a
		if temp < tmax && temp > tmin {
			record.t = temp
			record.p = r.PointAt(temp)
			record.normal = Sub(record.p, s.center).DivNum(s.radius)
			record.material = s.material
			return true
		}
	}
	return false
}

func NewRandomSpherePoint() *Vec3 {
	v := Vec3{}
	for {
		v.x, v.y, v.z = rand.Float64(), rand.Float64(), rand.Float64()
		v.MulNum(2.0).SubNum(1.0)
		if v.Length() < 1.0 {
			break
		}
	}
	return &v
}
