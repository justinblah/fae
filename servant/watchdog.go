// +build !plan9,!windows

package servant

import (
	log "github.com/funkygao/log4go"
	"time"
)

func (this *FunServantImpl) runWatchdog() {
	if this.conf.WatchdogInterval == 0 {
		return
	}

	ticker := time.NewTicker(time.Duration(this.conf.WatchdogInterval) * time.Second)
	defer ticker.Stop()

	for _ = range ticker.C {
		log.Debug("mg: %v, mc: %v, lc: %d",
			this.mg.FreeConn(),
			this.mc.FreeConn(),
			this.lc.Len())
	}

}
