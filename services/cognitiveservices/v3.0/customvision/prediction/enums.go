package prediction

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// CustomVisionErrorCodes enumerates the values for custom vision error codes.
type CustomVisionErrorCodes string

const (
	// BadRequest ...
	BadRequest CustomVisionErrorCodes = "BadRequest"
	// BadRequestCannotMigrateProjectWithName ...
	BadRequestCannotMigrateProjectWithName CustomVisionErrorCodes = "BadRequestCannotMigrateProjectWithName"
	// BadRequestClassificationTrainingValidationFailed ...
	BadRequestClassificationTrainingValidationFailed CustomVisionErrorCodes = "BadRequestClassificationTrainingValidationFailed"
	// BadRequestDetectionTrainingNotAllowNegativeTag ...
	BadRequestDetectionTrainingNotAllowNegativeTag CustomVisionErrorCodes = "BadRequestDetectionTrainingNotAllowNegativeTag"
	// BadRequestDetectionTrainingValidationFailed ...
	BadRequestDetectionTrainingValidationFailed CustomVisionErrorCodes = "BadRequestDetectionTrainingValidationFailed"
	// BadRequestDomainNotSupportedForAdvancedTraining ...
	BadRequestDomainNotSupportedForAdvancedTraining CustomVisionErrorCodes = "BadRequestDomainNotSupportedForAdvancedTraining"
	// BadRequestExceededBatchSize ...
	BadRequestExceededBatchSize CustomVisionErrorCodes = "BadRequestExceededBatchSize"
	// BadRequestExceededQuota ...
	BadRequestExceededQuota CustomVisionErrorCodes = "BadRequestExceededQuota"
	// BadRequestExceedIterationPerProjectLimit ...
	BadRequestExceedIterationPerProjectLimit CustomVisionErrorCodes = "BadRequestExceedIterationPerProjectLimit"
	// BadRequestExceedProjectLimit ...
	BadRequestExceedProjectLimit CustomVisionErrorCodes = "BadRequestExceedProjectLimit"
	// BadRequestExceedTagPerImageLimit ...
	BadRequestExceedTagPerImageLimit CustomVisionErrorCodes = "BadRequestExceedTagPerImageLimit"
	// BadRequestExceedTagPerProjectLimit ...
	BadRequestExceedTagPerProjectLimit CustomVisionErrorCodes = "BadRequestExceedTagPerProjectLimit"
	// BadRequestExportAlreadyInProgress ...
	BadRequestExportAlreadyInProgress CustomVisionErrorCodes = "BadRequestExportAlreadyInProgress"
	// BadRequestExportPlatformNotSupportedForAdvancedTraining ...
	BadRequestExportPlatformNotSupportedForAdvancedTraining CustomVisionErrorCodes = "BadRequestExportPlatformNotSupportedForAdvancedTraining"
	// BadRequestExportValidationFailed ...
	BadRequestExportValidationFailed CustomVisionErrorCodes = "BadRequestExportValidationFailed"
	// BadRequestImageBatch ...
	BadRequestImageBatch CustomVisionErrorCodes = "BadRequestImageBatch"
	// BadRequestImageExceededCount ...
	BadRequestImageExceededCount CustomVisionErrorCodes = "BadRequestImageExceededCount"
	// BadRequestImageFormat ...
	BadRequestImageFormat CustomVisionErrorCodes = "BadRequestImageFormat"
	// BadRequestImageRegions ...
	BadRequestImageRegions CustomVisionErrorCodes = "BadRequestImageRegions"
	// BadRequestImageSizeBytes ...
	BadRequestImageSizeBytes CustomVisionErrorCodes = "BadRequestImageSizeBytes"
	// BadRequestImageStream ...
	BadRequestImageStream CustomVisionErrorCodes = "BadRequestImageStream"
	// BadRequestImageTags ...
	BadRequestImageTags CustomVisionErrorCodes = "BadRequestImageTags"
	// BadRequestImageURL ...
	BadRequestImageURL CustomVisionErrorCodes = "BadRequestImageUrl"
	// BadRequestInvalid ...
	BadRequestInvalid CustomVisionErrorCodes = "BadRequestInvalid"
	// BadRequestInvalidEmailAddress ...
	BadRequestInvalidEmailAddress CustomVisionErrorCodes = "BadRequestInvalidEmailAddress"
	// BadRequestInvalidIds ...
	BadRequestInvalidIds CustomVisionErrorCodes = "BadRequestInvalidIds"
	// BadRequestInvalidPublishName ...
	BadRequestInvalidPublishName CustomVisionErrorCodes = "BadRequestInvalidPublishName"
	// BadRequestInvalidPublishTarget ...
	BadRequestInvalidPublishTarget CustomVisionErrorCodes = "BadRequestInvalidPublishTarget"
	// BadRequestIterationDescription ...
	BadRequestIterationDescription CustomVisionErrorCodes = "BadRequestIterationDescription"
	// BadRequestIterationIsNotTrained ...
	BadRequestIterationIsNotTrained CustomVisionErrorCodes = "BadRequestIterationIsNotTrained"
	// BadRequestIterationIsPublished ...
	BadRequestIterationIsPublished CustomVisionErrorCodes = "BadRequestIterationIsPublished"
	// BadRequestIterationName ...
	BadRequestIterationName CustomVisionErrorCodes = "BadRequestIterationName"
	// BadRequestIterationNameNotUnique ...
	BadRequestIterationNameNotUnique CustomVisionErrorCodes = "BadRequestIterationNameNotUnique"
	// BadRequestMultiClassClassificationTrainingValidationFailed ...
	BadRequestMultiClassClassificationTrainingValidationFailed CustomVisionErrorCodes = "BadRequestMultiClassClassificationTrainingValidationFailed"
	// BadRequestMultiLabelClassificationTrainingValidationFailed ...
	BadRequestMultiLabelClassificationTrainingValidationFailed CustomVisionErrorCodes = "BadRequestMultiLabelClassificationTrainingValidationFailed"
	// BadRequestMultipleNegativeTag ...
	BadRequestMultipleNegativeTag CustomVisionErrorCodes = "BadRequestMultipleNegativeTag"
	// BadRequestNegativeAndRegularTagOnSameImage ...
	BadRequestNegativeAndRegularTagOnSameImage CustomVisionErrorCodes = "BadRequestNegativeAndRegularTagOnSameImage"
	// BadRequestNotLimitedTrial ...
	BadRequestNotLimitedTrial CustomVisionErrorCodes = "BadRequestNotLimitedTrial"
	// BadRequestNotSupported ...
	BadRequestNotSupported CustomVisionErrorCodes = "BadRequestNotSupported"
	// BadRequestPredictionIdsExceededCount ...
	BadRequestPredictionIdsExceededCount CustomVisionErrorCodes = "BadRequestPredictionIdsExceededCount"
	// BadRequestPredictionIdsMissing ...
	BadRequestPredictionIdsMissing CustomVisionErrorCodes = "BadRequestPredictionIdsMissing"
	// BadRequestPredictionInvalidApplicationName ...
	BadRequestPredictionInvalidApplicationName CustomVisionErrorCodes = "BadRequestPredictionInvalidApplicationName"
	// BadRequestPredictionInvalidQueryParameters ...
	BadRequestPredictionInvalidQueryParameters CustomVisionErrorCodes = "BadRequestPredictionInvalidQueryParameters"
	// BadRequestPredictionResultsExceededCount ...
	BadRequestPredictionResultsExceededCount CustomVisionErrorCodes = "BadRequestPredictionResultsExceededCount"
	// BadRequestPredictionTagsExceededCount ...
	BadRequestPredictionTagsExceededCount CustomVisionErrorCodes = "BadRequestPredictionTagsExceededCount"
	// BadRequestProjectDescription ...
	BadRequestProjectDescription CustomVisionErrorCodes = "BadRequestProjectDescription"
	// BadRequestProjectName ...
	BadRequestProjectName CustomVisionErrorCodes = "BadRequestProjectName"
	// BadRequestProjectNameNotUnique ...
	BadRequestProjectNameNotUnique CustomVisionErrorCodes = "BadRequestProjectNameNotUnique"
	// BadRequestProjectUnknownClassification ...
	BadRequestProjectUnknownClassification CustomVisionErrorCodes = "BadRequestProjectUnknownClassification"
	// BadRequestProjectUnknownDomain ...
	BadRequestProjectUnknownDomain CustomVisionErrorCodes = "BadRequestProjectUnknownDomain"
	// BadRequestProjectUnsupportedDomainTypeChange ...
	BadRequestProjectUnsupportedDomainTypeChange CustomVisionErrorCodes = "BadRequestProjectUnsupportedDomainTypeChange"
	// BadRequestProjectUnsupportedExportPlatform ...
	BadRequestProjectUnsupportedExportPlatform CustomVisionErrorCodes = "BadRequestProjectUnsupportedExportPlatform"
	// BadRequestRequiredParamIsNull ...
	BadRequestRequiredParamIsNull CustomVisionErrorCodes = "BadRequestRequiredParamIsNull"
	// BadRequestReservedBudgetInHoursNotEnoughForAdvancedTraining ...
	BadRequestReservedBudgetInHoursNotEnoughForAdvancedTraining CustomVisionErrorCodes = "BadRequestReservedBudgetInHoursNotEnoughForAdvancedTraining"
	// BadRequestSubscriptionAPI ...
	BadRequestSubscriptionAPI CustomVisionErrorCodes = "BadRequestSubscriptionApi"
	// BadRequestTagDescription ...
	BadRequestTagDescription CustomVisionErrorCodes = "BadRequestTagDescription"
	// BadRequestTagName ...
	BadRequestTagName CustomVisionErrorCodes = "BadRequestTagName"
	// BadRequestTagNameNotUnique ...
	BadRequestTagNameNotUnique CustomVisionErrorCodes = "BadRequestTagNameNotUnique"
	// BadRequestTagType ...
	BadRequestTagType CustomVisionErrorCodes = "BadRequestTagType"
	// BadRequestTrainingAlreadyInProgress ...
	BadRequestTrainingAlreadyInProgress CustomVisionErrorCodes = "BadRequestTrainingAlreadyInProgress"
	// BadRequestTrainingNotNeeded ...
	BadRequestTrainingNotNeeded CustomVisionErrorCodes = "BadRequestTrainingNotNeeded"
	// BadRequestTrainingNotNeededButTrainingPipelineUpdated ...
	BadRequestTrainingNotNeededButTrainingPipelineUpdated CustomVisionErrorCodes = "BadRequestTrainingNotNeededButTrainingPipelineUpdated"
	// BadRequestTrainingValidationFailed ...
	BadRequestTrainingValidationFailed CustomVisionErrorCodes = "BadRequestTrainingValidationFailed"
	// BadRequestUnpublishFailed ...
	BadRequestUnpublishFailed CustomVisionErrorCodes = "BadRequestUnpublishFailed"
	// BadRequestWorkspaceCannotBeModified ...
	BadRequestWorkspaceCannotBeModified CustomVisionErrorCodes = "BadRequestWorkspaceCannotBeModified"
	// BadRequestWorkspaceNotDeletable ...
	BadRequestWorkspaceNotDeletable CustomVisionErrorCodes = "BadRequestWorkspaceNotDeletable"
	// Conflict ...
	Conflict CustomVisionErrorCodes = "Conflict"
	// ConflictInvalid ...
	ConflictInvalid CustomVisionErrorCodes = "ConflictInvalid"
	// ErrorExporterInvalidClassifier ...
	ErrorExporterInvalidClassifier CustomVisionErrorCodes = "ErrorExporterInvalidClassifier"
	// ErrorExporterInvalidFeaturizer ...
	ErrorExporterInvalidFeaturizer CustomVisionErrorCodes = "ErrorExporterInvalidFeaturizer"
	// ErrorExporterInvalidPlatform ...
	ErrorExporterInvalidPlatform CustomVisionErrorCodes = "ErrorExporterInvalidPlatform"
	// ErrorFeaturizationAugmentationError ...
	ErrorFeaturizationAugmentationError CustomVisionErrorCodes = "ErrorFeaturizationAugmentationError"
	// ErrorFeaturizationAugmentationUnavailable ...
	ErrorFeaturizationAugmentationUnavailable CustomVisionErrorCodes = "ErrorFeaturizationAugmentationUnavailable"
	// ErrorFeaturizationInvalidFeaturizer ...
	ErrorFeaturizationInvalidFeaturizer CustomVisionErrorCodes = "ErrorFeaturizationInvalidFeaturizer"
	// ErrorFeaturizationQueueTimeout ...
	ErrorFeaturizationQueueTimeout CustomVisionErrorCodes = "ErrorFeaturizationQueueTimeout"
	// ErrorFeaturizationServiceUnavailable ...
	ErrorFeaturizationServiceUnavailable CustomVisionErrorCodes = "ErrorFeaturizationServiceUnavailable"
	// ErrorFeaturizationUnrecognizedJob ...
	ErrorFeaturizationUnrecognizedJob CustomVisionErrorCodes = "ErrorFeaturizationUnrecognizedJob"
	// ErrorInvalid ...
	ErrorInvalid CustomVisionErrorCodes = "ErrorInvalid"
	// ErrorPrediction ...
	ErrorPrediction CustomVisionErrorCodes = "ErrorPrediction"
	// ErrorPredictionModelNotCached ...
	ErrorPredictionModelNotCached CustomVisionErrorCodes = "ErrorPredictionModelNotCached"
	// ErrorPredictionModelNotFound ...
	ErrorPredictionModelNotFound CustomVisionErrorCodes = "ErrorPredictionModelNotFound"
	// ErrorPredictionServiceUnavailable ...
	ErrorPredictionServiceUnavailable CustomVisionErrorCodes = "ErrorPredictionServiceUnavailable"
	// ErrorPredictionStorage ...
	ErrorPredictionStorage CustomVisionErrorCodes = "ErrorPredictionStorage"
	// ErrorProjectExportRequestFailed ...
	ErrorProjectExportRequestFailed CustomVisionErrorCodes = "ErrorProjectExportRequestFailed"
	// ErrorProjectInvalidDomain ...
	ErrorProjectInvalidDomain CustomVisionErrorCodes = "ErrorProjectInvalidDomain"
	// ErrorProjectInvalidPipelineConfiguration ...
	ErrorProjectInvalidPipelineConfiguration CustomVisionErrorCodes = "ErrorProjectInvalidPipelineConfiguration"
	// ErrorProjectInvalidWorkspace ...
	ErrorProjectInvalidWorkspace CustomVisionErrorCodes = "ErrorProjectInvalidWorkspace"
	// ErrorProjectTrainingRequestFailed ...
	ErrorProjectTrainingRequestFailed CustomVisionErrorCodes = "ErrorProjectTrainingRequestFailed"
	// ErrorRegionProposal ...
	ErrorRegionProposal CustomVisionErrorCodes = "ErrorRegionProposal"
	// ErrorUnknown ...
	ErrorUnknown CustomVisionErrorCodes = "ErrorUnknown"
	// Forbidden ...
	Forbidden CustomVisionErrorCodes = "Forbidden"
	// ForbiddenDRModeEnabled ...
	ForbiddenDRModeEnabled CustomVisionErrorCodes = "ForbiddenDRModeEnabled"
	// ForbiddenInvalid ...
	ForbiddenInvalid CustomVisionErrorCodes = "ForbiddenInvalid"
	// ForbiddenUser ...
	ForbiddenUser CustomVisionErrorCodes = "ForbiddenUser"
	// ForbiddenUserDisabled ...
	ForbiddenUserDisabled CustomVisionErrorCodes = "ForbiddenUserDisabled"
	// ForbiddenUserDoesNotExist ...
	ForbiddenUserDoesNotExist CustomVisionErrorCodes = "ForbiddenUserDoesNotExist"
	// ForbiddenUserInsufficientCapability ...
	ForbiddenUserInsufficientCapability CustomVisionErrorCodes = "ForbiddenUserInsufficientCapability"
	// ForbiddenUserResource ...
	ForbiddenUserResource CustomVisionErrorCodes = "ForbiddenUserResource"
	// ForbiddenUserSignupAllowanceExceeded ...
	ForbiddenUserSignupAllowanceExceeded CustomVisionErrorCodes = "ForbiddenUserSignupAllowanceExceeded"
	// ForbiddenUserSignupDisabled ...
	ForbiddenUserSignupDisabled CustomVisionErrorCodes = "ForbiddenUserSignupDisabled"
	// NoError ...
	NoError CustomVisionErrorCodes = "NoError"
	// NotFound ...
	NotFound CustomVisionErrorCodes = "NotFound"
	// NotFoundApimSubscription ...
	NotFoundApimSubscription CustomVisionErrorCodes = "NotFoundApimSubscription"
	// NotFoundDomain ...
	NotFoundDomain CustomVisionErrorCodes = "NotFoundDomain"
	// NotFoundImage ...
	NotFoundImage CustomVisionErrorCodes = "NotFoundImage"
	// NotFoundInvalid ...
	NotFoundInvalid CustomVisionErrorCodes = "NotFoundInvalid"
	// NotFoundIteration ...
	NotFoundIteration CustomVisionErrorCodes = "NotFoundIteration"
	// NotFoundIterationPerformance ...
	NotFoundIterationPerformance CustomVisionErrorCodes = "NotFoundIterationPerformance"
	// NotFoundProject ...
	NotFoundProject CustomVisionErrorCodes = "NotFoundProject"
	// NotFoundProjectDefaultIteration ...
	NotFoundProjectDefaultIteration CustomVisionErrorCodes = "NotFoundProjectDefaultIteration"
	// NotFoundTag ...
	NotFoundTag CustomVisionErrorCodes = "NotFoundTag"
	// UnsupportedMediaType ...
	UnsupportedMediaType CustomVisionErrorCodes = "UnsupportedMediaType"
)

// PossibleCustomVisionErrorCodesValues returns an array of possible values for the CustomVisionErrorCodes const type.
func PossibleCustomVisionErrorCodesValues() []CustomVisionErrorCodes {
	return []CustomVisionErrorCodes{BadRequest, BadRequestCannotMigrateProjectWithName, BadRequestClassificationTrainingValidationFailed, BadRequestDetectionTrainingNotAllowNegativeTag, BadRequestDetectionTrainingValidationFailed, BadRequestDomainNotSupportedForAdvancedTraining, BadRequestExceededBatchSize, BadRequestExceededQuota, BadRequestExceedIterationPerProjectLimit, BadRequestExceedProjectLimit, BadRequestExceedTagPerImageLimit, BadRequestExceedTagPerProjectLimit, BadRequestExportAlreadyInProgress, BadRequestExportPlatformNotSupportedForAdvancedTraining, BadRequestExportValidationFailed, BadRequestImageBatch, BadRequestImageExceededCount, BadRequestImageFormat, BadRequestImageRegions, BadRequestImageSizeBytes, BadRequestImageStream, BadRequestImageTags, BadRequestImageURL, BadRequestInvalid, BadRequestInvalidEmailAddress, BadRequestInvalidIds, BadRequestInvalidPublishName, BadRequestInvalidPublishTarget, BadRequestIterationDescription, BadRequestIterationIsNotTrained, BadRequestIterationIsPublished, BadRequestIterationName, BadRequestIterationNameNotUnique, BadRequestMultiClassClassificationTrainingValidationFailed, BadRequestMultiLabelClassificationTrainingValidationFailed, BadRequestMultipleNegativeTag, BadRequestNegativeAndRegularTagOnSameImage, BadRequestNotLimitedTrial, BadRequestNotSupported, BadRequestPredictionIdsExceededCount, BadRequestPredictionIdsMissing, BadRequestPredictionInvalidApplicationName, BadRequestPredictionInvalidQueryParameters, BadRequestPredictionResultsExceededCount, BadRequestPredictionTagsExceededCount, BadRequestProjectDescription, BadRequestProjectName, BadRequestProjectNameNotUnique, BadRequestProjectUnknownClassification, BadRequestProjectUnknownDomain, BadRequestProjectUnsupportedDomainTypeChange, BadRequestProjectUnsupportedExportPlatform, BadRequestRequiredParamIsNull, BadRequestReservedBudgetInHoursNotEnoughForAdvancedTraining, BadRequestSubscriptionAPI, BadRequestTagDescription, BadRequestTagName, BadRequestTagNameNotUnique, BadRequestTagType, BadRequestTrainingAlreadyInProgress, BadRequestTrainingNotNeeded, BadRequestTrainingNotNeededButTrainingPipelineUpdated, BadRequestTrainingValidationFailed, BadRequestUnpublishFailed, BadRequestWorkspaceCannotBeModified, BadRequestWorkspaceNotDeletable, Conflict, ConflictInvalid, ErrorExporterInvalidClassifier, ErrorExporterInvalidFeaturizer, ErrorExporterInvalidPlatform, ErrorFeaturizationAugmentationError, ErrorFeaturizationAugmentationUnavailable, ErrorFeaturizationInvalidFeaturizer, ErrorFeaturizationQueueTimeout, ErrorFeaturizationServiceUnavailable, ErrorFeaturizationUnrecognizedJob, ErrorInvalid, ErrorPrediction, ErrorPredictionModelNotCached, ErrorPredictionModelNotFound, ErrorPredictionServiceUnavailable, ErrorPredictionStorage, ErrorProjectExportRequestFailed, ErrorProjectInvalidDomain, ErrorProjectInvalidPipelineConfiguration, ErrorProjectInvalidWorkspace, ErrorProjectTrainingRequestFailed, ErrorRegionProposal, ErrorUnknown, Forbidden, ForbiddenDRModeEnabled, ForbiddenInvalid, ForbiddenUser, ForbiddenUserDisabled, ForbiddenUserDoesNotExist, ForbiddenUserInsufficientCapability, ForbiddenUserResource, ForbiddenUserSignupAllowanceExceeded, ForbiddenUserSignupDisabled, NoError, NotFound, NotFoundApimSubscription, NotFoundDomain, NotFoundImage, NotFoundInvalid, NotFoundIteration, NotFoundIterationPerformance, NotFoundProject, NotFoundProjectDefaultIteration, NotFoundTag, UnsupportedMediaType}
}
