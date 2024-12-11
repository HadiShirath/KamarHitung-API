package utility

import (
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		public_id := uuid.NewString()
		tokenString, err := GenerateToken(public_id, "fullname", "user", "IniSecret")

		require.Nil(t, err)
		require.NotEmpty(t, tokenString)
		log.Println(tokenString)
	})
}

func TestVerifyToken(t *testing.T) {
	t.Run("verify token success", func(t *testing.T) {
		publicId := uuid.NewString()
		role := "user"
		fullname := "user"
		tokenString, err := GenerateToken(publicId, fullname, role, "IniSecret")
		require.Nil(t, err)
		require.NotEmpty(t, tokenString)

		jwtId, jwtRole, err := ValidateToken(tokenString, "IniSecret")
		require.Nil(t, err)
		require.NotEmpty(t, jwtId)
		require.NotEmpty(t, jwtRole)

		require.Equal(t, publicId, jwtId)
		require.Equal(t, role, jwtRole)
	})
}
