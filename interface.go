package goopenai

import "context"

type ClientInterface interface {
	CreateChats(ctx context.Context, r CreateChatsRequest) (response CreateChatsResponse, err error)
	CreateCompletions(ctx context.Context, r CreateCompletionsRequest) (response CreateCompletionsResponse, err error)
	CreateEdits(ctx context.Context, r CreateEditsRequest) (response CreateEditsResponse, err error)
	CreateEmbeddings(ctx context.Context, r CreateEmbeddingsRequest) (response CreateEmbeddingsResponse, err error)
	CreateImages(ctx context.Context, r CreateImagesRequest) (response CreateImagesResponse, err error)
	CreateImagesEdits(ctx context.Context, r CreateImagesEditsRequest) (response CreateImagesEditsResponse, err error)
	CreateImagesVariations(ctx context.Context, r CreateImagesVariationsRequest) (response CreateImagesVariationsResponse, err error)
	GetModel(ctx context.Context, id string) (response GetModelResponse, err error)
	GetModels(ctx context.Context) (response GetModelsResponse, err error)
	CreateTranscriptions(ctx context.Context, r CreateTranscriptionsRequest) (response CreateTranscriptionsResponse, err error)
	CreateTranslations(ctx context.Context, r CreateTranslationsRequest) (response CreateTranslationsResponse, err error)
	CreateModerations(ctx context.Context, r CreateModerationsRequest) (response CreateModerationsResponse, err error)
}
