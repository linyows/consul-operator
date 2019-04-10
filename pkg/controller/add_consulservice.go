package controller

import (
	"github.com/linyows/consul-operator/pkg/controller/consulservice"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, consulservice.Add)
}
