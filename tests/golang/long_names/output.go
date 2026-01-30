package long_names

type ContainerItemsItem struct {
	AVeryLongInlinePropertyNameThatWillCreateAlongGeneratedTypeName   *string  `json:"aVeryLongInlinePropertyNameThatWillCreateALongGeneratedTypeName,omitempty"`
	AnotherLongPropertyNameForTestingPurposesOnlyDoNotUseInProduction *float64 `json:"anotherLongPropertyNameForTestingPurposesOnlyDoNotUseInProduction,omitempty"`
}

type ThisIsAnExtremelyLongTypeNameThatShouldStillWorkCorrectlyInTheCodeGenerator struct {
	AnotherVeryLongPropertyNameWithLotsOfWordsInItThatKeepsGoingAndGoing *int    `json:"anotherVeryLongPropertyNameWithLotsOfWordsInItThatKeepsGoingAndGoing,omitempty"`
	ThisIsAlsoAnExtremelyLongPropertyNameThatShouldBeHandledProperly     *string `json:"thisIsAlsoAnExtremelyLongPropertyNameThatShouldBeHandledProperly,omitempty"`
	YetAnotherIncrediblyLongPropertyNameThatTestsTheLimitsOfTheGenerator *bool   `json:"yetAnotherIncrediblyLongPropertyNameThatTestsTheLimitsOfTheGenerator,omitempty"`
}

type Container struct {
	Items                                                          []ContainerItemsItem                                                         `json:"items,omitempty"`
	ShortName                                                      *string                                                                      `json:"shortName,omitempty"`
	TheQuickBrownFoxJumpsOverTheLazyDogThisIsAveryLongPropertyName *ThisIsAnExtremelyLongTypeNameThatShouldStillWorkCorrectlyInTheCodeGenerator `json:"theQuickBrownFoxJumpsOverTheLazyDogThisIsAVeryLongPropertyName,omitempty"`
}
