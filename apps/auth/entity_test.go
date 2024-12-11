package auth

import (
	"kamar-hitung/infra/response"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateAuthEntity(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		authEntity := AuthEntity{
			Username: "hahah@gmail.com",
			Password: "12345678",
		}

		err := authEntity.Validate()
		require.Nil(t, err)
	})

	t.Run("email is require", func(t *testing.T) {
		authEntity := AuthEntity{
			Username: "",
			Password: "12345678",
		}

		err := authEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrEmailRequired, err)
	})

	t.Run("email is invalid", func(t *testing.T) {
		authEntity := AuthEntity{
			Username: "haha.gmail.com",
			Password: "12345678",
		}

		err := authEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrEmailInvalid, err)
	})

	t.Run("password is required", func(t *testing.T) {
		authEntity := AuthEntity{
			Username: "haha@gmail.com",
			Password: "",
		}

		err := authEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrPasswordRequired, err)
	})

	t.Run("password must have minimun 8 character", func(t *testing.T) {
		authEntity := AuthEntity{
			Username: "haha@gmail.com",
			Password: "123",
		}

		err := authEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrPasswordInvalidLength, err)
	})

}
