package auth

import (
    "log"
    "time"
    "errors"
    "os"
    "strings"

    "github.com/google/uuid"
    "github.com/dgrijalva/jwt-go"
    "github.com/joho/godotenv"
)

// JWTClaim is the struct of information in the JWT payload.
type JWTClaim struct {
    UUID uuid.UUID `json:"uuid"`
    jwt.StandardClaims
}

// getJWTKey returns the secret key defined in .env which signs the token.
func getJWTKey() []byte {
    err := godotenv.Load(".env")
    if (err != nil) {
        log.Fatalf("Error loading .env file")
    }

    return []byte(os.Getenv("JWT_KEY"))
}

// GenerateJWT generates a new token with the provided UUID embedded in the payload.
func GenerateJWT(uuid uuid.UUID) (tokenString string, err error) {
    jwtKey := getJWTKey()

    expirationTime := time.Now().Add(2 * time.Hour)
    claims := &JWTClaim {
        UUID: uuid,
        StandardClaims: jwt.StandardClaims {
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    return token.SignedString(jwtKey)
}

// ValidateToken validates the given token. Returns an error if there's any problem
// with the validation
func ValidateToken(signedToken string) error {
    jwtKey := getJWTKey()

    _, err := jwt.ParseWithClaims(
            signedToken,
            &JWTClaim{},
            func (token *jwt.Token) (interface{}, error) {
                _, ok := token.Method.(*jwt.SigningMethodHMAC)
                if (!ok) {
                    err := errors.New("wrongAlg")
                    return nil, err
                }

                return []byte(jwtKey), nil
            },
    )

    if (err != nil) {
        if (strings.Contains(err.Error(), "token is expired")) {
            err = errors.New("tokenExpired")
        } else if (strings.Contains(err.Error(), "token contains an invalid number of segments")) {
            err = errors.New("invalidTokenLength")
        } else if (strings.Contains(err.Error(), "invalid character")) {
            err = errors.New("invalidTokenChar")
        } 

        return err
    }

    return nil
}

// GetUUIDFromToken returns the UUID embedded in the provided token payload.
func GetUUIDFromToken(signedToken string) (uuid.UUID, error) {
    jwtKey := getJWTKey()

    token, err := jwt.ParseWithClaims(
            signedToken,
            &JWTClaim{},
            func (token *jwt.Token) (interface{}, error) {
                return []byte(jwtKey), nil
            },
    )

    var uuid uuid.UUID

    if (err != nil) {
        return uuid, err
    }

    claims, ok := token.Claims.(*JWTClaim)
    if (!ok) {
        err = errors.New("Couldn't parse claims")
        return uuid, err
    }

    return claims.UUID, nil
}
