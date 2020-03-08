package factory

import (
	"testing"

	"github.com/purwantogz/go-impact/models"
)

func TestBuildToken(t *testing.T) {

	token := New("asasasaaa8283s823823823", "go-impact-auth", 30)

	_, err := token.Build("purwanto.dev@gmail.com", &models.Roles{
		RoleType: "admin",
		Scope: models.Scopes{
			Create: true,
			Delete: false,
			Edit:   true,
			Read:   false,
		},
	})

	if err != nil {
		t.Error(err.Error())
	}

}

func TestRefresh(t *testing.T) {
	token := New("asasasaaa8283s823823823", "go-impact-auth", 30)
	_, err := token.Refresh("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODM2MDc4ODksImlzcyI6ImdvLWltcGFjdC1hdXRoIiwiZW1haWwiOiJwdXJ3YW50by5kZXZAZ21haWwuY29tIiwicm9sZXMiOnsicm9sZVR5cGUiOiJhZG1pbiIsInNjb3BlIjp7ImNyZWF0ZSI6dHJ1ZSwicmVhZCI6ZmFsc2UsImVkaXQiOnRydWUsImRlbGV0ZSI6ZmFsc2V9fX0.lHDNlzMNyZRhY8jcPQxIPuy4jexNnMKm_bOfU72lzFI")

	if err != nil {
		t.Error(err.Error())
	}
}
