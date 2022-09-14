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
func NewEntity(name string) *Entity {
	return &Entity{
		name: Name(name),
	}
}

// ID Getter
func (g *Entity) ID() ID {
	return g.id
}

// SetID Setter
func (g *Entity) SetID(id uint) {
	g.id = ID(id)
}

// Name Getter
func (g *Entity) Name() Name {
	return g.name
}

// Entities ...
type Entities []Entity
