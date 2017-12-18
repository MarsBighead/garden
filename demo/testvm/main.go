package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
	yaml "gopkg.in/yaml.v2"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	currentPath, _ := os.Getwd()
	srv := new(Server)
	body, err := ioutil.ReadFile(currentPath + "/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(body, srv)
	if err != nil {
		log.Fatal(err)
	}
	srv.pollFind(ctx)
	srv.pollContainer(ctx)

}

//Server Type Server for object server in config.yaml
type Server struct {
	Product  string `yaml:"product"`
	URI      string `yaml:"url"`
	IP       string `yaml:"ip"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Insecure bool   `yaml:"insecure"`
	Switch   string `yaml:"switch"`
}

// PollVMRaw UM sruct reference for dataBase
// MemSize, MemReserveSize: unit of measurement is MB
type PollVMRaw struct {
	ID             int    `vsql:"type:BIGSERIAL;column:id;primary_key"           csv:"id"`
	VcCollectionID int    `vsql:"type:BIGINT;column:vcCollectionId"              csv:"vcCollectionId"`
	Moref          string `vsql:"type:varchar(256);column:moref"`
	Name           string `vsql:"type:varchar(256);column:vmName"`
	InstanceUUUID  string `vsql:"type:varchar(256);column:instanceUuuid"         csv:"instanceUuuid"`
	BiosUUID       string `vsql:"type:varchar(128);column:biosUuid"`
	GuestFullName  string `vsql:"type:varchar(1024);column:guestFullName"        csv:"guestFullName"`
	ConnectOn      string `vsql:"type:BOOLEAN;column:connectOn"                  csv:"-"`
	PowerState     string `vsql:"type:SMALLINT;column:powerState"                csv:"State"`
	MemSizeMB      int32  `vsql:"type:integer;column:memSizeMB"                  csv:"memSizeMB"`
	ResMemSizeMB   int32  `vsql:"type:integer;column:resMemSizeMB"               csv:"resMemSizeMB"`
	HostMoref      string `vsql:"type:varchar(256);column:hostMoref"`
	CPUCount       int32  `vsql:"type:smallint;column:cpuCount"                  csv:"cpuCount"`
	CPUReservation int32  `vsql:"type:int;column:cpuReservation"                 csv:"cpuMHz"`
	//BootTime       *time.Time `vsql:"type:timestamp with time zone;column:bootTime"  csv:"Boot Time"`
	ChangeTime *time.Time `vsql:"type:timestamp with time zone NOT NULL;column:changeTime"  csv:"Boot Time"`
}

func (srv *Server) pollFind(ctx context.Context) error {
	now := time.Now()
	c, err := NewClient(ctx, *srv)
	if err != nil {
		return err
	}
	defer c.Logout(ctx)
	finder := find.NewFinder(c.Client, true)
	dc, err := finder.DefaultDatacenter(ctx)
	if err != nil {
		return err
	}
	finder.SetDatacenter(dc)

	args := []string{"*"}
	//var props []string
	props := []string{
		"summary",
	}
	var objects []*object.VirtualMachine
	for _, arg := range args {
		objects, err = finder.VirtualMachineList(ctx, arg)
		if err != nil {
			return err
		}
	}
	if len(objects) != 0 {
		refs := make([]types.ManagedObjectReference, 0, len(objects))
		for _, obj := range objects {
			refs = append(refs, obj.Reference())
		}

		pc := property.DefaultCollector(c.Client)
		vms := new([]mo.VirtualMachine)
		err = pc.Retrieve(ctx, refs, props, vms)
		if err != nil {
			return err
		}
		fmt.Println("By finder:length of vms is ", len(*vms))
		for _, vm := range *vms {
			pvmr := PollVMRaw{
				Moref:          vm.Summary.Vm.Value,
				Name:           vm.Summary.Config.Name,
				InstanceUUUID:  vm.Summary.Config.InstanceUuid,
				BiosUUID:       vm.Summary.Config.Uuid,
				GuestFullName:  vm.Summary.Config.GuestFullName,
				MemSizeMB:      vm.Summary.Config.MemorySizeMB,
				ResMemSizeMB:   vm.Summary.Config.MemoryReservation,
				HostMoref:      vm.Summary.Runtime.Host.Value,
				CPUCount:       vm.Summary.Config.NumCpu,
				CPUReservation: vm.Summary.Config.CpuReservation,
			}
			fmt.Println("Poll", pvmr.Moref)
			//vm.Summary.Runtime.ConnectionState, vm.Summary.Runtime.PowerState)
		}
	}
	fmt.Println("running time:", time.Now().Sub(now))
	return nil
}

// pollContainer Load data to table Vm
func (srv *Server) pollContainer(ctx context.Context) {
	now := time.Now()
	// Retrieve summary property for all machines
	// Reference: http://pubs.vmware.com/vsphere-60/topic/com.vmware.wssdk.apiref.doc/vim.VirtualMachine.html
	c, err := NewClient(ctx, *srv)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Logout(ctx)
	m := view.NewManager(c.Client)
	v, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, nil, true)
	if err != nil {
		log.Println(err)
	}
	var vms []mo.VirtualMachine
	err = v.Retrieve(ctx, []string{"VirtualMachine"}, []string{"summary"}, &vms)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("By ContainerView length of vms is ", len(vms))
	for _, vm := range vms {
		pvmr := &PollVMRaw{
			Moref: vm.Summary.Vm.Value,
			Name:  vm.Summary.Config.Name,
			//GuestOS: vm.Summary.Config.GuestFullName,
			//GuestID: vm.Summary.Config.GuestId,
			HostMoref:    vm.Summary.Runtime.Host.Value,
			MemSizeMB:    vm.Summary.Config.MemorySizeMB,
			ResMemSizeMB: vm.Summary.Config.MemoryReservation,
			PowerState:   string(vm.Summary.Runtime.PowerState),
			ConnectOn:    string(vm.Summary.Runtime.ConnectionState),
			//IPAddress:     vm.Summary.Guest.IpAddress,
			InstanceUUUID: vm.Summary.Config.InstanceUuid,
			BiosUUID:      vm.Summary.Config.Uuid,
			//BootTime:      vm.Summary.Runtime.BootTime,
		}
		fmt.Println("Poll", pvmr.Moref)

	}
	fmt.Println("running time:", time.Now().Sub(now))

}

// NewClient Build NewClient test with config.yaml
func NewClient(ctx context.Context, p Server) (*govmomi.Client, error) {
	// Parse URL from string
	u, err := soap.ParseURL(p.URI)
	if err != nil {
		return nil, err
	}

	// Override username and/or password as required
	formatURL(u, p.Username, p.Password)
	// Connect and log in to ESX or vCenter
	return govmomi.NewClient(ctx, u, p.Insecure)
}

// formatURL format url for connect vSphere/vCenter....
func formatURL(u *url.URL, cfgUsername, cfgPassword string) {
	// Override username if provided by configure
	if cfgUsername != "" {
		var password string
		var ok bool

		if u.User != nil {
			password, ok = u.User.Password()
		}

		if ok {
			u.User = url.UserPassword(cfgUsername, password)
		} else {
			u.User = url.User(cfgUsername)
		}
	}

	// Override password if provided
	if cfgPassword != "" {
		var username string

		if u.User != nil {
			username = u.User.Username()
		}

		u.User = url.UserPassword(username, cfgPassword)
	}
}
