package iec61850

type IEC61850Model struct {
	Name       string
	Datasets   []*Dataset
	RCBs       []*ReportControlBlock
	FirstChild *LogicalDevice
}
