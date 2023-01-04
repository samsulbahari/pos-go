package main

import (
	_categoryHadler "clean-arsitecture/internal/category/handler"
	_categoryRepo "clean-arsitecture/internal/category/repository"
	_categoryService "clean-arsitecture/internal/category/service"
	"clean-arsitecture/internal/middleware"

	_authHadler "clean-arsitecture/internal/auth/handler"
	_authRepo "clean-arsitecture/internal/auth/repository"
	_authService "clean-arsitecture/internal/auth/service"
	"clean-arsitecture/internal/database"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	r := gin.Default()

	r.Use()

	db := database.ConnectDatabase()

	middleware := middleware.WithAuth()

	authRepo := _authRepo.NewAuthRepository(db)
	authService := _authService.NewAuthService(authRepo)
	_authHadler.NewAuthHandler(r, authService)

	categoryRepo := _categoryRepo.NewCategoryRepository(db)
	categoryService := _categoryService.NewCategoryService(categoryRepo)

	authorized := r.Group("")
	authorized.Use(middleware)
	{
		_categoryHadler.NewCategoryHandler(authorized, categoryService)
	}

	r.Run()
}
