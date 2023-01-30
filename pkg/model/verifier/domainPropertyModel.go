package verifier

type DomainPropertiesData struct {
	Mx           bool
	Spf          bool
	Dmarc        bool
	SpfRecords   string
	DmarcRecords string
}
