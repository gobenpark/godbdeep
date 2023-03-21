package main

import (
	"database/sql"

	"github.com/spf13/cobra"
)

var mysqlDBCmd = &cobra.Command{
	Use:                        "",
	Aliases:                    nil,
	SuggestFor:                 nil,
	Short:                      "",
	GroupID:                    "",
	Long:                       "",
	Example:                    "",
	ValidArgs:                  nil,
	ValidArgsFunction:          nil,
	Args:                       nil,
	ArgAliases:                 nil,
	BashCompletionFunction:     "",
	Deprecated:                 "",
	Annotations:                nil,
	Version:                    "",
	PersistentPreRun:           nil,
	PersistentPreRunE:          nil,
	PreRun:                     nil,
	PreRunE:                    nil,
	Run:                        nil,
	RunE:                       nil,
	PostRun:                    nil,
	PostRunE:                   nil,
	PersistentPostRun:          nil,
	PersistentPostRunE:         nil,
	FParseErrWhitelist:         cobra.FParseErrWhitelist{},
	CompletionOptions:          cobra.CompletionOptions{},
	TraverseChildren:           false,
	Hidden:                     false,
	SilenceErrors:              false,
	SilenceUsage:               false,
	DisableFlagParsing:         false,
	DisableAutoGenTag:          false,
	DisableFlagsInUseLine:      false,
	DisableSuggestions:         false,
	SuggestionsMinimumDistance: 0,
}

type Database struct {
	Name    string   `json:"name"`
	Columns []Column `json:"columns"`
}

type Table struct {
	Name    string   `json:"name"`
	Columns []Column `json:"columns" diff:"columns"`
}

type Column struct {
	TableName              string         `json:"tableName" diff:"tableName"`
	ColumnName             string         `json:"columnName" diff:"columnName"`
	ColumnDefault          sql.NullString `json:"columnDefault" diff:"columnDefault"`
	IsNullAble             string         `json:"isNullAble" diff:"isNullAble"`
	DataType               string         `json:"dataType" diff:"dataType"`
	CharacterMaximumLength sql.NullInt64  `json:"characterMaximumLength" diff:"characterMaximumLength"`
	CharacterOctetLength   sql.NullInt64  `json:"characterOctetLength" diff:"characterOctetLength"`
	CharacterSetName       sql.NullString `json:"characterSetName" diff:"characterSetName"`
	CollationName          sql.NullString `json:"collationName" diff:"collationName"`
	ColumnType             string         `json:"columnType" diff:"columnType"`
	ColumnKey              string         `json:"columnKey" diff:"columnKey"`
}

func (d Database) Compare(b ...Database) {
	//
	//if !reflect.DeepEqual(a.Columns, b.Columns) {
	//	for i := range d.Columns {
	//		for j := range b.Columns {
	//			if d.Columns[i].ColumnName == b.Columns[j].ColumnName {
	//				if !reflect.DeepEqual(d.Columns[i], b.Columns[j]) {
	//					switch {
	//					case a.Columns[i].ColumnType != b.Columns[j].ColumnType:
	//						fmt.Println(SPrintColumnDifference(a.Name, a.Columns[i].ColumnName, "ColumnType", a.Columns[i].ColumnType, b.Columns[j].ColumnType))
	//					case a.Columns[i].ColumnKey != b.Columns[j].ColumnKey:
	//						fmt.Println(SPrintColumnDifference(a.Name, a.Columns[i].ColumnName, "ColumnKey", a.Columns[i].ColumnKey, b.Columns[j].ColumnKey))
	//					case a.Columns[i].ColumnDefault != b.Columns[j].ColumnDefault:
	//						fmt.Println(SPrintColumnDifference(a.Name, a.Columns[i].ColumnName, "ColumnDefault", a.Columns[i].ColumnDefault, b.Columns[j].ColumnDefault))
	//					case a.Columns[i].CollationName != b.Columns[j].CollationName:
	//						fmt.Println(SPrintColumnDifference(a.Name, a.Columns[i].ColumnName, "CollationName", a.Columns[i].CollationName, b.Columns[j].CollationName))
	//					case a.Columns[i].CharacterSetName != b.Columns[j].CharacterSetName:
	//						fmt.Println(SPrintColumnDifference(a.Name, a.Columns[i].ColumnName, "CharacterSetName", a.Columns[i].CharacterSetName, b.Columns[j].CharacterSetName))
	//					case a.Columns[i].DataType != b.Columns[j].DataType:
	//						fmt.Println(SPrintColumnDifference(a.Name, a.Columns[i].ColumnName, "DataType", a.Columns[i].DataType, b.Columns[j].DataType))
	//					case a.Columns[i].IsNullAble != b.Columns[j].IsNullAble:
	//						fmt.Println(SPrintColumnDifference(a.Name, a.Columns[i].ColumnName, "IsNullAble", a.Columns[i].IsNullAble, b.Columns[j].IsNullAble))
	//					case a.Columns[i].CharacterOctetLength != b.Columns[j].CharacterOctetLength:
	//						fmt.Println(SPrintColumnDifference(a.Name, a.Columns[i].ColumnName, "CharacterOctetLength", a.Columns[i].CharacterOctetLength, b.Columns[j].CharacterOctetLength))
	//					case a.Columns[i].CharacterMaximumLength != b.Columns[j].CharacterMaximumLength:
	//						fmt.Println(SPrintColumnDifference(a.Name, a.Columns[i].ColumnName, "CharacterMaximumLength", a.Columns[i].CharacterMaximumLength, b.Columns[j].CharacterMaximumLength))
	//						fmt.Println(SPrintColumnDifference(a.Name, a.Columns[i].ColumnName, "CharacterMaximumLength", a.Columns[i].CharacterMaximumLength, b.Columns[j].CharacterMaximumLength))
	//					}
	//
	//				}
	//			}
	//		}
	//	}
	//}

}
