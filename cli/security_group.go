package cli

import (
	log "github.com/Sirupsen/logrus"
	"github.com/hortonworks/hdc-cli/client/securitygroups"
	"github.com/hortonworks/hdc-cli/models"
	"strconv"
	"strings"
	"sync"
	"time"
)

var SECURITY_GROUP_DEFAULT_PORTS = []string{"22", "443", "9443"}

func (c *Cloudbreak) CreateSecurityGroup(skeleton ClusterSkeleton, channel chan int64, wg *sync.WaitGroup) {
	defer timeTrack(time.Now(), "create security group")

	secGroupName := "secg" + strconv.FormatInt(time.Now().UnixNano(), 10)

	defaultPorts := SECURITY_GROUP_DEFAULT_PORTS
	if skeleton.WebAccess {
		defaultPorts = append(defaultPorts, "8080", "18080", "18081", "9995", "6080", "21000", "10000")
	}

	modifiable := false
	secRules := []*models.SecurityRule{
		{
			Subnet:     skeleton.RemoteAccess,
			Protocol:   "tcp",
			Ports:      strings.Join(defaultPorts, ","),
			Modifiable: &modifiable,
		},
	}

	secGroup := models.SecurityGroupJSON{
		Name:          secGroupName,
		SecurityRules: secRules,
	}

	log.Infof("[CreateSecurityGroup] sending security group create request with name: %s", secGroupName)
	resp, err := c.Cloudbreak.Securitygroups.PostSecuritygroupsAccount(&securitygroups.PostSecuritygroupsAccountParams{&secGroup})

	if err != nil {
		logErrorAndExit(c.CreateSecurityGroup, err.Error())
	}

	log.Infof("[CreateSecurityGroup] security group created, id: %d", resp.Payload.ID)
	channel <- resp.Payload.ID

	wg.Done()
}

func (c *Cloudbreak) GetSecurityDetails(stack *models.StackResponse) (securityMap map[string][]*models.SecurityRule, err error) {
	defer timeTrack(time.Now(), "get security group by id")
	securityMap = make(map[string][]*models.SecurityRule)
	for _, v := range stack.InstanceGroups {
		if respSecurityGroup, err := c.Cloudbreak.Securitygroups.GetSecuritygroupsID(&securitygroups.GetSecuritygroupsIDParams{ID: v.SecurityGroupID}); err == nil {
			securityGroup := respSecurityGroup.Payload
			for _, sr := range securityGroup.SecurityRules {
				securityMap[sr.Subnet] = append(securityMap[sr.Subnet], sr)
			}
		}
	}
	return securityMap, err
}

func (c *Cloudbreak) GetPublicSecurityGroups() []*models.SecurityGroupJSON {
	defer timeTrack(time.Now(), "get public security groups")
	resp, err := c.Cloudbreak.Securitygroups.GetSecuritygroupsAccount(&securitygroups.GetSecuritygroupsAccountParams{})
	if err != nil {
		logErrorAndExit(c.GetPublicSecurityGroups, err.Error())
	}
	return resp.Payload
}

func (c *Cloudbreak) DeleteSecurityGroup(name string) error {
	defer timeTrack(time.Now(), "delete security group")
	log.Infof("[DeleteSecurityGroup] delete security group: %s", name)
	return c.Cloudbreak.Securitygroups.DeleteSecuritygroupsAccountName(&securitygroups.DeleteSecuritygroupsAccountNameParams{Name: name})
}