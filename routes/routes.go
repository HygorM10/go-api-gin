package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hygorm10/go-api-gin/controllers"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosOsAlunos)
	r.GET("/alunos/:id", controllers.BuscaAluno)
	r.POST("/alunos", controllers.CriaAluno)
	r.GET("/:nome", controllers.Saudacao)
	r.DELETE("/alunos/:id", controllers.DeletarAluno)
	r.PUT("/alunos/:id", controllers.AtualizarAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.Run()
}
