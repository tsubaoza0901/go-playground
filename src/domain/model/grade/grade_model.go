package grade

type (
	// ID ...
	ID uint
	// Name ...
	Name string
)

// NonGrade グレードなし（初期グレード）
const NonGrade uint = 6

// Entity ...
type Entity struct {
	id   ID
	name Name
}

// NewEntity ...
func NewEntity(id uint) *Entity {
	return &Entity{
		id: ID(id),
	}
}

// MakeEntity ...
func MakeEntity(id ID, name Name) *Entity {
	return &Entity{
		id:   ID(id),
		name: Name(name),
	}
}

// ID Getter
func (g *Entity) ID() ID {
	return g.id
}

func (g *Entity) setID(id ID) {
	g.id = id
}

// Name Getter
func (g *Entity) Name() Name {
	return g.name
}

// Entities ...
type Entities []Entity
