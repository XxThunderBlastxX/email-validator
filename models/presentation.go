package models

type ResponseErr struct {
	Status bool   `json:"status"`
	Error  string `json:"error"`
}

type ResponseSuccess struct {
	Status        bool     `json:"status"`
	Error         string   `json:"error"`
	ContainsMx    bool     `json:"contains_mx"`
	MxRecord      []string `json:"mx_record"`
	ContainsSpf   bool     `json:"contains_spf"`
	Spf           string   `json:"spf"`
	ContainsDmarc bool     `json:"contains_dmarc"`
	Dmarc         string   `json:"dmarc"`
}

// PresentErr is a method to present error message
func (r *ResponseErr) PresentErr() *ResponseErr {
	return &ResponseErr{
		Status: false,
		Error:  r.Error,
	}
}

// PresentSuccess is method to present success response
func (r *ResponseSuccess) PresentSuccess() *ResponseSuccess {
	return &ResponseSuccess{
		Status:        true,
		Error:         "",
		ContainsMx:    r.ContainsMx,
		MxRecord:      r.MxRecord,
		ContainsSpf:   r.ContainsSpf,
		Spf:           r.Spf,
		ContainsDmarc: r.ContainsDmarc,
		Dmarc:         r.Dmarc,
	}
}
