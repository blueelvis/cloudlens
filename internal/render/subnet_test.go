package render

import (
	"testing"

	"github.com/one2nc/cloudlens/internal/aws"
	"github.com/stretchr/testify/assert"
)

func TestSubnetRender(t *testing.T) {
	pom := aws.SubnetResp{SubnetId: "subnet-1", OwnerId: "000000000000", CidrBlock: "172.31.0.0/16", AvailabilityZone: "us-east-1", State: "disabled"}

	var subnet Subnet
	r := NewRow(5)
	err := subnet.Render(pom, "subnet", &r)
	assert.Nil(t, err)

	assert.Equal(t, "subnet", r.ID)
	e := Fields{"subnet-1", "000000000000", "172.31.0.0/16", "us-east-1", "disabled"}
	assert.Equal(t, e, r.Fields[0:])
}
