package main

import (
	"errors"

	"github.com/Unknwon/goconfig"
)

type Package struct {
	PackageName string
	PackageRoot string
	//Args        string
}

func Loadconfig(filename string) ([]*Package, error) {
	c, err := goconfig.LoadConfigFile(filename)
	if err != nil {
		return nil, err
	}
	sectionlist := c.GetSectionList()

	if len(sectionlist) <= 0 {
		return nil, errors.New("config file format err.")
	}

	result := make([]*Package, len(sectionlist))
	for i := range sectionlist {

		packageroot, err := c.GetValue(sectionlist[i], "path")
		if err != nil {
			return nil, err
		}
		conf := new(Package)
		conf.PackageRoot = packageroot
		conf.PackageName = sectionlist[i]
		//conf.Args, _ = c.GetValue(sectionlist[i], "args")
		result[i] = conf

	}
	return result, nil
}
