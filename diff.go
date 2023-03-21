package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/briandowns/spinner"
	_ "github.com/go-sql-driver/mysql"
	"github.com/olekukonko/tablewriter"
	"github.com/r3labs/diff/v3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	cfgFile string

	diffCmd = &cobra.Command{
		Use:                        "diff",
		Short:                      "This is the Diff Two Database Schema",
		Long:                       "This is compare two database schema on mysql,postgres or etc",
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
		Run:                        Diff,
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
)

func init() {
	rootCmd.AddCommand(diffCmd)
	diffCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "-c ./source.yaml")
	initConfig()
}

func initConfig() {
	viper.SetConfigFile(cfgFile)
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func Diff(cmd *cobra.Command, args []string) {
	s := spinner.New(spinner.CharSets[7], 100*time.Millisecond) // Build our new spinner
	s.Color("red", "bold")
	s.Start()                   // Start the spinner
	time.Sleep(4 * time.Second) // Run for some time to simulate work

	ctx := context.Background()
	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		zap.L().Panic("err marshal")
	}
	s.Start()
	stb, err := TableData(ctx, c.Source.ID, c.Source.Password, c.Source.IP, c.Source.Port, c.Source.Name, c.Source.Database)
	if err != nil {
		zap.L().Panic("TableData error", zap.Error(err))
	}

	stb2, err := TableData(ctx, c.Source.ID, c.Source.Password, c.Source.IP, c.Source.Port, c.Source.Name, c.Source.Database)
	if err != nil {
		zap.L().Panic("TableData error", zap.Error(err))
	}

	stb2.Columns[0].ColumnDefault = sql.NullString{
		String: "hi",
		Valid:  false,
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Type", "From", "To", "Path"})

	l, err := diff.Diff(stb, stb2)
	if err != nil {
		zap.L().Panic("diff err")
	}

	for _, i := range l {
		idx, err := strconv.Atoi(i.Path[1])
		if err != nil {
			zap.L().Error("index error", zap.Error(err))
			break
		}

		table.Append([]string{stb.Columns[idx].ColumnName, i.Type, fmt.Sprintf("%v", i.From), fmt.Sprintf("%v", i.To), fmt.Sprintf("%v", i.Path)})
	}
	s.Stop()
	table.Render()

}

func TableData(ctx context.Context, id, password, ip, port, name, database string) (*Database, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", id, password, ip, port, database))
	if err != nil {
		zap.L().Panic("err", zap.Error(err))
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return nil, err
	}

	d := &Database{
		Name:    name,
		Columns: []Column{},
	}

	rows, err := db.QueryContext(ctx, fmt.Sprintf(`
SELECT distinct
       ic.TABLE_NAME as TABLE_NAME,
       COLUMN_NAME,
       COLUMN_DEFAULT,
       IS_NULLABLE,
       DATA_TYPE,
       CHARACTER_MAXIMUM_LENGTH,
       CHARACTER_OCTET_LENGTH,
       CHARACTER_SET_NAME,
       COLLATION_NAME,
       COLUMN_TYPE,
       COLUMN_KEY
from information_schema.COLUMNS ic join INFORMATION_SCHEMA.TABLES it on ic.TABLE_SCHEMA = it.TABLE_SCHEMA
where it.TABLE_SCHEMA = '%s'
`, database))
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		c := Column{}
		if err := rows.Scan(&c.TableName, &c.ColumnName, &c.ColumnDefault, &c.IsNullAble, &c.DataType,
			&c.CharacterMaximumLength, &c.CharacterOctetLength, &c.CharacterSetName, &c.CollationName,
			&c.ColumnType, &c.ColumnKey); err != nil {
			zap.L().Error("scan error", zap.Error(err))
			continue
		}

		d.Columns = append(d.Columns, c)
	}
	if err := rows.Close(); err != nil {
		zap.L().Panic("close error", zap.Error(err))
	}
	return d, nil
}
