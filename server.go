package main

import (
	"fmt"
	"net/http"

	"github.com/luisabarbalho/go_cache_pragmatic_review/controller"
	router "github.com/luisabarbalho/go_cache_pragmatic_review/http"
	"github.com/luisabarbalho/go_cache_pragmatic_review/repository"
	"github.com/luisabarbalho/go_cache_pragmatic_review/service"
)

var (
	postRepository repository.PostRepository = repository.NewFirestoreRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewChiRouter()
)

func main() {
	const port string = ":8080"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(port)
}
