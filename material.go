package main

import (
	"math"
)

type Scatterable interface {
	scatter(*Ray, *HitRecord, *Vec3, *Ray) bool
}

type Material struct {
	albedo *Vec3
}

type Metal struct {
	albedo *Vec3
	fuzz   float
}

func (m *Metal) scatter(rin *Ray, record *HitRecord, attenuation *Vec3, scattered *Ray) bool {
	*scattered = *rin.Reflect(record)
	if m.fuzz > 0 {
		scattered.Direction.Add(MulNum(NewRandomSpherePoint(), math.Min(1.0, m.fuzz)))
	}
	*attenuation = *m.albedo
	return Dot(scattered.Direction, record.normal) > 0
}

type Lambertian Material

func (m *Lambertian) scatter(rin *Ray, record *HitRecord, attenuation *Vec3, scattered *Ray) bool {
	target := Sum(record.p, record.normal, NewRandomSpherePoint())
	*scattered = Ray{record.p, Sub(target, record.p)}
	*attenuation = *m.albedo
	return true
}

type Dielectric struct {
	ref_idx float
}

func (m *Dielectric) scatter(rin *Ray, record *HitRecord, attenuation *Vec3, scattered *Ray) bool {
	*attenuation = Vec3{1.0, 1.0, 1.0}
	r := rin.Refract(record, m.ref_idx)
	if r != nil {
		*scattered = *r
		return true
	}
	return false
}
