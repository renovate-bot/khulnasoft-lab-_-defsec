package compute

import (
	"github.com/khulnasoft-lab/defsec/pkg/providers/google/compute"
	"github.com/khulnasoft-lab/defsec/pkg/terraform"
	defsecTypes "github.com/khulnasoft-lab/defsec/pkg/types"
	"github.com/zclconf/go-cty/cty"
)

func adaptInstances(modules terraform.Modules) (instances []compute.Instance) {

	for _, instanceBlock := range modules.GetResourcesByType("google_compute_instance") {

		instance := compute.Instance{
			Metadata: instanceBlock.GetMetadata(),
			Name:     instanceBlock.GetAttribute("name").AsStringValueOrDefault("", instanceBlock),
			ShieldedVM: compute.ShieldedVMConfig{
				Metadata:                   instanceBlock.GetMetadata(),
				SecureBootEnabled:          defsecTypes.BoolDefault(false, instanceBlock.GetMetadata()),
				IntegrityMonitoringEnabled: defsecTypes.BoolDefault(false, instanceBlock.GetMetadata()),
				VTPMEnabled:                defsecTypes.BoolDefault(false, instanceBlock.GetMetadata()),
			},
			ServiceAccount: compute.ServiceAccount{
				Metadata:  instanceBlock.GetMetadata(),
				Email:     defsecTypes.StringDefault("", instanceBlock.GetMetadata()),
				IsDefault: defsecTypes.BoolDefault(false, instanceBlock.GetMetadata()),
				Scopes:    nil,
			},
			CanIPForward:                instanceBlock.GetAttribute("can_ip_forward").AsBoolValueOrDefault(false, instanceBlock),
			OSLoginEnabled:              defsecTypes.BoolDefault(true, instanceBlock.GetMetadata()),
			EnableProjectSSHKeyBlocking: defsecTypes.BoolDefault(false, instanceBlock.GetMetadata()),
			EnableSerialPort:            defsecTypes.BoolDefault(false, instanceBlock.GetMetadata()),
			NetworkInterfaces:           nil,
			BootDisks:                   nil,
			AttachedDisks:               nil,
		}

		// network interfaces
		for _, networkInterfaceBlock := range instanceBlock.GetBlocks("network_interface") {
			ni := compute.NetworkInterface{
				Metadata:    networkInterfaceBlock.GetMetadata(),
				Network:     nil,
				SubNetwork:  nil,
				HasPublicIP: defsecTypes.BoolDefault(false, networkInterfaceBlock.GetMetadata()),
				NATIP:       defsecTypes.StringDefault("", networkInterfaceBlock.GetMetadata()),
			}
			if accessConfigBlock := networkInterfaceBlock.GetBlock("access_config"); accessConfigBlock.IsNotNil() {
				ni.HasPublicIP = defsecTypes.Bool(true, accessConfigBlock.GetMetadata())
			}
			instance.NetworkInterfaces = append(instance.NetworkInterfaces, ni)
		}

		// vm shielding
		if shieldedBlock := instanceBlock.GetBlock("shielded_instance_config"); shieldedBlock.IsNotNil() {
			instance.ShieldedVM.Metadata = shieldedBlock.GetMetadata()
			instance.ShieldedVM.IntegrityMonitoringEnabled = shieldedBlock.GetAttribute("enable_integrity_monitoring").AsBoolValueOrDefault(true, shieldedBlock)
			instance.ShieldedVM.VTPMEnabled = shieldedBlock.GetAttribute("enable_vtpm").AsBoolValueOrDefault(true, shieldedBlock)
			instance.ShieldedVM.SecureBootEnabled = shieldedBlock.GetAttribute("enable_secure_boot").AsBoolValueOrDefault(false, shieldedBlock)
		}

		// metadata
		if metadataAttr := instanceBlock.GetAttribute("metadata"); metadataAttr.IsNotNil() {
			if val := metadataAttr.MapValue("enable-oslogin"); val.Type() == cty.Bool {
				instance.OSLoginEnabled = defsecTypes.BoolExplicit(val.True(), metadataAttr.GetMetadata())
			}
			if val := metadataAttr.MapValue("block-project-ssh-keys"); val.Type() == cty.Bool {
				instance.EnableProjectSSHKeyBlocking = defsecTypes.BoolExplicit(val.True(), metadataAttr.GetMetadata())
			}
			if val := metadataAttr.MapValue("serial-port-enable"); val.Type() == cty.Bool {
				instance.EnableSerialPort = defsecTypes.BoolExplicit(val.True(), metadataAttr.GetMetadata())
			}
		}

		// disks
		for _, diskBlock := range instanceBlock.GetBlocks("boot_disk") {
			disk := compute.Disk{
				Metadata: diskBlock.GetMetadata(),
				Name:     diskBlock.GetAttribute("device_name").AsStringValueOrDefault("", diskBlock),
				Encryption: compute.DiskEncryption{
					Metadata:   diskBlock.GetMetadata(),
					RawKey:     diskBlock.GetAttribute("disk_encryption_key_raw").AsBytesValueOrDefault(nil, diskBlock),
					KMSKeyLink: diskBlock.GetAttribute("kms_key_self_link").AsStringValueOrDefault("", diskBlock),
				},
			}
			instance.BootDisks = append(instance.BootDisks, disk)
		}
		for _, diskBlock := range instanceBlock.GetBlocks("attached_disk") {
			disk := compute.Disk{
				Metadata: diskBlock.GetMetadata(),
				Name:     diskBlock.GetAttribute("device_name").AsStringValueOrDefault("", diskBlock),
				Encryption: compute.DiskEncryption{
					Metadata:   diskBlock.GetMetadata(),
					RawKey:     diskBlock.GetAttribute("disk_encryption_key_raw").AsBytesValueOrDefault(nil, diskBlock),
					KMSKeyLink: diskBlock.GetAttribute("kms_key_self_link").AsStringValueOrDefault("", diskBlock),
				},
			}
			instance.AttachedDisks = append(instance.AttachedDisks, disk)
		}

		if serviceAccountBlock := instanceBlock.GetBlock("service_account"); serviceAccountBlock.IsNotNil() {
			emailAttr := serviceAccountBlock.GetAttribute("email")
			instance.ServiceAccount.Email = emailAttr.AsStringValueOrDefault("", serviceAccountBlock)

			if instance.ServiceAccount.Email.IsEmpty() || instance.ServiceAccount.Email.EndsWith("-compute@developer.gserviceaccount.com") {
				instance.ServiceAccount.IsDefault = defsecTypes.Bool(true, serviceAccountBlock.GetMetadata())
			}

			if emailAttr.IsResourceBlockReference("google_service_account") {
				if accBlock, err := modules.GetReferencedBlock(emailAttr, instanceBlock); err == nil {
					instance.ServiceAccount.IsDefault = defsecTypes.Bool(false, serviceAccountBlock.GetMetadata())
					instance.ServiceAccount.Email = accBlock.GetAttribute("email").AsStringValueOrDefault("", accBlock)
				}
			}

			if scopesAttr := serviceAccountBlock.GetAttribute("scopes"); scopesAttr.IsNotNil() {
				instance.ServiceAccount.Scopes = append(instance.ServiceAccount.Scopes, scopesAttr.AsStringValues()...)
			}
		}

		instances = append(instances, instance)
	}

	return instances
}
