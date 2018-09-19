package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/types"
	"github.com/containernetworking/cni/pkg/types/current"
	"github.com/containernetworking/cni/pkg/version"
)

type NetConf struct {
	types.NetConf
	IsTest string `json:"istest"`
}

func loadNetConf(bytes []byte) (*NetConf, string, error) {
	n := &NetConf{
		IsTest: "true",
	}
	if err := json.Unmarshal(bytes, n); err != nil {
		return nil, "", fmt.Errorf("failed to load netconf: %v", err)
	}
	return n, n.CNIVersion, nil
}

func cmdAdd(args *skel.CmdArgs) error {

	/*
		args.Args
		args.ContainerID
		args.IfName
		args.Netns
		args.Path
		args.StdinData
	*/
	_, cniVersion, err := loadNetConf(args.StdinData)
	if err != nil {
		return err
	}

	var interfaces = &current.Interface{
		Name:    "eth0",
		Mac:     "aabbccdd",
		Sandbox: "sandboxid",
	}
	var i = 1
	_, address, _ := net.ParseCIDR("10.10.10.10/24")
	gateway := net.ParseIP("10.10.10.1")
	var ips = &current.IPConfig{
		Version:   "4",
		Interface: &i,
		Address:   *address,
		Gateway:   gateway,
	}

	var dns = types.DNS{
		Nameservers: []string{"114.114.114.114", "8.8.8.8"},
		Domain:      "mkb.com",
		Search:      []string{"cluster.local"},
	}

	_, dst, _ := net.ParseCIDR("10.10.10.0/24")

	var route = &types.Route{
		Dst: *dst,
		GW:  net.ParseIP("10.10.10.1"),
	}

	var result current.Result
	result.CNIVersion = cniVersion
	result.Interfaces = []*current.Interface{interfaces}
	result.IPs = []*current.IPConfig{ips}
	result.Routes = []*types.Route{route}
	result.DNS = dns

	return types.PrintResult(&result, cniVersion)
}

func cmdDel(args *skel.CmdArgs) error {

	//just test interface
	_, cniVersion, err := loadNetConf(args.StdinData)
	if err != nil {
		return err
	}

	var interfaces = &current.Interface{
		Name:    "eth0",
		Mac:     "aabbccdd",
		Sandbox: "sandboxid",
	}
	var i = 1
	_, address, _ := net.ParseCIDR("10.10.10.10/24")
	gateway := net.ParseIP("10.10.10.1")
	var ips = &current.IPConfig{
		Version:   "4",
		Interface: &i,
		Address:   *address,
		Gateway:   gateway,
	}

	var dns = types.DNS{
		Nameservers: []string{"114.114.114.114", "8.8.8.8"},
		Domain:      "mkb.com",
		Search:      []string{"cluster.local"},
	}

	_, dst, _ := net.ParseCIDR("10.10.10.0/24")

	var route = &types.Route{
		Dst: *dst,
		GW:  net.ParseIP("10.10.10.1"),
	}

	var result current.Result
	result.CNIVersion = cniVersion
	result.Interfaces = []*current.Interface{interfaces}
	result.IPs = []*current.IPConfig{ips}
	result.Routes = []*types.Route{route}
	result.DNS = dns

	os.Create("/tmp/del")

	return types.PrintResult(&result, cniVersion)

}

func cmdGet(args *skel.CmdArgs) error {
	return fmt.Errorf("not implemented")
}

func main() {
	skel.PluginMain(cmdAdd, cmdGet, cmdDel, version.All, "TODO")
}
