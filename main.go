package main

import (
	"fmt"
	"os"
)

type float = float64

func main() {
	// nx, ny, ns := 1280, 720, 10
	nx, ny, ns := 240, 180, 200
	// nx, ny, ns := 160, 120, 10

	world := NewRandomWorld()
	lookfrom := &Vec3{13.0, 2.0, 3.0}
	lookat := &Vec3{0.0, 0.0, -1.0}
	camera := NewCamera(lookfrom, lookat, &Vec3{0.0, 1.0, 0.0}, 20.0, nx, ny, ns)
	move := Sub(lookfrom, lookat).Norm()
	for i := 0; i < 240; i++ {
		print(i)
		f, _ := os.Create(fmt.Sprintf("image/img%d.ppm", i))
		fmt.Fprint(f, fmt.Sprintf("P3\n%d %d\n255\n", nx, ny))
		camera.Draw(f, world)
		camera.Move(MulNum(move, 0.04))
		f.Close()
	}
}
