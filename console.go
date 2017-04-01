package client

import (
	"github.com/nilshell/xmlrpc"
)

type Console XenAPIObject

func (self *Console) GetConsoleAll() (consoles []*Console, err error) {
	consoles = make([]*Console, 0)
	result := APIResult{}
	err = self.Client.APICall(&result, "console.get_all")
	if err != nil {
		return
	}

	for _, elem := range result.Value.([]interface{}) {
		vm := new(Console)
		vm.Ref = elem.(string)
		vm.Client = self.Client
		consoles = append(consoles, vm)
	}

	return
}

func (self *Console) GetUUID() (uuid string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "console.get_uuid", self.Ref)
	if err != nil {
		return
	}
	uuid = result.Value.(string)
	return
}

func (self *Console) GetProtocal() (ptl string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "console.get_protocol", self.Ref)
	if err != nil {
		return
	}
	ptl = result.Value.(string)
	return
}

func (self *Console) GetLocation() (loc string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "console.get_location", self.Ref)
	if err != nil {
		return
	}
	loc = result.Value.(string)
	return
}

func (self *Console) GetOtherConfig() (ret map[string]string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "console.get_other_config", self.Ref)
	if err != nil {
		return
	}
	ret = make(map[string]string)
	for k, v := range result.Value.(xmlrpc.Struct) {
		if temp, ok := v.(string); ok {
			ret[k] = temp
		}
	}
	return
}
