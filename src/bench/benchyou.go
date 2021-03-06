/*
 * benchyou
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"runtime"
	"xcmd"
)

var (
	write_threads      int
	read_threads       int
	update_threads     int
	delete_threads     int
	mysql_host         string
	mysql_port         int
	mysql_user         string
	mysql_password     string
	mysql_db           string
	mysql_table_engine string
	mysql_range_order  string
	mysql_enable_xa    int
	rows_per_insert    int
	batch_per_commit   int
	max_time           int
	max_request        uint64
	oltp_tables_count  int
	ssh_host           string
	ssh_user           string
	ssh_password       string
	ssh_port           int
)

var (
	rootCmd = &cobra.Command{
		Use:        "benchyou",
		Short:      "",
		SuggestFor: []string{"benchyou"},
	}
)

func init() {
	cobra.EnableCommandSorting = false
	rootCmd.PersistentFlags().IntVar(&write_threads, "write-threads", 32, "number of write threads to use(Default 32)")
	rootCmd.PersistentFlags().IntVar(&read_threads, "read-threads", 32, "number of read threads to use(Default 32)")
	rootCmd.PersistentFlags().IntVar(&update_threads, "update-threads", 0, "number of update threads to use(Default 0)")
	rootCmd.PersistentFlags().IntVar(&delete_threads, "delete-threads", 0, "number of delete threads to use(Default 0)")
	rootCmd.PersistentFlags().StringVar(&mysql_host, "mysql-host", "", "MySQL server host(Default NULL)")
	rootCmd.PersistentFlags().IntVar(&mysql_port, "mysql-port", 3306, "MySQL server port(Default 3306)")
	rootCmd.PersistentFlags().StringVar(&mysql_user, "mysql-user", "benchyou", "MySQL user(Default benchyou)")
	rootCmd.PersistentFlags().StringVar(&mysql_password, "mysql-password", "benchyou", "MySQL password(Default benchyou)")
	rootCmd.PersistentFlags().StringVar(&mysql_db, "mysql-db", "sbtest", "MySQL database name(Default sbtest)")
	rootCmd.PersistentFlags().StringVar(&mysql_table_engine, "mysql-table-engine", "tokudb", "storage engine to use for the test table {tokudb,innodb,...}(Default tokudb)")
	rootCmd.PersistentFlags().StringVar(&mysql_range_order, "mysql-range-order", "ASC", "range query sort the result-set in {ASC|DESC} (Default ASC)")
	rootCmd.PersistentFlags().IntVar(&mysql_enable_xa, "mysql-enable-xa", 0, "enable MySQL xa transaction for insertion {0|1} (Default 0)")
	rootCmd.PersistentFlags().IntVar(&rows_per_insert, "rows-per-insert", 1, "#rows per insert(Default 1)")
	rootCmd.PersistentFlags().IntVar(&batch_per_commit, "batch-per-commit", 1, "#rows per transaction(Default 1)")
	rootCmd.PersistentFlags().IntVar(&max_time, "max-time", 3600, "limit for total execution time in seconds(Default 3600)")
	rootCmd.PersistentFlags().Uint64Var(&max_request, "max-request", 0, "limit for total requests, including write and read(Default 0, means no limits)")
	rootCmd.PersistentFlags().IntVar(&oltp_tables_count, "oltp-tables-count", 8, "number of tables to create(Default 8)")
	rootCmd.PersistentFlags().StringVar(&ssh_host, "ssh-host", "", "SSH server host(Default NULL, same as mysql-host)")
	rootCmd.PersistentFlags().StringVar(&ssh_user, "ssh-user", "benchyou", "SSH server user(Default benchyou)")
	rootCmd.PersistentFlags().StringVar(&ssh_password, "ssh-password", "benchyou", "SSH server password(Default benchyou)")
	rootCmd.PersistentFlags().IntVar(&ssh_port, "ssh-port", 22, "SSH server port(Default 22)")

	rootCmd.AddCommand(xcmd.NewPrepareCommand())
	rootCmd.AddCommand(xcmd.NewCleanupCommand())
	rootCmd.AddCommand(xcmd.NewRandomCommand())
	rootCmd.AddCommand(xcmd.NewSeqCommand())
	rootCmd.AddCommand(xcmd.NewRangeCommand())
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
