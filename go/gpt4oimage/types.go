package gpt4oimage

type TaskStatus string

type GenerationParams struct {
	Model          string   `json:"model" help:"required; gpt4o-image"`
	Prompt         string   `json:"prompt,omitempty" help:"optional; prompt text, required when files_url and file_url are blank"`
	Size           string   `json:"size" help:"required; 1:1, 3:2, or 2:3"`
	FilesURL       []string `json:"files_url,omitempty" help:"optional; up to 5 image URLs for edits or variants"`
	MaskURL        string   `json:"mask_url,omitempty" help:"optional; mask image URL"`
	NVariants      int      `json:"n_variants,omitempty" help:"optional; 1, 2, or 4 output variants"`
	CallbackURL    string   `json:"callback_url,omitempty" help:"optional; webhook URL"`
	IsEnhance      *bool    `json:"is_enhance,omitempty" help:"optional; prompt enhancement toggle"`
	UploadCN       *bool    `json:"upload_cn,omitempty" help:"optional; China upload routing toggle"`
	EnableFallback *bool    `json:"enable_fallback,omitempty" help:"optional; backup route toggle"`
	FallbackModel  string   `json:"fallback_model,omitempty" help:"optional; GPT_IMAGE_1 or FLUX_MAX"`
	FileURL        string   `json:"file_url,omitempty" help:"optional; deprecated single input URL"`
}

type AsyncTaskResponse struct {
	ID       string     `json:"id"`
	Status   TaskStatus `json:"status"`
	Progress string     `json:"progress,omitempty"`
	Error    string     `json:"error,omitempty"`
}

func (r AsyncTaskResponse) GetID() string     { return r.ID }
func (r AsyncTaskResponse) GetStatus() string { return string(r.Status) }
func (r AsyncTaskResponse) GetError() string  { return r.Error }

type Image struct {
	URL string `json:"url"`
}

type GenerationResponse struct {
	AsyncTaskResponse
	Images []Image `json:"images,omitempty"`
}
