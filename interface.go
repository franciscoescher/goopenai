package goopenai

type ClientInterface interface {
	CreateChats(r CreateChatsRequest) (response CreateChatsResponse, err error)
	CreateCompletions(r CreateCompletionsRequest) (response CreateCompletionsResponse, err error)
	CreateEdits(r CreateEditsRequest) (response CreateEditsResponse, err error)
	CreateEmbeddings(r CreateEmbeddingsRequest) (response CreateEmbeddingsResponse, err error)
	CreateImages(r CreateImagesRequest) (response CreateImagesResponse, err error)
	CreateImagesEdits(r CreateImagesEditsRequest) (response CreateImagesEditsResponse, err error)
	CreateImagesVariations(r CreateImagesVariationsRequest) (response CreateImagesVariationsResponse, err error)
	GetModel(id string) (response GetModelResponse, err error)
	GetModels() (response GetModelsResponse, err error)
	CreateTranscriptions(r CreateTranscriptionsRequest) (response CreateTranscriptionsResponse, err error)
	CreateTranslations(r CreateTranslationsRequest) (response CreateTranslationsResponse, err error)
	CreateModerations(r CreateModerationsRequest) (response CreateModerationsResponse, err error)
}
