package navmesh

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/ungerik/go3d/float64/vec2"
	"testing"
)

func TestTriangle(t *testing.T) {
	Convey("Triangle", t, func() {
		v1 := &vec2.T{5, 5}
		v2 := &vec2.T{-5, 5}
		v3 := &vec2.T{0, -5}
		triangle := newTriangle(v1, v2, v3)

		Convey("#containsPoint", func() {
			Convey("with a point in the triangle", func() {
				So(triangle.containsPoint(&vec2.T{0, 0}), ShouldBeTrue)
			})

			Convey("with a point out the triangle", func() {
				So(triangle.containsPoint(&vec2.T{0, 6}), ShouldBeFalse)
			})

			Convey("with a point on the triangle edge", func() {
				So(triangle.containsPoint(&vec2.T{-1, 5}), ShouldBeTrue)
			})

			Convey("with a point on the triangle vertex", func() {
				So(triangle.containsPoint(v1), ShouldBeTrue)
				So(triangle.containsPoint(v2), ShouldBeTrue)
				So(triangle.containsPoint(v3), ShouldBeTrue)
			})
		})

		Convey("#hasVertex", func() {
			Convey("with a vertex on the triangle", func() {
				So(triangle.hasVertex(triangle.Vertices[0]), ShouldBeTrue)
			})

			Convey("with a vertex not on the triangle", func() {
				So(triangle.hasVertex(&vec2.T{-1, -1}), ShouldBeFalse)
			})
		})
	})
}