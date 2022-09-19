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

// ID Getter
func (g *Entity) ID() ID {
	return g.id
}

// Name Getter
func (g *Entity) Name() string {
	return string(g.name)
}

// InitEntity 初期化関数。原則、新規レコード登録用。
func InitEntity(id uint) *Entity {
	return &Entity{
		id: ID(id),
	}
}

// NewEntity 新規インスタンス生成関数（初期値の設定なし）
func NewEntity(id ID, name Name) *Entity {
	return &Entity{
		id:   ID(id),
		name: Name(name),
	}
}

// Entities ...
type Entities []Entity
