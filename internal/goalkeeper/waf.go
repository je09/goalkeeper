package goalkeeper

import (
	"Goalkeeper/pkg/config"
	"github.com/corazawaf/coraza/v2"
	"github.com/corazawaf/coraza/v2/seclang"
	"github.com/corazawaf/coraza/v2/types"
	"log"
	"net/http"
)

var waf *coraza.Waf

func InitWaf(cnf config.Config) error {
	waf = coraza.NewWaf()
	err := waf.SetDebugLogLevel(5)
	if err != nil {
		log.Fatal(err)
	}
	parser, _ := seclang.NewParser(waf)
	files := []string{
		cnf.Modsec.Conf,
		cnf.Modsec.Setup,
		cnf.Modsec.Ruleset,
	}
	for _, f := range files {
		if err := parser.FromFile(f); err != nil {
			return err
		}
	}

	return nil
}

func processRequest(w http.ResponseWriter, r *http.Request) int {
	tx := waf.NewTransaction()
	defer tx.ProcessLogging()
	in, err := tx.ProcessRequest(r)
	if err != nil {
		return 405
	}

	if in != nil {
		processInterruption(w, in)
		return in.Status
	}

	return 200
}

func processInterruption(w http.ResponseWriter, in *types.Interruption) {
	if in.Status == 0 {
		in.Status = 403
	}
	log.Printf("Request denied. Found rule: %d, action: %s.", in.RuleID, in.Action)
}

type interceptor struct {
	origWriter  http.ResponseWriter
	tx          *coraza.Transaction
	headersSent bool
}

func (i *interceptor) WriteHeader(rc int) {
	if i.headersSent {
		return
	}
	for k, vv := range i.origWriter.Header() {
		for _, v := range vv {
			i.tx.AddResponseHeader(k, v)
		}
	}
	i.headersSent = true
	if it := i.tx.ProcessResponseHeaders(rc, "http/1.1"); it != nil {
		processInterruption(i.origWriter, it)
		return
	}
	i.origWriter.WriteHeader(rc)
}

func (i *interceptor) Write(b []byte) (int, error) {
	return i.tx.ResponseBodyBuffer.Write(b)
}

func (i *interceptor) Header() http.Header {
	return i.origWriter.Header()
}
