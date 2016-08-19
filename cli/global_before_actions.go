package main

var globalBeforeActions = []action{
	ensureMetadataAction,
	injectContainerAction,
	injectProjectConfigAction,
	validateProjectConfigAction,
	injectAuthTokenAction,
}
