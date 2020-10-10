package main

import (
	"github.com/nkomiljon/wallet_new/pkg/wallet"
)

func main()  {
	svc := &wallet.Service{}
    svc.RegisterAccount("+992929000001")
	svc.RegisterAccount("+992929000002")
	svc.RegisterAccount("+992929000003")
	svc.ExportToFile("data/export.txt")
	//svc.ImportFromFile("data/import.txt")
}
