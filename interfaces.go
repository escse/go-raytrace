package main

import (
	"math"
)

type HitRecord struct {
	t        float
	p        *Vec3
	normal   *Vec3
	material Scatterable
}

type Hitable interface {
	hit(*Ray, float, float, *HitRecord) bool
}

func HitableVec3(r *Ray, world Hitable, depth int) *Vec3 {
	var v *Vec3
	record := &HitRecord{}
	if world.hit(r, 0.00001, math.MaxFloat64, record) {
		scattered := &Ray{}
		attenuation := &Vec3{}
		if depth > 0 && record.material.scatter(r, record, attenuation, scattered) {
			return Mul(attenuation, HitableVec3(scattered, world, depth-1))
		}
		return &Vec3{0.1, 0.1, 0.1}
	} else {
		var n1 = &Vec3{1.0, 1.0, 1.0}
		var n2 = &Vec3{0.5, 0.7, 1.0}
		unitDirection := r.Direction.Norm()
		t := 0.5 * (unitDirection.y + 1.0)
		v = Add(MulNum(n1, 1.0-t), MulNum(n2, t))
	}
	return v
}
