package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/you06/sqlsmith-client/executor"
	"github.com/you06/sqlsmith-client/util"
	"github.com/juju/errors"
	"github.com/ngaut/log"
	_ "github.com/pingcap/tidb/types/parser_driver"
)

var (
	printVersion bool
	dsn1         string
	dsn2         string
	printSchema  bool
)

func init() {
	flag.BoolVar(&printVersion, "V", false, "print version")
	flag.BoolVar(&printSchema, "schema", false, "print schema and exit")
	flag.StringVar(&dsn1, "dsn1", "", "dsn1")
	flag.StringVar(&dsn2, "dsn2", "", "dsn2")
}

func main() {
	flag.Parse()
	if printVersion {
		util.PrintInfo()
		os.Exit(0)
	}

	var (
		exec *executor.Executor
		err error
	)
	if dsn1 == "" {
		log.Fatalf("dsn1 can not be empty")
	} else if dsn2 == "" {
		exec, err = executor.New(dsn1)
	} else {
		exec, err = executor.NewABTest(dsn1, dsn2)
	}
	if err != nil {
		log.Fatalf("create executor error %v", errors.ErrorStack(err))
	}

	if printSchema {
		if err := exec.PrintSchema(); err != nil {
			log.Fatalf("print schema err %v", errors.ErrorStack(err))
		}
		os.Exit(0)
	}
	go exec.Start()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		os.Kill,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	sig := <-sc
	log.Infof("Got signal %d to exit.", sig)
}
