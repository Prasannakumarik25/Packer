// Code generated by "hcl2-schema"; DO NOT EDIT.\n

package ecs

import (
	"github.com/hashicorp/hcl2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

func (*RunConfig) HCL2Spec() hcldec.ObjectSpec {
	s := map[string]hcldec.Spec{
		"AssociatePublicIpAddress": &hcldec.AttrSpec{Name:"AssociatePublicIpAddress", Type:cty.Bool, Required:false},
		"ZoneId": &hcldec.AttrSpec{Name:"ZoneId", Type:cty.String, Required:false},
		"IOOptimized": &hcldec.AttrSpec{Name:"IOOptimized", Type:cty.Bool, Required:false},
		"InstanceType": &hcldec.AttrSpec{Name:"InstanceType", Type:cty.String, Required:false},
		"Description": &hcldec.AttrSpec{Name:"Description", Type:cty.String, Required:false},
		"AlicloudSourceImage": &hcldec.AttrSpec{Name:"AlicloudSourceImage", Type:cty.String, Required:false},
		"ForceStopInstance": &hcldec.AttrSpec{Name:"ForceStopInstance", Type:cty.Bool, Required:false},
		"DisableStopInstance": &hcldec.AttrSpec{Name:"DisableStopInstance", Type:cty.Bool, Required:false},
		"SecurityGroupId": &hcldec.AttrSpec{Name:"SecurityGroupId", Type:cty.String, Required:false},
		"SecurityGroupName": &hcldec.AttrSpec{Name:"SecurityGroupName", Type:cty.String, Required:false},
		"UserData": &hcldec.AttrSpec{Name:"UserData", Type:cty.String, Required:false},
		"UserDataFile": &hcldec.AttrSpec{Name:"UserDataFile", Type:cty.String, Required:false},
		"VpcId": &hcldec.AttrSpec{Name:"VpcId", Type:cty.String, Required:false},
		"VpcName": &hcldec.AttrSpec{Name:"VpcName", Type:cty.String, Required:false},
		"CidrBlock": &hcldec.AttrSpec{Name:"CidrBlock", Type:cty.String, Required:false},
		"VSwitchId": &hcldec.AttrSpec{Name:"VSwitchId", Type:cty.String, Required:false},
		"VSwitchName": &hcldec.AttrSpec{Name:"VSwitchName", Type:cty.String, Required:false},
		"InstanceName": &hcldec.AttrSpec{Name:"InstanceName", Type:cty.String, Required:false},
		"InternetChargeType": &hcldec.AttrSpec{Name:"InternetChargeType", Type:cty.String, Required:false},
		"InternetMaxBandwidthOut": &hcldec.AttrSpec{Name:"InternetMaxBandwidthOut", Type:cty.Number, Required:false},
		"WaitSnapshotReadyTimeout": &hcldec.AttrSpec{Name:"WaitSnapshotReadyTimeout", Type:cty.Number, Required:false},
		"Comm": nil,
		"SSHPrivateIp": &hcldec.AttrSpec{Name:"SSHPrivateIp", Type:cty.Bool, Required:false},
	}
	return hcldec.ObjectSpec(s)
}
