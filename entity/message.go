package entity

type ZlmResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type OnPlayReq struct {
	MediaServerId string `json:"mediaServerId"`
	App           string `json:"app"`
	Id            string `json:"id"`
	Ip            string `json:"ip"`
	Params        string `json:"params"`
	Port          int    `json:"port"`
	Schema        string `json:"schema"`
	Stream        string `json:"stream"`
	Vhost         string `json:"vhost"`
}

type OnStreamNotFoundReq struct {
	MediaServerId string `json:"mediaServerId"`
	App           string `json:"app"`
	Id            string `json:"id"`
	Ip            string `json:"ip"`
	Params        string `json:"params"`
	Port          int    `json:"port"`
	Schema        string `json:"schema"`
	Stream        string `json:"stream"`
	Vhost         string `json:"vhost"`
}

type OnPublishReq struct {
	MediaServerId string `json:"mediaServerId"`
	App           string `json:"app"`
	Id            string `json:"id"`
	Ip            string `json:"ip"`
	Params        string `json:"params"`
	Port          int    `json:"port"`
	Schema        string `json:"schema"`
	Stream        string `json:"stream"`
	Vhost         string `json:"vhost"`
}

type OnPublishResp struct {
	Code           int    `json:"code"`
	EnableHls      bool   `json:"enable_hls,omitempty"`
	EnableMp4      bool   `json:"enable_mp4,omitempty"`
	EnableRtmp     bool   `json:"enable_rtmp,omitempty"`
	EnableRtsp     bool   `json:"enable_rtsp,omitempty"`
	EnableTs       bool   `json:"enable_ts,omitempty"`
	EnableFmp4     bool   `json:"enable_fmp4,omitempty"`
	EnableAudio    bool   `json:"enable_audio,omitempty"`
	AddMuteAudio   bool   `json:"add_mute_audio,omitempty"`
	Mp4SavePath    string `json:"mp4_save_path,omitempty"`
	Mp4MaxSecond   int    `json:"mp4_max_second,omitempty"`
	HlsSavePath    string `json:"hls_save_path,omitempty"`
	ContinuePushMs uint32 `json:"continue_push_ms,omitempty"`
	Msg            string `json:"msg"`
}
type OnRecordMp4Req struct {
	MediaServerId string `json:"mediaServerId"`
	App           string `json:"app"`
	FileName      string `json:"file_name"`
	FilePath      string `json:"file_path"`
	FileSize      int    `json:"file_size"`
	Folder        string `json:"folder"`
	StartTime     int    `json:"start_time"`
	Stream        string `json:"stream"`
	TimeLen       int    `json:"time_len"`
	Url           string `json:"url"`
	Vhost         string `json:"vhost"`
}

type OnRtspRealmReq struct {
	MediaServerId string `json:"mediaServerId"`
	App           string `json:"app"`
	Id            string `json:"id"`
	Ip            string `json:"ip"`
	Params        string `json:"params"`
	Port          int    `json:"port"`
	Schema        string `json:"schema"`
	Stream        string `json:"stream"`
	Vhost         string `json:"vhost"`
}

type OnRtspRealmResp struct {
	Code  int    `json:"code"`
	Realm string `json:"realm"`
	Msg   string `json:"msg"`
}
type OnRtspAuthReq struct {
	MediaServerId string `json:"mediaServerId"`
	App           string `json:"app"`
	Id            string `json:"id"`
	Ip            string `json:"ip"`
	MustNoEncrypt bool   `json:"must_no_encrypt"`
	Params        string `json:"params"`
	Port          int    `json:"port"`
	Realm         string `json:"realm"`
	Schema        string `json:"schema"`
	Stream        string `json:"stream"`
	UserName      string `json:"user_name"`
	Vhost         string `json:"vhost"`
}
type OnRtspAuthResp struct {
	Code      int    `json:"code"`
	Encrypted bool   `json:"encrypted"`
	Passwd    string `json:"passwd"`
	Msg       string `json:"msg"`
}

type OnStreamNoneReaderReq struct {
	MediaServerId string `json:"mediaServerId"`
	App           string `json:"app"`
	Schema        string `json:"schema"`
	Stream        string `json:"stream"`
	Vhost         string `json:"vhost"`
}

type OnStreamNoneReaderResp struct {
	Close bool `json:"close"`
	Code  int  `json:"code"`
}

type OnServerKeepaliveReq struct {
	Data          OnServerKeepaliveData `json:"data"`
	MediaServerId string                `json:"mediaServerId"`
}
type OnServerKeepaliveData struct {
	Buffer                int `json:"Buffer"`
	BufferLikeString      int `json:"BufferLikeString"`
	BufferList            int `json:"BufferList"`
	BufferRaw             int `json:"BufferRaw"`
	Frame                 int `json:"Frame"`
	FrameImp              int `json:"FrameImp"`
	MediaSource           int `json:"MediaSource"`
	MultiMediaSourceMuxer int `json:"MultiMediaSourceMuxer"`
	RtmpPacket            int `json:"RtmpPacket"`
	RtpPacket             int `json:"RtpPacket"`
	Socket                int `json:"Socket"`
	TcpClient             int `json:"TcpClient"`
	TcpServer             int `json:"TcpServer"`
	TcpSession            int `json:"TcpSession"`
	UdpServer             int `json:"UdpServer"`
	UdpSession            int `json:"UdpSession"`
}
