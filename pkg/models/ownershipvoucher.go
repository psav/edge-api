package models

// FDODevice has a one OwnershipVoucherData and a one FDO User
type FDODevice struct {
	Model
	OwnershipVoucherData            *OwnershipVoucherData `json:"ownership_voucher_data"`
	Connected                       bool                  `json:"connected" gorm:"default:false"`
	UUID                            string                `json:"uuid"`
	SubscriptionIdentityCertificate string                `json:"subscription_identity_certificate"`
	InitialUser                     *FDOUser              `json:"initial_user"`
}

// OwnershipVoucherData represents the data of an ownership voucher
type OwnershipVoucherData struct {
	Model
	ProtocolVersion uint   `json:"protocol_version"`
	GUID            string `json:"guid"`
	DeviceName      string `json:"device_name"`
	FDODeviceID     uint   `json:"fdo_device_id"`
}

// FDOUser represents an initial user of FDO, FDOUser has many SSH keys
type FDOUser struct {
	Model
	Username    string   `json:"username"`
	SSHKeys     []SSHKey `json:"ssh_keys"`
	FDODeviceID uint     `json:"fdo_device_id"`
}

// SSHKey represents an SSH key of a user
type SSHKey struct {
	Model
	Key       string `json:"key"`
	FDOUserID uint   `json:"fdo_user_id"`
}