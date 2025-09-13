package ipfs

type Credentials struct {
	Name		 	string 			`json:"name"`
	Description 	string 			`json:"description"`
	Image 			[]byte 			`json:"image"`
	ExternalURL 	string  		`json:"external_url"`
	Attributes  	[]Attribute 	`json:"attributes"`
	AnimationURL 	string 			`json:"animation_url"`
	YoutubeURL 	 	string 			`json:"youtube_url"`
	BgColor 	 	string 			`json:"background_color"`
	CustomFields 	[]CustomField 	`json:"custom_field"`
}

type Attribute struct {
	TraitType  string 	`json:"trait_type"`
	Value	   string 	`json:"value"`
}

type CustomField struct {
	DeanSignatureHash	string  `json:"dean_signature_hash"`
}