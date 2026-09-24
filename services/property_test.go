package services

import (
	"errors"
	"os"
	"testing"
	"w3-a4/models"

	"github.com/beego/beego/v2/server/web"
)

func TestMain(m *testing.M) {

	err := os.Chdir("..")

	if err != nil {
		panic(err)
	}

	err = web.LoadAppConfig(
		"ini",
		"conf/app.conf",
	)

	if err != nil {
		panic(err)
	}

	os.Exit(
		m.Run(),
	)
}

func TestGetPropertyByID(t *testing.T) {

	properties := []models.SourceProperty{

		{
			ID:           "TEST-1",
			PropertyName: "Hotel One",
		},

		{
			ID:           "TEST-2",
			PropertyName: "Hotel Two",
		},
	}

	tests := []struct {
		name      string
		id        string
		wantFound bool
	}{

		{
			name:      "existing property",
			id:        "TEST-1",
			wantFound: true,
		},

		{
			name:      "another existing property",
			id:        "TEST-2",
			wantFound: true,
		},

		{
			name:      "unknown property",
			id:        "UNKNOWN",
			wantFound: false,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			result, err := FindPropertyByID(
				properties,
				tt.id,
			)

			if tt.wantFound {

				if err != nil {
					t.Fatal(err)
				}

				if result == nil {

					t.Fatal(
						"expected property",
					)

				}

			} else {

				if !errors.Is(
					err,
					ErrPropertyNotFound,
				) {

					t.Errorf(
						"expected not found error",
					)

				}

				if result != nil {

					t.Error(
						"expected nil result",
					)

				}

			}

		})

	}

}

func TestListProperties_WithLimit(t *testing.T) {

	filter := models.PropertyFilter{
		Limit: 5,
	}

	response, err := ListProperties(filter)

	if err != nil {
		t.Fatal(err)
	}

	if response.Result.Count > 5 {

		t.Errorf(
			"limit not applied",
		)

	}

}
