/*
** Copyright [2013-2015] [Megam Systems]
**
** Licensed under the Apache License, Version 2.0 (the "License");
** you may not use this file except in compliance with the License.
** You may obtain a copy of the License at
**
** http://www.apache.org/licenses/LICENSE-2.0
**
** Unless required by applicable law or agreed to in writing, software
** distributed under the License is distributed on an "AS IS" BASIS,
** WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
** See the License for the specific language governing permissions and
** limitations under the License.
 */
package ceph

import (
	"fmt"
	"github.com/megamsys/libgo/cmd"
	"github.com/megamsys/megdc/handler"
	"launchpad.net/gnuflag"
	"reflect"
	//	"strconv"
)

type Cephinstall struct {
	Fs           			*gnuflag.FlagSet

	CephInstall  	bool
	CEPH_LOG       string
  Ceph_user      string
  Ceph_password  string
  Ceph_group     string
  User_home      string
	Host		 			string
	Osd1	 			string
	Osd2   			string
  Osd3        string
	Quiet        			bool
}

func (g *Cephinstall) Info() *cmd.Info {
	desc := `starts megdc.

If you use the '--quiet' flag megdc doesn't print the logs.

`
	return &cmd.Info{
		Name:    "cephinstall",
		Usage:   `cephinstall [--ceph] [--ceph_user] ...`,
		Desc:    desc,
		MinArgs: 0,
	}
}

func (c *Cephinstall) Run(context *cmd.Context) error {
	fmt.Println("[main] starting megdc ...")

	packages := make(map[string]string)
	options := make(map[string]string)

	s := reflect.ValueOf(c).Elem()
	typ := s.Type()
	if s.Kind() == reflect.Struct {
		for i := 0; i < s.NumField(); i++ {
			key := s.Field(i)
			value := s.FieldByName(typ.Field(i).Name)
			switch key.Interface().(type) {
			case bool:
				if value.Bool() {
					packages[typ.Field(i).Name] = typ.Field(i).Name
				}
			case string:
				if value.String() != "" {
					options[typ.Field(i).Name] = value.String()
				}
			}
		}
	}

	if handler, err := handler.NewHandler(); err != nil {
		return err
	} else {
		handler.SetTemplates(packages, options)
        err := handler.Run()
        if err != nil {
        	return err
        }
	}

	// goodbye.
	return nil
}

func (c *Cephinstall) Flags() *gnuflag.FlagSet {
	if c.Fs == nil {
		c.Fs = gnuflag.NewFlagSet("megdc", gnuflag.ExitOnError)

		/* Install package commands */
		c.Fs.BoolVar(&c.CephInstall, "ceph", false, "Install ceph package")
		c.Fs.BoolVar(&c.CephInstall, "c", false, "Install ceph package")


    c.Fs.StringVar(&c.CEPH_LOG, "ceph_log", "", "ceph_log path for hosted machine")
    c.Fs.StringVar(&c.CEPH_LOG, "l", "", "ceph_log path for hosted machine")
		c.Fs.StringVar(&c.Ceph_user, "ceph_user", "", "ceph_user for hosted machine")
		c.Fs.StringVar(&c.Ceph_user, "u", "", "ceph_user for hosted machine")
		c.Fs.StringVar(&c.Ceph_password, "ceph_password", "", "ceph_password for hosted machine")
		c.Fs.StringVar(&c.Ceph_password, "p", "", "ceph_password for hosted machine")
    c.Fs.StringVar(&c.Ceph_group, "ceph_group", "", "ceph_group for hosted machine")
		c.Fs.StringVar(&c.Ceph_group, "g", "", "ceph_group for hosted machine")
    c.Fs.StringVar(&c.User_home, "user_home", "", "user_home path for hosted machine")
		c.Fs.StringVar(&c.User_home, "uh", "", "user_home path for hosted machine")
    c.Fs.StringVar(&c.Host, "host", "", "host address for machine")
		c.Fs.StringVar(&c.Host, "h", "", "host address for machine")
    c.Fs.StringVar(&c.Osd1, "osd1", "", "osd1 storage drive for hosted machine")
		c.Fs.StringVar(&c.Osd1, "os1", "", "osd1 strorage drive for hosted machine")
    c.Fs.StringVar(&c.Osd2, "osd2", "", "osd2 storage drive for hosted machine")
		c.Fs.StringVar(&c.Osd2, "os2", "", "osd2 strorage drive for hosted machine")
    c.Fs.StringVar(&c.Osd3, "osd3", "", "osd3 storage drive for hosted machine")
    c.Fs.StringVar(&c.Osd3, "os3", "", "osd3 strorage drive for hosted machine")


		c.Fs.BoolVar(&c.Quiet, "quiet", false, "")
		c.Fs.BoolVar(&c.Quiet, "q", false, "")
	}
	return c.Fs
}
