package components

import (
	"encoding/json"
	"github.com/lefelys/apibara/examples/generated/codes"
	"github.com/lefelys/apibara/examples/generated/errors"
	"io"
	"net/http"
	"strconv"
	"time"
)

///////////
// Schemas
///////////

type Order struct {
	Id       int64     `json:"id"`
	PetId    int64     `json:"petId"`
	Quantity int32     `json:"quantity"`
	ShipDate time.Time `json:"shipDate"`
	Status   string    `json:"status"`   // Order Status. placed || approved || delivered
	Complete bool      `json:"complete"` // default: false
}

type Category struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type User struct {
	Id         int64  `json:"id"`
	Username   string `json:"username"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	UserStatus int32  `json:"userStatus"`
}

type Tag struct {
	Id   int64 `json:"id"`
	Name int64 `json:"name"`
}

type Pet struct {
	Id        int64    `json:"id"`
	Category  Category `json:"category"`
	Name      string   `json:"name"`      // required, example: doggie
	PhotoUrls []string `json:"photoUrls"` // required
	Tags      []Tag    `json:"tags"`
	Status    bool     `json:"status"` // description: pet status in the store, available || pending || sold
}

func (p *Pet) DecodeJSON(r io.Reader) (*Pet, error) {
	err := json.NewDecoder(r).Decode(p)
	return p, err
}

type Cat struct {
	Id        int64    `json:"id"`
	Category  Category `json:"category"`
	Name      string   `json:"name"`      // required, example: doggie
	PhotoUrls []string `json:"photoUrls"` // required
	Tags      []Tag    `json:"tags"`
	Status    bool     `json:"status"` // description: pet status in the store, available || pending || sold
	Nya       bool     `json:"nya"`
}

type Dog struct {
	Id        int64    `json:"id"`
	Category  Category `json:"category"`
	Name      string   `json:"name"`      // required, example: doggie
	PhotoUrls []string `json:"photoUrls"` // required
	Tags      []Tag    `json:"tags"`
	Status    bool     `json:"status"` // description: pet status in the store, available || pending || sold
	Bark      bool     `json:"bark"`
}

type ApiResponse struct {
	Code    int32  `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

/////////////
// Responses
/////////////

type BadRequestObject struct {
	Message *string `json:"message"`
}

// Invalid request
type BadRequest struct {
	Body BadRequestObject
	codes.StatusBadRequest
}

func (resp *BadRequest) Respond(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.Body)
}

/////////////
// Parameter
/////////////

// Limits the number of items on a page
func PageLimit(r *http.Request) (pageLimit int32, err error) {
	param := "limit"
	limitStr := r.URL.Query().Get(param)
	if len(limitStr) == 0 {
		pageLimit = 20 // default
		return
	}

	limitInt64, err := strconv.ParseInt(limitStr, 10, 32)
	if err != nil {
		err = errors.NewBadRequestError(param, err)
		return
	}
	pageLimit = int32(limitInt64)

	if pageLimit < 1 {
		err = errors.NewMinValidationError(param, 1)
		return
	}

	if pageLimit > 100 {
		err = errors.NewMaxValidationError(param, 1)
		return
	}

	return
}

// Specifies the page number of the artists to be displayed
func PageOffset(r *http.Request) (pageOffset int64, err error) {
	param := "offset"
	offsetStr := r.URL.Query().Get(param)
	if len(offsetStr) == 0 {
		err = errors.NewRequiredValidationError(param)
		return
	}
	pageOffset, err = strconv.ParseInt(offsetStr, 10, 32)
	if err != nil {
		err = errors.NewBadRequestError(param, err)
		return
	}
	return
}
