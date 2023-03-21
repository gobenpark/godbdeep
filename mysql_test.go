package main

import (
	"fmt"
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

type ParseOption int

const (
	Second         ParseOption = 1 << iota // Seconds field, default 0
	SecondOptional                         // Optional seconds field, default 0
	Minute                                 // Minutes field, default 0
	Hour                                   // Hours field, default 0
	Dom                                    // Day of month field, default *
	Month                                  // Month field, default *
	Dow                                    // Day of week field, default *
	DowOptional                            // Optional day of week field, default *
	Descriptor                             // Allow descriptors such as @monthly, @weekly, etc.
)

func TestParse(t *testing.T) {

	data := Dom

	result := Hour | Month | Dow | Descriptor

	fmt.Println(data & result)

}

//	//1
//	//10
//	//100
//	//1000
//	//10000
//	//100000
//	//1000000
//	result := 2 | 4
//	fmt.Println(result & 1)
//	fmt.Println(reflect.TypeOf(result))
//}

//
//func TestDB(t *testing.T) {
//	db, err := sql.Open("mysql", "ben:MnAobpfl8I/H@tcp(localhost:3328)/swit")
//	require.NoError(t, err)
//
//	d := &Database{
//		Name:   "swit",
//		Tables: []Table{},
//	}
//
//	rows, err := db.QueryContext(context.TODO(), `SELECT ic.TABLE_NAME,ic.COLUMN_NAME,ic.COLUMN_DEFAULT,ic.IS_NULLABLE,ic.DATA_TYPE,ic.CHARACTER_MAXIMUM_LENGTH,ic.CHARACTER_OCTET_LENGTH,ic.CHARACTER_SET_NAME,ic.COLLATION_NAME,ic.COLUMN_TYPE,ic.COLUMN_KEY
//FROM INFORMATION_SCHEMA.TABLES it
//join  information_schema.COLUMNS ic on it.TABLE_NAME = ic.TABLE_NAME WHERE ic.TABLE_SCHEMA = 'swit';`)
//	require.NoError(t, err)
//
//	for rows.Next() {
//		tab := Table{}
//		c := Column{}
//		rows.Scan(&tab.Name, &c.ColumnName, &c.ColumnDefault, &c.IsNullAble, &c.DataType,
//			&c.CharacterMaximumLength, &c.CharacterOctetLength, &c.CharacterSetName, &c.CollationName,
//			&c.ColumnType, &c.ColumnKey)
//		rows.Scan(&tab.Name)
//		tab.Columns = append(tab.Columns, c)
//		d.Tables = append(d.Tables, tab)
//		fmt.Println(tab.Columns)
//	}
//	rows.Close()
//	//
//	//	for i := range d.Tables {
//	//		rows, err = db.QueryContext(context.TODO(), `select COLUMN_NAME,
//	//       COLUMN_DEFAULT,
//	//       IS_NULLABLE,
//	//       DATA_TYPE,
//	//       CHARACTER_MAXIMUM_LENGTH,
//	//       CHARACTER_OCTET_LENGTH,
//	//       CHARACTER_SET_NAME,
//	//       COLLATION_NAME,
//	//       COLUMN_TYPE,
//	//       COLUMN_KEY
//	//from information_schema.COLUMNS
//	//WHERE TABLE_SCHEMA = ?
//	//  and TABLE_NAME = ?;`, d.Name, d.Tables[i].Name)
//	//
//	//		for rows.Next() {
//	//			c := Column{}
//	//			rows.Scan(&c.ColumnName, &c.ColumnDefault, &c.IsNullAble, &c.DataType,
//	//				&c.CharacterMaximumLength, &c.CharacterOctetLength, &c.CharacterSetName, &c.CollationName,
//	//				&c.ColumnType, &c.ColumnKey)
//	//			d.Tables[i].Columns = append(d.Tables[i].Columns, c)
//	//		}
//	//		rows.Close()
//	//	}
//	//
//	//	f, err := os.Create("express-asia.json")
//	//	require.NoError(t, err)
//	//	bt, err := json.Marshal(d)
//	//	require.NoError(t, err)
//	//	f.Write(bt)
//}
//
//func TestDB2(t *testing.T) {
//	db, err := sql.Open("mysql", "ben:MnAobpfl8I/H@tcp(localhost:3328)/swit")
//	require.NoError(t, err)
//
//	d := &Database{
//		Name:   "swit",
//		Tables: []Table{},
//	}
//
//	rows, err := db.QueryContext(context.TODO(), "SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = 'swit';")
//	require.NoError(t, err)
//
//	for rows.Next() {
//		tab := Table{}
//		rows.Scan(&tab.Name)
//		d.Tables = append(d.Tables, tab)
//	}
//	rows.Close()
//
//	for i := range d.Tables {
//		rows, err = db.QueryContext(context.TODO(), `select COLUMN_NAME,
//       COLUMN_DEFAULT,
//       IS_NULLABLE,
//       DATA_TYPE,
//       CHARACTER_MAXIMUM_LENGTH,
//       CHARACTER_OCTET_LENGTH,
//       CHARACTER_SET_NAME,
//       COLLATION_NAME,
//       COLUMN_TYPE,
//       COLUMN_KEY
//from information_schema.COLUMNS
//WHERE TABLE_SCHEMA = ?
//  and TABLE_NAME = ?;`, d.Name, d.Tables[i].Name)
//
//		for rows.Next() {
//			c := Column{}
//			rows.Scan(&c.ColumnName, &c.ColumnDefault, &c.IsNullAble, &c.DataType,
//				&c.CharacterMaximumLength, &c.CharacterOctetLength, &c.CharacterSetName, &c.CollationName,
//				&c.ColumnType, &c.ColumnKey)
//			d.Tables[i].Columns = append(d.Tables[i].Columns, c)
//		}
//		rows.Close()
//	}
//
//	f, err := os.Create("express-asia.json")
//	require.NoError(t, err)
//	bt, err := json.Marshal(d)
//	require.NoError(t, err)
//	f.Write(bt)
//}

func SPrintColumnDifference[T comparable](table, schema, infomation string, source, destination T) string {
	return fmt.Sprintf("table: %s column:%s infomation: %s source: %v dest: %v", table, schema, infomation, source, destination)
}

func Compare(a, b Table) {
	if !reflect.DeepEqual(a.Columns, b.Columns) {
		for i := range a.Columns {
			for j := range b.Columns {
				if a.Columns[i].ColumnName == b.Columns[j].ColumnName {
					if !reflect.DeepEqual(a.Columns[i], b.Columns[j]) {
						switch {
						case a.Columns[i].ColumnType != b.Columns[j].ColumnType:
							fmt.Println(SPrintColumnDifference(a.Name, a.Columns[i].ColumnName, "ColumnType", a.Columns[i].ColumnType, b.Columns[j].ColumnType))
						case a.Columns[i].ColumnKey != b.Columns[j].ColumnKey:
							fmt.Println(SPrintColumnDifference(a.Name, a.Columns[i].ColumnName, "ColumnKey", a.Columns[i].ColumnKey, b.Columns[j].ColumnKey))
						case a.Columns[i].ColumnDefault != b.Columns[j].ColumnDefault:
							fmt.Println(SPrintColumnDifference(a.Name, a.Columns[i].ColumnName, "ColumnDefault", a.Columns[i].ColumnDefault, b.Columns[j].ColumnDefault))
						case a.Columns[i].CollationName != b.Columns[j].CollationName:
							fmt.Println(SPrintColumnDifference(a.Name, a.Columns[i].ColumnName, "CollationName", a.Columns[i].CollationName, b.Columns[j].CollationName))
						case a.Columns[i].CharacterSetName != b.Columns[j].CharacterSetName:
							fmt.Println(SPrintColumnDifference(a.Name, a.Columns[i].ColumnName, "CharacterSetName", a.Columns[i].CharacterSetName, b.Columns[j].CharacterSetName))
						case a.Columns[i].DataType != b.Columns[j].DataType:
							fmt.Println(SPrintColumnDifference(a.Name, a.Columns[i].ColumnName, "DataType", a.Columns[i].DataType, b.Columns[j].DataType))
						case a.Columns[i].IsNullAble != b.Columns[j].IsNullAble:
							fmt.Println(SPrintColumnDifference(a.Name, a.Columns[i].ColumnName, "IsNullAble", a.Columns[i].IsNullAble, b.Columns[j].IsNullAble))
						case a.Columns[i].CharacterOctetLength != b.Columns[j].CharacterOctetLength:
							fmt.Println(SPrintColumnDifference(a.Name, a.Columns[i].ColumnName, "CharacterOctetLength", a.Columns[i].CharacterOctetLength, b.Columns[j].CharacterOctetLength))
						case a.Columns[i].CharacterMaximumLength != b.Columns[j].CharacterMaximumLength:
							fmt.Println(SPrintColumnDifference(a.Name, a.Columns[i].ColumnName, "CharacterMaximumLength", a.Columns[i].CharacterMaximumLength, b.Columns[j].CharacterMaximumLength))
						}

					}
				}
			}
		}
	}

}

//
//func TestDeepEqual(t *testing.T) {
//	bt, err := os.ReadFile("prod-asia.json")
//	require.NoError(t, err)
//	var prod Database
//
//	err = json.Unmarshal(bt, &prod)
//	require.NoError(t, err)
//
//	bt, err = os.ReadFile("express-asia.json")
//	require.NoError(t, err)
//	var exp Database
//
//	err = json.Unmarshal(bt, &exp)
//	require.NoError(t, err)
//
//	for i := range prod.Tables {
//		for j := range exp.Tables {
//			if prod.Tables[i].Name == exp.Tables[j].Name {
//				Compare(prod.Tables[i], exp.Tables[j])
//			}
//		}
//	}
//}
