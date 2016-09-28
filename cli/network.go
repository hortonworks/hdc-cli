package cli

import (
	log "github.com/Sirupsen/logrus"
	"github.com/hortonworks/hdc-cli/client/networks"
	"github.com/hortonworks/hdc-cli/models"
	"github.com/urfave/cli"
	"strconv"
	"sync"
	"time"
)

var NetworkHeader []string = []string{"ID", "Name"}

type NetworkList struct {
	Id   int64  `json:"Id" yaml:"Id"`
	Name string `json:"Name" yaml:"Name"`
}

func (n *NetworkList) DataAsStringArray() []string {
	return []string{strconv.FormatInt(n.Id, 10), n.Name}
}

func (c *Cloudbreak) CreateNetwork(skeleton ClusterSkeleton, channel chan int64, wg *sync.WaitGroup) {
	defer timeTrack(time.Now(), "create network")
	defer wg.Done()

	networkName := "net" + strconv.FormatInt(time.Now().UnixNano(), 10)

	vpc := skeleton.Network
	var vpcParams = make(map[string]interface{})
	if vpc != nil && len(vpc.VpcId) > 0 && len(vpc.SubnetId) > 0 {
		vpcParams["vpcId"] = vpc.VpcId
		vpcParams["subnetId"] = vpc.SubnetId
	} else {
		defaultNetwork := c.GetNetwork("aws-network")
		vpcParams["vpcId"] = defaultNetwork.Parameters["vpcId"]
		vpcParams["internetGatewayId"] = defaultNetwork.Parameters["internetGatewayId"]
	}

	network := models.NetworkJSON{
		Name:          networkName,
		CloudPlatform: "AWS",
		Parameters:    vpcParams,
	}

	log.Infof("[CreateNetwork] sending network create request with name: %s", networkName)
	resp, err := c.Cloudbreak.Networks.PostNetworksAccount(&networks.PostNetworksAccountParams{&network})

	if err != nil {
		logErrorAndExit(c.CreateNetwork, err.Error())
	}

	log.Infof("[CreateNetwork] network created, id: %d", resp.Payload.ID)
	channel <- resp.Payload.ID
}

func CreateNetworkCommand(c *cli.Context) error {
	checkRequiredFlags(c, CreateCredential)
	defer timeTrack(time.Now(), "create network")

	oAuth2Client := NewOAuth2HTTPClient(c.String(FlServer.Name), c.String(FlUsername.Name), c.String(FlPassword.Name))

	networkName := c.String(FlNetworkName.Name)
	log.Infof("[CreateNetworkCommand] create network with name: %s", networkName)

	var vpcParams = make(map[string]interface{})
	vpcParams["vpcId"] = c.String(FlVPC.Name)
	vpcParams["internetGatewayId"] = c.String(FlIGW.Name)
	subnet := c.String(FlSubnet.Name)

	network := models.NetworkJSON{
		Name:          networkName,
		CloudPlatform: "AWS",
		Parameters:    vpcParams,
		SubnetCIDR:    &subnet,
	}

	resp, err := oAuth2Client.Cloudbreak.Networks.PostNetworksUser(&networks.PostNetworksUserParams{&network})

	if err != nil {
		logErrorAndExit(CreateNetworkCommand, err.Error())
	}

	log.Infof("[CreateNetworkCommand] network created, id: %d", resp.Payload.ID)
	return nil
}

func (c *Cloudbreak) GetNetwork(name string) models.NetworkJSON {
	log.Infof("[GetNetwork] sending get request to find network with name: %s", name)
	resp, err := c.Cloudbreak.Networks.GetNetworksAccountName(&networks.GetNetworksAccountNameParams{Name: name})

	if err != nil {
		logErrorAndExit(c.GetNetwork, err.Error())
	}

	defaultNetwork := *resp.Payload
	log.Infof("[GetNetwork] found network, name: %s", defaultNetwork.Name)
	return defaultNetwork
}

func (c *Cloudbreak) GetNetworkById(id int64) *models.NetworkJSON {
	log.Infof("[GetNetwork] sending get request to find network with id: %d", id)
	resp, err := c.Cloudbreak.Networks.GetNetworksID(&networks.GetNetworksIDParams{ID: id})

	if err != nil {
		logErrorAndExit(c.GetNetwork, err.Error())
	}

	network := resp.Payload
	log.Infof("[GetNetwork] found network, name: %s", network.Name)
	return network
}

func DeleteNetwork(c *cli.Context) error {
	checkRequiredFlags(c, DeleteCredential)
	oAuth2Client := NewOAuth2HTTPClient(c.String(FlServer.Name), c.String(FlUsername.Name), c.String(FlPassword.Name))
	if err := oAuth2Client.DeleteNetwork(c.String(FlNetworkName.Name)); err != nil {
		logErrorAndExit(DeleteNetwork, err.Error())
	}
	return nil
}

func (c *Cloudbreak) GetPublicNetworks() []*models.NetworkJSON {
	defer timeTrack(time.Now(), "get public networks")
	resp, err := c.Cloudbreak.Networks.GetNetworksAccount(&networks.GetNetworksAccountParams{})
	if err != nil {
		logErrorAndExit(c.GetPublicNetworks, err.Error())
	}
	return resp.Payload
}

func (c *Cloudbreak) DeleteNetwork(name string) error {
	defer timeTrack(time.Now(), "delete network")
	log.Infof("[DeleteNetwork] delete network: %s", name)
	return c.Cloudbreak.Networks.DeleteNetworksAccountName(&networks.DeleteNetworksAccountNameParams{Name: name})
}

func ListPrivateNetworks(c *cli.Context) error {
	checkRequiredFlags(c, CreateCredential)
	defer timeTrack(time.Now(), "list the private networks")

	oAuth2Client := NewOAuth2HTTPClient(c.String(FlServer.Name), c.String(FlUsername.Name), c.String(FlPassword.Name))
	networkResp := oAuth2Client.GetPrivateNetworks()

	var tableRows []Row
	for _, net := range networkResp {
		row := &NetworkList{Id: *net.ID, Name: net.Name}
		tableRows = append(tableRows, row)
	}
	output := Output{Format: c.String(FlOutput.Name)}
	output.WriteList(NetworkHeader, tableRows)
	return nil
}

func (c *Cloudbreak) GetPrivateNetworks() []*models.NetworkJSON {
	defer timeTrack(time.Now(), "get private networks")
	resp, err := c.Cloudbreak.Networks.GetNetworksUser(&networks.GetNetworksUserParams{})
	if err != nil {
		logErrorAndExit(c.GetPrivateNetworks, err.Error())
	}
	return resp.Payload
}
