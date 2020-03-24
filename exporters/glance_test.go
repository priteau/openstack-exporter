package exporters

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
)

type GlanceTestSuite struct {
	BaseOpenStackTestSuite
}

var glanceExpectedUp = `
# HELP openstack_glance_image_size image_size
# TYPE openstack_glance_image_size gauge
openstack_glance_image_size{name="F17-x86_64-cfntools",resource_id="781b3762-9469-4cec-b58d-3349e5de4e9c",tenant_id="5ef70662f8b34079a6eddb8da9d75fe8"} 4.76704768e+08
openstack_glance_image_size{name="cirros-0.3.2-x86_64-disk",resource_id="1bea47ed-f6a9-463b-b423-14b9cca9ad27",tenant_id="5ef70662f8b34079a6eddb8da9d75fe8"} 1.3167616e+07
# HELP openstack_glance_images images
# TYPE openstack_glance_images gauge
openstack_glance_images 2
# HELP openstack_glance_up up
# TYPE openstack_glance_up gauge
openstack_glance_up 1
`

var glanceExpectedDown = `
# HELP openstack_glance_up up
# TYPE openstack_glance_up gauge
openstack_glance_up 0
`

func (suite *GlanceTestSuite) TestGlanceExporter() {
	err := testutil.CollectAndCompare(*suite.Exporter, strings.NewReader(glanceExpectedUp))
	assert.NoError(suite.T(), err)
}

func (suite *GlanceTestSuite) TestGlanceExporterWithEndpointDown() {
	suite.teardownFixtures()
	defer suite.installFixtures()

	err := testutil.CollectAndCompare(*suite.Exporter, strings.NewReader(glanceExpectedDown))
	assert.NoError(suite.T(), err)
}
