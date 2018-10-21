package main

type Ray struct {
	Origin    *Vec3
	Direction *Vec3
}

func (ray *Ray) PointAt(t float) *Vec3 {
	return Add(ray.Origin, MulNum(ray.Direction, t))
}

func (ray *Ray) Reflect(record *HitRecord) *Ray {
	reflected := ray.Direction.Reflect(record.normal)
	r := &Ray{record.p, reflected}
	return r
}

func (ray *Ray) Refract(record *HitRecord, ratio_n float) *Ray {
	normal := record.normal.Copy()
	if Dot(normal, ray.Direction) > 0 {
		normal.MulNum(-1.0)
	} else {
		ratio_n = 1.0 / ratio_n
	}
	refracted := ray.Direction.Refract(normal, ratio_n)
	if refracted == nil {
		return nil
	}
	r := &Ray{record.p, refracted}
	return r
}

// func (ray *Ray) ColorLine() *Color {
// 	unitDirection := ray.Direction.Norm()
// 	t := 0.5 * (unitDirection.y + 1.0)
// 	v := Add(MulNum(n1, 1.0-t), MulNum(n2, t))
// 	return v.ToColor()
// }

// func (ray *Ray) ColorSphere() *Color {
// 	r := &Color{255, 0, 0}
// 	if s.hitRay(ray) {
// 		return r
// 	}
// 	unit_direction := ray.Direction.Norm()
// 	t := 0.5 * (unit_direction.y + 1.0)
// 	v := Add(MulNum(n1, 1.0-t), MulNum(n2, t))
// 	return v.ToColor()
// }
