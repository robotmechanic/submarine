package navmesh

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/ungerik/go3d/float64/vec2"
	"testing"
)

func TestMesh(t *testing.T) {
	Convey("mesh", t, func() {
		m, _ := LoadMeshFromJSONFile("fixtures/mesh.json")

		Convey(".LoadMeshFromJSONFile", func() {
			Convey("with a valid json path", func() {
				Convey("should creates a mesh", func() {
					So(m.Vertices[0], ShouldResemble, &vec2.T{0, 0})
					So(m.Triangles[0], ShouldResemble, newTriangle(
						m.Vertices[0],
						m.Vertices[1],
						m.Vertices[2],
					))
				})
			})

			Convey("should initialize outerEdges", func() {
				So(m.outerEdges, ShouldHaveLength, 8)
				So(m.outerEdges, ShouldContain, Edge{m.Vertices[0], m.Vertices[1]})
				So(m.outerEdges, ShouldContain, Edge{m.Vertices[1], m.Vertices[2]})
				So(m.outerEdges, ShouldContain, Edge{m.Vertices[2], m.Vertices[3]})
				So(m.outerEdges, ShouldContain, Edge{m.Vertices[3], m.Vertices[4]})
				So(m.outerEdges, ShouldContain, Edge{m.Vertices[0], m.Vertices[5]})
				So(m.outerEdges, ShouldContain, Edge{m.Vertices[4], m.Vertices[7]})
				So(m.outerEdges, ShouldContain, Edge{m.Vertices[5], m.Vertices[6]})
				So(m.outerEdges, ShouldContain, Edge{m.Vertices[6], m.Vertices[7]})
			})

			Convey("should initialize trianglesByVertex", func() {
				triangles := m.trianglesByVertex[m.Vertices[0]]
				So(triangles, ShouldHaveLength, 4)
				So(triangles, ShouldContain, m.Triangles[0])
				So(triangles, ShouldContain, m.Triangles[1])
				So(triangles, ShouldContain, m.Triangles[2])
				So(triangles, ShouldContain, m.Triangles[3])
			})

			Convey("should initialize adjoiningVertices", func() {
				vertices := m.adjoiningVertices[m.Vertices[0]]
				So(vertices, ShouldHaveLength, 5)
				So(vertices, ShouldContain, m.Vertices[1])
				So(vertices, ShouldContain, m.Vertices[2])
				So(vertices, ShouldContain, m.Vertices[3])
				So(vertices, ShouldContain, m.Vertices[4])
				So(vertices, ShouldContain, m.Vertices[5])
			})
		})

		Convey("#findTriangleByPoint", func() {
			Convey("should find the triangle that contains the specified point", func() {
				triangle := m.findTriangleByPoint(&vec2.T{1, -11})
				So(triangle, ShouldEqual, m.Triangles[5])
			})
		})

		Convey("#isIntersectWithLine", func() {
			Convey("with intersected points", func() {
				Convey("should return true", func() {
					p1 := &vec2.T{1, 0}
					p2 := &vec2.T{1, 10}
					So(m.isIntersectedWithLine(p1, vec2.Sub(p2, p1)), ShouldBeTrue)
				})
			})

			Convey("with no-intersected points", func() {
				Convey("should return false", func() {
					p1 := &vec2.T{999, 999}
					p2 := &vec2.T{1000, 1000}
					So(m.isIntersectedWithLine(p1, vec2.Sub(p2, p1)), ShouldBeFalse)
				})
			})
		})

		Convey("#intersectWithLine", func() {
			Convey("with intersected points", func() {
				Convey("should return the intersection point", func() {
					p1 := &vec2.T{1, 0}
					p2 := &vec2.T{10, 0}
					result := m.intersectWithLine(p1, vec2.Sub(p2, p1))
					So(result[0], ShouldEqual, 8.875)
					So(result[1], ShouldEqual, 0)
				})
			})

			Convey("with intersected points which has multi intersection points", func() {
				Convey("should return the most neaby intersection point", func() {
					p1 := &vec2.T{-7, 3.5}
					p2 := &vec2.T{-1, -11}
					result := m.intersectWithLine(p1, vec2.Sub(p2, p1))
					So(result[0], ShouldEqual, -5.551724137931035)
					So(result[1], ShouldEqual, 0)
				})
			})

			Convey("with no-intersected points", func() {
				Convey("should return nil", func() {
					p1 := &vec2.T{0.00000000001, 0}
					p2 := &vec2.T{5, 0}
					So(m.intersectWithLine(p1, vec2.Sub(p2, p1)), ShouldBeNil)

					p1 = &vec2.T{3, 2}
					p2 = &vec2.T{2, -10}
					So(m.intersectWithLine(p1, vec2.Sub(p2, p1)), ShouldBeNil)
				})
			})
		})
	})
}
