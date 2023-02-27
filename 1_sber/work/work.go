package work

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/cr00z/parser/utils"
)

type Work struct {
	Idx        int
	ctx        context.Context
	startLink  string
	cookies    []*http.Cookie
	formValues map[string]string
	Step       int
	Finished   bool
}

func NewWork(idx int, ctx context.Context, startLink string) *Work {
	return &Work{
		Idx:       idx,
		ctx:       ctx,
		startLink: startLink,
	}
}

func (w *Work) Do() error {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"

	switch w.Step {
	case 0:
		var err error
		_, _, w.cookies, err = utils.GetPage(w.ctx,
			w.startLink, http.MethodGet, nil, nil, nil)
		if err != nil {
			return fmt.Errorf("auth error: %w", err)
		}
		log.Printf("Work %d auth complete: %v", w.Idx, w.cookies)

	default:
		method := http.MethodPost
		linkId := w.Step - 1
		if w.Step == 1 {
			method = http.MethodGet
			linkId = w.Step
		}
		activeLink := w.startLink + "/question/" + strconv.Itoa(linkId)

		doc, code, _, err := utils.GetPage(w.ctx,
			activeLink, method, w.cookies, headers, w.formValues)
		if err != nil {
			return fmt.Errorf("Work %d get question %d error: %w",
				w.Idx, w.Step, err)
		}
		if code != 200 {
			return fmt.Errorf("Work %d wrong question %d answer code: %d",
				w.Idx, w.Step, code)
		}

		w.formValues = utils.ParseForm(doc)
		if len(w.formValues) == 0 {
			// check complete
			if strings.Contains(doc.Text(), "Test successfully passed") {
				log.Printf("Work %d test successfully passed", w.Idx)
				w.Finished = true
				return nil
			}
			// else report error
			return errors.New("parse form error")
		}

		log.Printf("Work %d question %d answer: %v", w.Idx, w.Step, w.formValues)
	}

	w.Step++
	return nil
}
