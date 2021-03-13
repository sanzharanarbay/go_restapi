package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/sanzharanarbay/go_restapi/api/models"
	"github.com/sanzharanarbay/go_restapi/api/responses"
)

// CreateVenue parses request, validates data and saves the new venue
func (a *App) CreateVenue(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": "success", "message": "Venue successfully created"}

	venue := &models.Venue{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = json.Unmarshal(body, &venue)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	venue.Prepare()

	if err = venue.Validate(); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if vne, _ := venue.GetVenue(a.DB); vne != nil {
		resp["status"] = "failed"
		resp["message"] = "Venue already registered, please choose another name"
		responses.JSON(w, http.StatusBadRequest, resp)
		return
	}

	venueCreated, err := venue.Save(a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	resp["venue"] = venueCreated
	responses.JSON(w, http.StatusCreated, resp)
	return
}

func (a *App) GetVenues(w http.ResponseWriter, r *http.Request) {
	venues, err := models.GetVenues(a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, venues)
	return
}

func (a *App) GetVenue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	venue, err := models.GetVenueById(id, a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, venue)
	return
}



func (a *App) UpdateVenue(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": "success", "message": "Venue updated successfully"}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	venueUpdate := models.Venue{}
	if err = json.Unmarshal(body, &venueUpdate); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	venueUpdate.Prepare()

	_, err = venueUpdate.UpdateVenue(id, a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, resp)
	return
}

func (a *App) DeleteVenue(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": "success", "message": "Venue deleted successfully"}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	venue, err := models.GetVenueById(id, a.DB)

	if venue ==  nil{
		resp["status"] = "error"
		resp["message"] = "Not Found"
		responses.JSON(w, http.StatusNotFound, resp)
		return
	}

	err = models.DeleteVenue(id, a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, resp)
	return
}