package persona

import "path"

func New(name, location string) (Persona, error) {
	return Persona{
		APIVersion: apiVersion,
		location: path.Join(location, name),
		name: name,
	}, nil
}

