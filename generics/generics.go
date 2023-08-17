package generics

// https://go.dev/play/p/UvPmiBwAwhh
import "context"

type DB struct{}

type Coll[T Model] struct{}

type Model interface {
	GetID() any
	SetID(any) error
	TableName() string
	// If the model uses single table inheritance, this should return the column name and type value
	// for the STI column. For example, if the column name is "type" and the type value is "incident",
	// this should return ("type", "incident").
	//
	// If the model does not use STI, this should return ("", "").
	//
	// In general, you should avoid using STI. It's a legacy feature.
	STI_Type() (string, string)
}

type DefaultModel struct {
	ID int `bson:"_id" json:"_id"`
}

func (m DefaultModel) GetID() any { return 1 }
func (m *DefaultModel) SetID(id any) error {
	m.ID = id.(int)
	return nil
}
func (m DefaultModel) TableName() string          { return "hello" }
func (m DefaultModel) STI_Type() (string, string) { return "", "" }

func Collection[T Model](db *DB) *Coll[T] {
	return &Coll[T]{}
}

func (c *Coll[T]) FindID(ctx context.Context, id int) (T, error) {
	var t T
	return t, nil
}

func (c *Coll[T]) Find(ctx context.Context, q Query[T]) (T, error) {
	var t T
	// do something with the payload
	_ = q.Resolve()
	return t, nil
}

type Query[T any] struct {
	payload any
}

func (q Query[T]) Resolve() any {
	return q.payload
}

type Incident struct {
	DefaultModel
	Name      string
	ChannelID string
}

var IncidentsInState = func(state string) Query[*Incident] {
	return Query[*Incident]{payload: state}
}

func ALL[T any]() Query[T] { return Query[T]{} }

func MyTest(ctx context.Context, db *DB) {
	coll := Collection[*Incident](db)
	coll.FindID(ctx, 1)
	coll.Find(ctx, IncidentsInState("hello"))
	coll.Find(ctx, ALL[*Incident]())
}
