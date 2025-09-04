package packlit

// StreamDescriptor represents a single stream descriptor in JSON.
// It closely maps to a descriptor built by WithInput/WithStream/WithOutput and other options.
type StreamDescriptor struct {
	Input           string            `json:"input"`
	Stream          string            `json:"stream"`
	Output          string            `json:"output"`
	InitSegment     string            `json:"init_segment,omitempty"`
	SegmentTemplate string            `json:"segment_template,omitempty"`
	Extra           map[string]string `json:"extra,omitempty"`
}

// PackagerOptions is the top-level JSON payload accepted by parser API.
type PackagerOptions struct {
	Binary     string             `json:"binary,omitempty"`
	Streams    []StreamDescriptor `json:"streams"`
	Packager   *PackagerFlags     `json:"packager,omitempty"`
	Muxer      *MuxerFlags        `json:"muxer,omitempty"`
	MPD        *MpdFlags          `json:"mpd,omitempty"`
	Manifest   *ManifestFlags     `json:"manifest,omitempty"`
	HTTP       *HttpFlags         `json:"http,omitempty"`
	HLS        *HlsFlags          `json:"hls,omitempty"`
	Crypto     *CryptoFlags       `json:"crypto,omitempty"`
	Protection *ProtectionFlags   `json:"protection,omitempty"`
	RawKey     *RawKeyFlags       `json:"raw_key,omitempty"`
	Widevine   *WideVineFlags     `json:"widevine,omitempty"`
	Playready  *PlayReadyFlags    `json:"playready,omitempty"`
}

// Flag configuration structs
type PackagerFlags struct {
	Licenses             bool   `json:"licenses,omitempty"`
	Quiet                bool   `json:"quiet,omitempty"`
	UseFakeClockForMuxer bool   `json:"use_fake_clock_for_muxer,omitempty"`
	TestPackagerVersion  bool   `json:"test_packager_version,omitempty"`
	SingleThreaded       bool   `json:"single_threaded,omitempty"`
	VModule              string `json:"vmodule,omitempty"`
	Version              bool   `json:"version,omitempty"`
}

type MuxerFlags struct {
	ClearLead                        float32 `json:"clear_lead,omitempty"`
	SegmentDuration                  float32 `json:"segment_duration,omitempty"`
	SegmentSapAligned                bool    `json:"segment_sap_aligned,omitempty"`
	FragmentDuration                 float32 `json:"fragment_duration,omitempty"`
	FragmentSapAligned               bool    `json:"fragment_sap_aligned,omitempty"`
	GenerateSidxInMediaSegments      bool    `json:"generate_sidx_in_media_segments,omitempty"`
	TempDir                          string  `json:"temp_dir,omitempty"`
	Mp4IncludePsshInStream           bool    `json:"mp4_include_pssh_in_stream,omitempty"`
	TransportStreamTimestampOffsetMs int     `json:"transport_stream_timestamp_offset_ms,omitempty"`
	DefaultTextZeroBiasMs            int     `json:"default_text_zero_bias_ms,omitempty"`
	StartSegmentNumber               int     `json:"start_segment_number,omitempty"`
	UseDoviSupplementalCodecs        bool    `json:"use_dovi_supplemental_codecs,omitempty"`
}

type MpdFlags struct {
	OutputMediaInfo                 bool     `json:"output_media_info,omitempty"`
	MpdOutput                       string   `json:"mpd_output,omitempty"`
	BaseUrls                        []string `json:"base_urls,omitempty"`
	GenerateStaticLiveMpd           bool     `json:"generate_static_live_mpd,omitempty"`
	AllowApproximateSegmentTimeline bool     `json:"allow_approximate_segment_timeline,omitempty"`
	DumpStreamInfo                  bool     `json:"dump_stream_info,omitempty"`
	MinBufferTime                   float32  `json:"min_buffer_time,omitempty"`
	MinimumUpdatePeriod             float32  `json:"minimum_update_period,omitempty"`
	SuggestedPresentationDelay      float32  `json:"suggested_presentation_delay,omitempty"`
	UtcTimings                      []string `json:"utc_timings,omitempty"`
	GenerateDashIfIopCompliantMpd   bool     `json:"generate_dash_if_iop_compliant_mpd,omitempty"`
	AllowCodecSwitching             bool     `json:"allow_codec_switching,omitempty"`
	IncludeMsprProForPlayReady      bool     `json:"include_mspr_pro_for_playready,omitempty"`
	DashForceSegmentList            bool     `json:"dash_force_segment_list,omitempty"`
	LowLatencyDashMode              bool     `json:"low_latency_dash_mode,omitempty"`
}

type ManifestFlags struct {
	TimeShiftBufferDepth               float32 `json:"time_shift_buffer_depth,omitempty"`
	PreservedSegmentsOutsideLiveWindow int     `json:"preserved_segments_outside_live_window,omitempty"`
	DefaultLanguage                    string  `json:"default_language,omitempty"`
	DefaultTextLanguage                string  `json:"default_text_language,omitempty"`
	ForceClIndex                       bool    `json:"force_cl_index,omitempty"`
}

type HttpFlags struct {
	UserAgent                    string `json:"user_agent,omitempty"`
	CaFile                       string `json:"ca_file,omitempty"`
	ClientCertFile               string `json:"client_cert_file,omitempty"`
	ClientCertPrivateKeyFile     string `json:"client_cert_private_key_file,omitempty"`
	ClientCertPrivateKeyPassword string `json:"client_cert_private_key_password,omitempty"`
	DisablePeerVerification      bool   `json:"disable_peer_verification,omitempty"`
	IgnoreHttpOutputFailures     bool   `json:"ignore_http_output_failures,omitempty"`
	IoCacheSize                  uint   `json:"io_cache_size,omitempty"`
}

type HlsFlags struct {
	MasterPlaylistOutput string  `json:"hls_master_playlist_output,omitempty"`
	BaseUrl              string  `json:"hls_base_url,omitempty"`
	KeyUri               string  `json:"hls_key_uri,omitempty"`
	PlaylistType         string  `json:"hls_playlist_type,omitempty"`
	MediaSequenceNumber  int     `json:"hls_media_sequence_number,omitempty"`
	StartTimeOffset      float32 `json:"hls_start_time_offset,omitempty"`
	CreateSessionKeys    bool    `json:"create_session_keys,omitempty"`
}

type CryptoFlags struct {
	CryptByteBlock           int    `json:"crypt_byte_block,omitempty"`
	ProtectionScheme         string `json:"protection_scheme,omitempty"`
	SkipByteBlock            int    `json:"skip_byte_block,omitempty"`
	Vp9SubsampleEncryption   bool   `json:"vp9_subsample_encryption,omitempty"`
	PlayreadyExtraHeaderData string `json:"playready_extra_header_data,omitempty"`
}

type ProtectionFlags struct {
	ProtectionSystems []string `json:"protection_systems,omitempty"`
}

type RawKeyFlags struct {
	EnableEncryption bool     `json:"enable_raw_key_encryption,omitempty"`
	EnableDecryption bool     `json:"enable_raw_key_decryption,omitempty"`
	Keys             []string `json:"keys,omitempty"`
	Iv               string   `json:"iv,omitempty"`
	Pssh             string   `json:"pssh,omitempty"`
}

type WideVineFlags struct {
	EnableEncryption     bool   `json:"enable_widevine_encryption,omitempty"`
	EnableDecryption     bool   `json:"enable_widevine_decryption,omitempty"`
	KeyServerUrl         string `json:"key_server_url,omitempty"`
	ContentId            string `json:"content_id,omitempty"`
	Policy               string `json:"policy,omitempty"`
	MaxSdPixels          int    `json:"max_sd_pixels,omitempty"`
	MaxHdPixels          int    `json:"max_hd_pixels,omitempty"`
	MaxUhd1Pixels        int    `json:"max_uhd1_pixels,omitempty"`
	Signer               string `json:"signer,omitempty"`
	AesSigningKey        string `json:"aes_signing_key,omitempty"`
	AesSigningIv         string `json:"aes_signing_iv,omitempty"`
	RsaSigningKeyPath    string `json:"rsa_signing_key_path,omitempty"`
	CryptoPeriodDuration int    `json:"crypto_period_duration,omitempty"`
	GroupId              string `json:"group_id,omitempty"`
	EnableEntitlement    bool   `json:"enable_entitlement_license,omitempty"`
}

type PlayReadyFlags struct {
	EnableEncryption  bool   `json:"enable_playready_encryption,omitempty"`
	ServerUrl         string `json:"playready_server_url,omitempty"`
	ProgramIdentifier string `json:"program_identifier,omitempty"`
}
