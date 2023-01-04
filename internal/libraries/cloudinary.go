package libraries

import (
	"context"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
)

func Setupcloudinary() (*cloudinary.Cloudinary, context.Context) {
	cloudinary_name := os.Getenv("CLODINARY_NAME")
	cloudinary_key := os.Getenv("CLODINARY_KEY")
	cloudinary_secret := os.Getenv("CLODINARY_SECRET")

	cld, _ := cloudinary.NewFromParams(cloudinary_name, cloudinary_key, cloudinary_secret)
	ctx := context.Background()

	return cld, ctx
}
