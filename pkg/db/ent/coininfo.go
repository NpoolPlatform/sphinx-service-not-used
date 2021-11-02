// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"sphinx/ent/coininfo"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// CoinInfo is the model entity for the CoinInfo schema.
type CoinInfo struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Unit holds the value of the "unit" field.
	Unit string `json:"unit,omitempty"`
	// NeedSigninfo holds the value of the "need_signinfo" field.
	NeedSigninfo bool `json:"need_signinfo,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CoinInfoQuery when eager-loading is set.
	Edges CoinInfoEdges `json:"edges"`
}

// CoinInfoEdges holds the relations/edges for other nodes in the graph.
type CoinInfoEdges struct {
	// Keys holds the value of the keys edge.
	Keys []*KeyStore `json:"keys,omitempty"`
	// Transactions holds the value of the transactions edge.
	Transactions []*Transaction `json:"transactions,omitempty"`
	// Reviews holds the value of the reviews edge.
	Reviews []*Review `json:"reviews,omitempty"`
	// WalletNodes holds the value of the wallet_nodes edge.
	WalletNodes []*WalletNode `json:"wallet_nodes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// KeysOrErr returns the Keys value or an error if the edge
// was not loaded in eager-loading.
func (e CoinInfoEdges) KeysOrErr() ([]*KeyStore, error) {
	if e.loadedTypes[0] {
		return e.Keys, nil
	}
	return nil, &NotLoadedError{edge: "keys"}
}

// TransactionsOrErr returns the Transactions value or an error if the edge
// was not loaded in eager-loading.
func (e CoinInfoEdges) TransactionsOrErr() ([]*Transaction, error) {
	if e.loadedTypes[1] {
		return e.Transactions, nil
	}
	return nil, &NotLoadedError{edge: "transactions"}
}

// ReviewsOrErr returns the Reviews value or an error if the edge
// was not loaded in eager-loading.
func (e CoinInfoEdges) ReviewsOrErr() ([]*Review, error) {
	if e.loadedTypes[2] {
		return e.Reviews, nil
	}
	return nil, &NotLoadedError{edge: "reviews"}
}

// WalletNodesOrErr returns the WalletNodes value or an error if the edge
// was not loaded in eager-loading.
func (e CoinInfoEdges) WalletNodesOrErr() ([]*WalletNode, error) {
	if e.loadedTypes[3] {
		return e.WalletNodes, nil
	}
	return nil, &NotLoadedError{edge: "wallet_nodes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CoinInfo) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case coininfo.FieldNeedSigninfo:
			values[i] = new(sql.NullBool)
		case coininfo.FieldID:
			values[i] = new(sql.NullInt64)
		case coininfo.FieldName, coininfo.FieldUnit:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CoinInfo", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CoinInfo fields.
func (ci *CoinInfo) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case coininfo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ci.ID = int(value.Int64)
		case coininfo.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ci.Name = value.String
			}
		case coininfo.FieldUnit:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field unit", values[i])
			} else if value.Valid {
				ci.Unit = value.String
			}
		case coininfo.FieldNeedSigninfo:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field need_signinfo", values[i])
			} else if value.Valid {
				ci.NeedSigninfo = value.Bool
			}
		}
	}
	return nil
}

// QueryKeys queries the "keys" edge of the CoinInfo entity.
func (ci *CoinInfo) QueryKeys() *KeyStoreQuery {
	return (&CoinInfoClient{config: ci.config}).QueryKeys(ci)
}

// QueryTransactions queries the "transactions" edge of the CoinInfo entity.
func (ci *CoinInfo) QueryTransactions() *TransactionQuery {
	return (&CoinInfoClient{config: ci.config}).QueryTransactions(ci)
}

// QueryReviews queries the "reviews" edge of the CoinInfo entity.
func (ci *CoinInfo) QueryReviews() *ReviewQuery {
	return (&CoinInfoClient{config: ci.config}).QueryReviews(ci)
}

// QueryWalletNodes queries the "wallet_nodes" edge of the CoinInfo entity.
func (ci *CoinInfo) QueryWalletNodes() *WalletNodeQuery {
	return (&CoinInfoClient{config: ci.config}).QueryWalletNodes(ci)
}

// Update returns a builder for updating this CoinInfo.
// Note that you need to call CoinInfo.Unwrap() before calling this method if this CoinInfo
// was returned from a transaction, and the transaction was committed or rolled back.
func (ci *CoinInfo) Update() *CoinInfoUpdateOne {
	return (&CoinInfoClient{config: ci.config}).UpdateOne(ci)
}

// Unwrap unwraps the CoinInfo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ci *CoinInfo) Unwrap() *CoinInfo {
	tx, ok := ci.config.driver.(*txDriver)
	if !ok {
		panic("ent: CoinInfo is not a transactional entity")
	}
	ci.config.driver = tx.drv
	return ci
}

// String implements the fmt.Stringer.
func (ci *CoinInfo) String() string {
	var builder strings.Builder
	builder.WriteString("CoinInfo(")
	builder.WriteString(fmt.Sprintf("id=%v", ci.ID))
	builder.WriteString(", name=")
	builder.WriteString(ci.Name)
	builder.WriteString(", unit=")
	builder.WriteString(ci.Unit)
	builder.WriteString(", need_signinfo=")
	builder.WriteString(fmt.Sprintf("%v", ci.NeedSigninfo))
	builder.WriteByte(')')
	return builder.String()
}

// CoinInfos is a parsable slice of CoinInfo.
type CoinInfos []*CoinInfo

func (ci CoinInfos) config(cfg config) {
	for _i := range ci {
		ci[_i].config = cfg
	}
}