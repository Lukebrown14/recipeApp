package handlers

import (
	"log"
	"net/http"
	"recipeapp/utils"
)

func SearchRecipeHandler(w http.ResponseWriter, r *http.Request) {
	mealName := r.PostFormValue("meal-name")
	mealData, err := utils.FetchMealData("https://www.themealdb.com/api/json/v1/1/search.php?s=" + mealName)
	if err != nil {
		log.Fatal(err)
	}
	//err = utils.RenderHome(w, mealData.Meals[0])
	err = utils.RenderTemplate(w, "recipe_list", mealData.Meals[0])
	if err != nil {
		log.Fatal(err)
	}
}
