package client

import (
	"fmt"
	"net/url"

	Cli "github.com/docker/docker/cli"
	flag "github.com/docker/docker/pkg/mflag"
	//"github.com/docker/docker/opts"
)

//func presentInHelp(usage string) string { return usage }
//func absentFromHelp(string) string      { return "" }

// Usage: docker registry
func (cli *DockerCli) CmdRegistry(args ...string) (err error) {
	cmd := Cli.Subcmd("registry", nil, "lalalla.", true)
	cmd.Require(flag.Exact, 0)
	flInsecureRegistry:= cmd.String([]string{"-insecure-registry"}, "", "Enable insecure registry communication")
	//flInsecureRegistries := opts.NewListOpts(nil)
	//cmd.Var(&flInsecureRegistries, []string{"-insecure-registry"}, "Enable insecure registry communication")
	
	cmd.ParseFlags(args, true)
	
	v := url.Values{}
//	for _, registry := range flInsecureRegistries.GetAll() {
//		v.Add("registry", registry)
//	}
	v.Set("registry", *flInsecureRegistry)
	
	if _, _, err := readBody(cli.call("POST", fmt.Sprintf("/registry/add?%s", v.Encode()), nil, nil)); err != nil {
		fmt.Fprintf(cli.err, "%s\n", err)
		return fmt.Errorf("Error: failed to add insecure registry.")
	}
	
	return nil
}
