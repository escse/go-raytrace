package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
)

type Camera struct {
	nx         int
	ny         int
	ns         int
	lowerLeft  *Vec3
	horizontal *Vec3
	vertical   *Vec3
	origin     *Vec3
}

func (camera *Camera) Move(v *Vec3) {
	camera.lowerLeft.Add(v)
	camera.origin.Add(v)
}

func (camera *Camera) Draw(f *os.File, world *Objects) {
	vc := &Vec3{}
	for j := camera.ny - 1; j >= 0; j-- {
		for i := 0; i < camera.nx; i++ {
			vc.Clear()
			for s := 0; s < camera.ns; s++ {
				u := (float(i) + rand.Float64()) / float(camera.nx-1)
				v := (float(j) + rand.Float64()) / float(camera.ny-1)
				ray := camera.Get_ray(u, v)
				c := HitableVec3(ray, world, 30)
				vc.Add(c)
			}
			vc.DivNum(float(camera.ns))
			color := vc.ToColor()
			fmt.Fprint(f, fmt.Sprintf("%d %d %d\n", color.r, color.g, color.b))
		}
	}
}

func (camera *Camera) DrawFuzz(f *os.File, world *Objects) {
	var color *Color
	for j := camera.ny - 1; j >= 0; j-- {
		for i := 0; i < camera.nx; i++ {
			u := (float(i)) / float(camera.nx-1)
			v := (float(j)) / float(camera.ny-1)
			ray := camera.Get_ray(u, v)
			c := HitableVec3(ray, world, 50)
			if !c.isZero() {
				color = c.ToColor()
			}
			fmt.Fprint(f, fmt.Sprintf("%d %d %d\n", color.r, color.g, color.b))
		}
	}
}

func NewCamera(lookfrom, lookat, vup *Vec3, vfov float, nx, ny, ns int) *Camera {
	aspect := float(nx) / float(ny)
	theta := vfov * math.Pi / 180.0
	half_height := math.Tan(theta / 2.0)
	half_width := aspect * half_height

	c := &Camera{nx: nx, ny: ny, ns: ns, origin: lookfrom}

	w := Sub(lookfrom, lookat).Norm()
	u := Cross(vup, w).Norm()
	v := Cross(w, u)
	c.lowerLeft = Sum(c.origin, MulNum(u, -half_width), MulNum(v, -half_height), MulNum(w, -1.0))
	c.horizontal = MulNum(u, 2.0*half_width)
	c.vertical = MulNum(v, 2.0*half_height)
	return c
}

func (c *Camera) Get_ray(u, v float) *Ray {
	return &Ray{c.origin, Sum(c.lowerLeft, MulNum(c.horizontal, u), MulNum(c.vertical, v), MulNum(c.origin, -1.0))}
}
