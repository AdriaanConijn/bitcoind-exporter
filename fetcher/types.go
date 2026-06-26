package fetcher

import "encoding/json"

type StringOrArray []string

// UnmarshalJSON implements json.Unmarshaler interface
func (sa *StringOrArray) UnmarshalJSON(data []byte) error {
	var strings []string
	if err := json.Unmarshal(data, &strings); err == nil {
		*sa = strings
		return nil
	}

	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	*sa = []string{str}
	return nil
}

type BlockchainInfo struct {
	Chain                string        `json:"chain"`
	Blocks               int           `json:"blocks"`
	Headers              int           `json:"headers"`
	BestBlockhash        string        `json:"bestblockhash"`
	Difficulty           float64       `json:"difficulty"`
	Time                 int           `json:"time"`
	MedianTime           int           `json:"mediantime"`
	VerificationProgress float64       `json:"verificationprogress"`
	InitialBlockDownload bool          `json:"initialblockdownload"`
	ChainWork            string        `json:"chainwork"`
	SizeOnDisk           int           `json:"size_on_disk"`
	Pruned               bool          `json:"pruned"`
	Warnings             StringOrArray `json:"warnings"`
}

type MempoolInfo struct {
	Loaded              bool    `json:"loaded"`
	Size                int     `json:"size"`
	Bytes               int     `json:"bytes"`
	Usage               int     `json:"usage"`
	TotalFee            float64 `json:"total_fee"`
	MaxMempool          int     `json:"maxmempool"`
	MempoolMinFee       float64 `json:"mempoolminfee"`
	MinRelayTxFee       float64 `json:"minrelaytxfee"`
	IncrementalRelayFee float64 `json:"incrementalrelayfee"`
	UnbroadcastCount    int     `json:"unbroadcastcount"`
	FullRBF             bool    `json:"fullrbf"`
}

type MemoryInfo struct {
	Locked struct {
		Used       int `json:"used"`
		Free       int `json:"free"`
		Total      int `json:"total"`
		Locked     int `json:"locked"`
		ChunksUsed int `json:"chunks_used"`
		ChunksFree int `json:"chunks_free"`
	} `json:"locked"`
}

type IndexInfo struct {
	TxIndex struct {
		Synced          bool `json:"synced"`
		BestBlockHeight int  `json:"best_block_height"`
	}
}

type Network struct {
	Name                      string `json:"name"`
	Limited                   bool   `json:"limited"`
	Reachable                 bool   `json:"reachable"`
	Proxy                     string `json:"proxy"`
	ProxyRandomizeCredentials bool   `json:"proxy_randomize_credentials"`
}

type LocalAddress struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
	Score   int    `json:"score"`
}

type NetworkInfo struct {
	Version            int            `json:"version"`
	Subversion         string         `json:"subversion"`
	ProtocolVersion    int            `json:"protocolversion"`
	LocalServices      string         `json:"localservices"`
	LocalServicesNames []string       `json:"localservicesnames"`
	LocalRelay         bool           `json:"localrelay"`
	Timeoffset         int            `json:"timeoffset"`
	TotalConnections   int            `json:"connections"`
	ConnectionsIn      int            `json:"connections_in"`
	ConnectionsOut     int            `json:"connections_out"`
	Networks           []Network      `json:"networks"`
	RelayFee           float64        `json:"relayfee"`
	IncrementalFee     float64        `json:"incrementalfee"`
	LocalAddresses     []LocalAddress `json:"localaddresses"`
}

type SmartFee struct {
	Feerate float64 `json:"feerate"`
	Blocks  int     `json:"blocks"`
}

type NetTotals struct {
	TotalBytesRecv int `json:"totalbytesrecv"`
	TotalBytesSent int `json:"totalbytessent"`
	TimeMillis     int `json:"timemillis"`
}

type Peer struct {
	ID                int              `json:"id"`
	Addr              string           `json:"addr"`
	AddrBind          string           `json:"addrbind"`
	Network           string           `json:"network"`
	ConnectionType    string           `json:"connection_type"`
	Inbound           bool             `json:"inbound"`
	Services          string           `json:"services"`
	RelayTxes         bool             `json:"relaytxes"`
	Version           int              `json:"version"`
	SubVer            string           `json:"subver"`
	StartingHeight    int64            `json:"startingheight"`
	SyncedHeaders     int64            `json:"synced_headers"`
	SyncedBlocks      int64            `json:"synced_blocks"`
	BytesSent         int64            `json:"bytessent"`
	BytesRecv         int64            `json:"bytesrecv"`
	BytesSentPerMsg   map[string]int64 `json:"bytessent_per_msg"`
	BytesRecvPerMsg   map[string]int64 `json:"bytesrecv_per_msg"`
	PingTime          float64          `json:"pingtime"`
	MinPing           float64          `json:"minping"`
	ConnTime          int64            `json:"conntime"`
	LastSend          int64            `json:"lastsend"`
	LastRecv          int64            `json:"lastrecv"`
	TransportProtocol string           `json:"transport_protocol_type"`
	SessionID         string           `json:"session_id"`
}
