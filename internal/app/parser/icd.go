package parser

import "encoding/xml"

type SCL struct {
	XMLName xml.Name `xml:"SCL"`
	Header  struct {
		ID            string `xml:"id,attr"`
		NameStructure string `xml:"nameStructure,attr"`
	} `xml:"Header"`
	IED               IED               `xml:"IED"`
	DataTypeTemplates DataTypeTemplates `xml:"DataTypeTemplates"`
}

type IED struct {
	Name         string        `xml:"name,attr"`
	Description  string        `xml:"desc,attr"`
	Manufacturer string        `xml:"manufacturer,attr"`
	Services     Services      `xml:"Services"`
	AccessPoint  []AccessPoint `xml:"AccessPoint"`
}

type Services struct {
	DynAssociation          struct{} `xml:"DynAssociation"`
	GetDirectory            struct{} `xml:"GetDirectory"`
	GetDataObjectDefinition struct{} `xml:"GetDataObjectDefinition"`
	GetDataSetValue         struct{} `xml:"GetDataSetValue"`
	DataSetDirectory        struct{} `xml:"DataSetDirectory"`
	ReadWrite               struct{} `xml:"ReadWrite"`
	GetCBValues             struct{} `xml:"GetCBValues"`
	ConfLNs                 struct {
		FixPrefix bool `xml:"fixPrefix,attr"`
		FixLnInst bool `xml:"fixLnInst,attr"`
	} `xml:"ConfLNs"`
	GOOSE struct {
		Max string `xml:"max,attr"`
	} `xml:"GOOSE"`
	GSSE struct {
		Max string `xml:"max,attr"`
	} `xml:"GSSE"`
	FileHandling          struct{} `xml:"FileHandling"`
	GSEDir                struct{} `xml:"GSEDir"`
	TimerActivatedControl struct{} `xml:"TimerActivatedControl"`
}

type AccessPoint struct {
	Name   string `xml:"name,attr"`
	Server Server `xml:"Server"`
}

type Server struct {
	Authentication struct{}  `xml:"Authentication"`
	LDevice        []LDevice `xml:"LDevice"`
}

type LDevice struct {
	Inst string `xml:"inst,attr"`
	LN0  LN0    `xml:"LN0"`
	LN   []LN   `xml:"LN"`
}

type LN0 struct {
	LnClass       string          `xml:"lnClass,attr"`
	LnType        string          `xml:"lnType,attr"`
	Inst          string          `xml:"inst,attr"`
	Prefix        string          `xml:"prefix,attr"`
	DataSet       []DataSet       `xml:"DataSet"`
	ReportControl []ReportControl `xml:"ReportControl"`
	DOI           []DOI           `xml:"DOI"`
}

type LN struct {
	LnClass string `xml:"lnClass,attr"`
	LnType  string `xml:"lnType,attr"`
	Inst    string `xml:"inst,attr"`
	Prefix  string `xml:"prefix,attr"`
	DOI     []DOI  `xml:"DOI"`
}

type DataSet struct {
	Name string `xml:"name,attr"`
	Desc string `xml:"desc,attr"`
	FCDA []FCDA `xml:"FCDA"`
}

type ReportControl struct {
	Name       string     `xml:"name,attr"`
	ConfRev    string     `xml:"confRev,attr"`
	DatSet     string     `xml:"datSet,attr"`
	RptID      string     `xml:"rptID,attr"`
	Buffered   bool       `xml:"buffered,attr"`
	IntgPd     string     `xml:"intgPd,attr"`
	BufTime    string     `xml:"bufTime,attr"`
	TrgOps     TrgOps     `xml:"TrgOps"`
	OptFields  OptFields  `xml:"OptFields"`
	RptEnabled RptEnabled `xml:"RptEnabled"`
	Indexed    bool       `xml:"indexed,attr"`
}

type FCDA struct {
	LdInst  string `xml:"ldInst,attr"`
	LnClass string `xml:"lnClass,attr"`
	FC      string `xml:"fc,attr"`
	LnInst  string `xml:"lnInst,attr"`
	DoName  string `xml:"doName,attr"`
	DaName  string `xml:"daName,attr"`
}

type TrgOps struct {
	Period bool `xml:"period,attr"`
}

type OptFields struct {
	SeqNum     bool `xml:"seqNum,attr"`
	TimeStamp  bool `xml:"timeStamp,attr"`
	DataSet    bool `xml:"dataSet,attr"`
	ReasonCode bool `xml:"reasonCode,attr"`
	ConfigRef  bool `xml:"configRef,attr"`
	EntryID    bool `xml:"entryID,attr"`
}

type RptEnabled struct {
	Max string `xml:"max,attr"`
}

type DOI struct {
	Name string `xml:"name,attr"`
	DAI  []DAI  `xml:"DAI"`
}

type DAI struct {
	Name string `xml:"name,attr"`
	Val  string `xml:"Val"`
}

type DataTypeTemplates struct {
	LNodeType []LNodeType `xml:"LNodeType"`
	DOType    []DOType    `xml:"DOType"`
	DAType    []DAType    `xml:"DAType"`
	EnumType  []EnumType  `xml:"EnumType"`
}

type LNodeType struct {
	ID      string `xml:"id,attr"`
	LnClass string `xml:"lnClass,attr"`
	DO      []DO   `xml:"DO"`
}

type DO struct {
	Name string `xml:"name,attr"`
	Type string `xml:"type,attr"`
}

type DOType struct {
	ID  string `xml:"id,attr"`
	CDC string `xml:"cdc,attr"`
	DA  []DA   `xml:"DA"`
}

type DA struct {
	Name  string `xml:"name,attr"`
	BType string `xml:"bType,attr"`
	Type  string `xml:"type,attr"`
	FC    string `xml:"fc,attr"`
	Dchg  bool   `xml:"dchg,attr"`
	Qchg  bool   `xml:"qchg,attr"`
}

type DAType struct {
	ID  string `xml:"id,attr"`
	BDA []BDA  `xml:"BDA"`
}

type BDA struct {
	Name  string `xml:"name,attr"`
	BType string `xml:"bType,attr"`
	Type  string `xml:"type,attr"`
}

type EnumType struct {
	ID      string    `xml:"id,attr"`
	EnumVal []EnumVal `xml:"EnumVal"`
}

type EnumVal struct {
	Ord string `xml:"ord,attr"`
	Val string `xml:"#text"`
}
