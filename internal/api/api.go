package api

import (
	"RainbowRunner/internal/api/types"
	"RainbowRunner/internal/objects"
	"encoding/json"
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type query struct{}

func (_ *query) GetZones() *types.ZoneCollection {
	//zone := &objects.Zone{
	//	Name: "townston",
	//}
	//
	//zones := types.NewZoneCollection([]*types.Zone{types.NewZone(zone)})

	list := make([]*types.Zone, 0)

	for _, zone := range objects.Zones.GetZones() {
		list = append(list, types.NewZone(zone))
	}

	return types.NewZoneCollection(list)
}

func (_ *query) GetEntities() *types.EntityCollection {
	//avatar := objects.NewAvatar("avatar")
	//
	//avatar.RREntityProperties().ID = 12
	//avatar.RREntityProperties().OwnerID = 1
	//avatar.RREntityProperties().Zone = &objects.Zone{
	//	Name: "townston",
	//}
	//
	//return types.NewEntityCollection([]*types.Entity{types.NewEntity(avatar)})

	list := make([]*types.Entity, 0)

	for _, entity := range objects.Entities.GetEntities() {
		list = append(list, types.NewEntity(entity))
	}

	return types.NewEntityCollection(list)
}

func (_ *query) GetPlayers() *types.PlayerCollection {
	//player := objects.NewPlayer("Ellie")
	//
	//player.RREntityProperties().ID = 12
	//
	//zone := &objects.Zone{
	//	Name: "townston",
	//}
	//
	//player.RREntityProperties().Zone = zone
	//
	//avatar := objects.NewAvatar("avatar")
	//
	//component := objects.NewUnitBehavior("unitbehavior")
	//avatar.AddChild(component)
	//
	//player.AddChild(avatar)
	//
	//rrPlayer := &objects.RRPlayer{
	//	Conn: &connections.RRConn{
	//		Client: &connections.RRConnClient{
	//			ID: 1,
	//		},
	//	},
	//	CurrentCharacter: player,
	//	Zone:             zone,
	//}
	//
	//return types.NewPlayerCollection([]*types.Player{types.NewPlayer(rrPlayer)})
	list := make([]*types.Player, 0)

	for _, p := range objects.Players.GetPlayers() {
		list = append(list, types.NewPlayer(p))
	}

	return types.NewPlayerCollection(list)
}

var schema = `
type Query {
	getZones: ZoneCollection
	getEntities: EntityCollection
	getPlayers: PlayerCollection
}

type EntityCollection {
	entities: [Entity]
}

type Entity {
	id: Int
	ownerId: Int
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

type ZoneCollection {
	zones: [Zone]
}

type Zone {
	name: String
	entities: [Entity]
	players: [Player]
}
`

type MyHandler struct {
	relay.Handler
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		return
	}

	var params struct {
		Query         string                 `json:"query"`
		OperationName string                 `json:"operationName"`
		Variables     map[string]interface{} `json:"variables"`
	}

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := h.Schema.Exec(r.Context(), params.Query, params.OperationName, params.Variables)
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

func NewMyHandler(s *graphql.Schema) *MyHandler {
	h := &MyHandler{}

	h.Schema = s

	return h
}

func StartGraphqlAPI() {
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schema := graphql.MustParseSchema(schema, &query{}, opts...)
	http.Handle("/query", NewMyHandler(schema))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
