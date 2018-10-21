package main

import (
	"math/rand"
	"time"
)

type Objects struct {
	list []Hitable
}

func (h *Objects) Add(hs ...Hitable) {
	h.list = append(h.list, hs...)
}

func (h *Objects) hit(r *Ray, tmin, tmax float, record *HitRecord) bool {
	tmpRecord := &HitRecord{}
	hit_anything := false
	closest := tmax
	for i := 0; i < len(h.list); i++ {
		if h.list[i].hit(r, tmin, closest, tmpRecord) {
			hit_anything = true
			closest = tmpRecord.t
			*record = *tmpRecord
		}
	}
	return hit_anything
}

func NewWorld() *Objects {
	world := &Objects{}
	objects := []Hitable{
		&Sphere{&Vec3{0.0, 0.0, -1.0}, 0.5, &Lambertian{&Vec3{0.1, 0.2, 0.5}}},
		&Sphere{&Vec3{0.0, -100.5, -1.0}, 100.0, &Lambertian{&Vec3{0.8, 0.8, 0.0}}},
		&Sphere{&Vec3{1.0, 0.0, -1.0}, 0.5, &Metal{&Vec3{0.8, 0.6, 0.2}, 0.0}},
		&Sphere{&Vec3{-1.0, 0.0, -1.0}, 0.5, &Dielectric{1.5}},
	}
	world.Add(objects...)
	return world
}

func NewRandomWorld() *Objects {
	rand.Seed(time.Now().UTC().UnixNano())
	world := &Objects{}
	n := 500
	objects := make([]Hitable, n)
	objects[0] = &Sphere{&Vec3{0.0, -1000.0, 0.0}, 1000.0, &Lambertian{&Vec3{0.5, 0.5, 0.5}}}
	objects[1] = &Sphere{&Vec3{0.0, 1.0, 0.0}, 1.0, &Dielectric{1.5}}
	objects[2] = &Sphere{&Vec3{-4.0, 1.0, 0.0}, 1.0, &Lambertian{&Vec3{0.4, 0.2, 0.1}}}
	objects[3] = &Sphere{&Vec3{4.0, 1.0, 0.0}, 1.0, &Metal{&Vec3{0.7, 0.6, 0.5}, 0.0}}
	i := 4
	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			choose_mat := rand.Float64()
			center := &Vec3{float(a) + 0.9*choose_mat, 0.2, float(b) + 0.9*choose_mat}
			if Sub(center, &Vec3{4.0, 0.2, 0.0}).Length() > 0.9 {
				if choose_mat < 0.8 {
					objects[i] = &Sphere{center, 0.2,
						&Lambertian{&Vec3{rand.Float64() * rand.Float64(), rand.Float64() * rand.Float64(), rand.Float64() * rand.Float64()}}}
				} else if choose_mat < 0.95 {
					// r1, r2, r3 := rand.Float64(), rand.Float64(), rand.Float64()
					objects[i] = &Sphere{center, 0.2,
						&Metal{&Vec3{0.5 * (1.0 + rand.Float64()), 0.5 * (1.0 + rand.Float64()), 0.5 * (1.0 + rand.Float64())}, 0.5 * (1.0 + rand.Float64())}}
				} else {
					objects[i] = &Sphere{center, 0.2, &Dielectric{1.5}}
				}
				i += 1
			}
		}
	}
	world.Add(objects[:i]...)
	return world
}
