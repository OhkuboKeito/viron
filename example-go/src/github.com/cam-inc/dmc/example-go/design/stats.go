package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// StatsPlanetMediaType of media type.
var StatsPlanetMediaType = MediaType("application/vnd.statsplanet+json", func() {
	Description("Planets data")
	ContentType("application/json")

	Attributes(func() {
		Attribute("keys", ArrayOf(String), "key names of graph data")
		Attribute("data", ArrayOf(ArrayOf(Any)), "graph data")
		Required("keys", "data")
	})

	View("default", func() {
		Attribute("keys")
		Attribute("data")
	})
})

var _ = Resource("stats_dau", func() {
	Origin(OriginURL, OriginAllowAll)
	BasePath("/stats/dau")
	Action("show", func() {
		Routing(GET(""))
		Description("Service Daily Activity User")
		Response(OK, Number)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("stats_mau", func() {
	Origin(OriginURL, OriginAllowAll)
	BasePath("/stats/mau")
	Action("show", func() {
		Routing(GET(""))
		Description("Service Monthly Activity User")
		Response(OK, Number)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("stats_planet", func() {
	Origin(OriginURL, OriginAllowAll)
	BasePath("/stats/planet")
	DefaultMedia(StatsPlanetMediaType)
	Action("bar", func() {
		Routing(GET("/bar"))
		Description("Planets Information")
		Response(OK, func() {
			Media(StatsPlanetMediaType)
		})
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
	Action("scatterplot", func() {
		Routing(GET("/scatterplot"))
		Description("Planets Information")
		Response(OK, func() {
			Media(StatsPlanetMediaType)
		})
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
	Action("line", func() {
		Routing(GET("/line"))
		Description("Planets Information")
		Response(OK, func() {
			Media(StatsPlanetMediaType)
		})
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
	Action("horizontal-bar", func() {
		Routing(GET("/horizontal-bar"))
		Description("Planets Information")
		Response(OK, func() {
			Media(StatsPlanetMediaType)
		})
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
	Action("stacked-bar", func() {
		Routing(GET("/stacked-bar"))
		Description("Planets Information")
		Response(OK, func() {
			Media(StatsPlanetMediaType)
		})
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
	Action("horizontal-stacked-bar", func() {
		Routing(GET("/horizontal-stacked-bar"))
		Description("Planets Information")
		Response(OK, func() {
			Media(StatsPlanetMediaType)
		})
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
	Action("stacked-area", func() {
		Routing(GET("/stacked-area"))
		Description("Planets Information")
		Response(OK, func() {
			Media(StatsPlanetMediaType)
		})
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})
