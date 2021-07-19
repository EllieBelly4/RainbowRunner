package api

import (
	"RainbowRunner/internal/api/types"
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/objects"
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type query struct{}

func (_ *query) GetEntities() *types.EntityCollection {
	avatar := objects.NewAvatar("avatar")

	avatar.RREntityProperties().ID = 12
	avatar.RREntityProperties().OwnerID = 1
	avatar.RREntityProperties().Zone = &objects.Zone{
		Name: "townston",
	}

	return types.NewEntityCollection([]*types.Entity{types.NewEntity(avatar)})
}

func (_ *query) GetPlayers() *types.PlayerCollection {
	player := objects.NewPlayer("Ellie")

	player.RREntityProperties().ID = 12

	zone := &objects.Zone{
		Name: "townston",
	}

	player.RREntityProperties().Zone = zone

	avatar := objects.NewAvatar("avatar")

	component := objects.NewUnitBehavior("unitbehavior")
	avatar.AddChild(component)

	player.AddChild(avatar)

	rrPlayer := &objects.RRPlayer{
		Conn: &connections.RRConn{
			Client: &connections.RRConnClient{
				ID: 1,
			},
		},
		CurrentCharacter: player,
		Zone:             zone,
	}

	return types.NewPlayerCollection([]*types.Player{types.NewPlayer(rrPlayer)})
}

var schema = `
type Query {
	getEntities: EntityCollection
	getPlayers: PlayerCollection
}

type EntityCollection {
	entities: [Entity]
}

type Entity {
	id: Int
	typeName: String
	zone: Zone
	children: [Entity]
}

type Component {
	id: Int
	typeName: String
}

type PlayerCollection {
	players: [Player]
}

type Player {
	id: Int
	name: String
	zone: Zone
	currentCharacter: Entity
}

type Zone {
	name: String
}
`

func StartGraphqlAPI() {
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schema := graphql.MustParseSchema(schema, &query{}, opts...)
	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
