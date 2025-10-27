package student

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/sachinpoudel/learning_go/RestApi/internal/types"
	"github.com/sachinpoudel/learning_go/RestApi/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		if err != nil {
			response.WriteJson(w,http.StatusBadRequest, response.GeneralError(err))
		}

//request validation

if err:=validator.New().Struct(student); err != nil {
	validateErrs := err.(validator.ValidationErrors)
	response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
	return 
}


		w.Write([]byte("welcome to students api"))
		response.WriteJson(w, http.StatusCreated, map[string]string {"success":"ok"})
	}
}
