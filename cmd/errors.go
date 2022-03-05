package cmd

import "errors"

var (
	errIdentityPluginNotFound = errors.New("cannot find identity plugin")
	errLoadingCorePlugins     = errors.New("cannot load core plugins")
	errPersonaDontExists      = errors.New("persona does not exists")
	errPluginNotInstanceOf    = errors.New("retrieved plugin is not an instance of identity.Plugin")
	//nolint:lll // line too long but too lazy for making it multiline
	errRehashWithActiveProfile = errors.New("trying to rehash with an active profile, this may go very wrong so it's forbidden")
)
