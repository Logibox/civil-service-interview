package v1

import (
	bpdts_client "github.com/Logibox/civil-service-interview/v1/bpdts/client"
	bpdts_params "github.com/Logibox/civil-service-interview/v1/bpdts/client/default_operations"
	bpdts_models "github.com/Logibox/civil-service-interview/v1/bpdts/models"
	"github.com/Logibox/civil-service-interview/v1/nominatim"
	"github.com/Logibox/civil-service-interview/v1/models"
)


func FindUsers(city, country string, within float64) (models.UserList, error) {
	ld, err := nominatim.GetCityLocation(city, country)
	if err != nil {
		return nil, err
	}
	lat, err := ld.Latitude.Float64()
	if err != nil {
		return nil, err
	}
	long, err := ld.Longitude.Float64()
	if err != nil {
		return nil, err
	}
	cityLoc := &Location{lat, long}

	resp, err := bpdts_client.Default.DefaultOperations.GetAllUsers(nil)
	if err != nil {
		return nil, err
	}
	userList := make([]*bpdts_models.User, 0, 256)
	for _, user := range resp.Payload {
		userLoc, err := UserLocation(user)
		if err != nil {
			return nil, err
		}
		distance := Distance(*cityLoc, *userLoc)
		if distance < within {
			userList = append(userList, user)
		}
	}

	extraUsers, err := FindRegisteredUsers(city)
	if err != nil {
		return nil, err
	}

	userList = append(userList, extraUsers...)

	return convertList(userList)
}

func FindRegisteredUsers(city string) (bpdts_models.UserList, error) {
	resp, err := bpdts_client.Default.DefaultOperations.GetUsersInCity(bpdts_params.NewGetUsersInCityParams().WithCity(city))
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func convertList(input bpdts_models.UserList) (models.UserList, error) {
	out := make(models.UserList, 0, 64)
	for _, u := range input {
		lat, err := u.Latitude.Float64()
		if err != nil {
			return nil, err
		}
		long, err := u.Longitude.Float64()
		if err != nil {
			return nil, err
		}

		out = append(out, &models.User{
			Email: u.Email,
			FirstName: u.FirstName,
			ID: u.ID,
			IPAddress: u.IPAddress,
			LastName: u.LastName,
			Latitude: lat,
			Longitude: long,
		})
	}

	return out, nil
}
