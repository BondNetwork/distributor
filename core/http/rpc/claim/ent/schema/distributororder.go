package schema

import (
	"database/sql"
	"database/sql/driver"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"fmt"
	"math/big"
)

// DistributorOrder holds the schema definition for the DistributorOrder entity.
type DistributorOrder struct {
	ent.Schema
}

// Fields of the DistributorOrder.
func (DistributorOrder) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("orderId"),
		field.String("batchId"),
		field.String("projectId"),
		field.String("account"),
		field.Int("amount").
			GoType(new(BigInt)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(78, 0)",
			}),
		field.Int("index"),
		field.Time("createAt"),
		field.String("pTaskId"),
		field.Strings("proof"),
	}
}

// Edges of the DistributorOrder.
func (DistributorOrder) Edges() []ent.Edge {
	return nil
}

type BigInt struct {
	big.Int
}

func (b *BigInt) Scan(src any) error {
	var i sql.NullString
	if err := i.Scan(src); err != nil {
		return err
	}
	if !i.Valid {
		return nil
	}
	if _, ok := b.Int.SetString(i.String, 10); ok {
		return nil
	}
	return fmt.Errorf("could not scan type %T with value %v into BigInt", src, src)
}

func (b *BigInt) Value() (driver.Value, error) {
	return b.String(), nil
}
