package main

import (
	"context"

	log "github.com/Sirupsen/logrus"

	"github.com/vmware/govmomi/examples"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

func main() {
	ctx := context.Background()

	c, err := examples.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	var ticket string
	req := types.AcquireCloneTicket{
		This: *c.ServiceContent.SessionManager,
	}

	res, err := methods.AcquireCloneTicket(ctx, c, &req)
	if err != nil {
		log.Fatal(err)
	}

	ticket = res.Returnval
	log.Infof("Ticket: %s", ticket)
}
