package main

import (
	"fmt"
    "log"
    "context"
    "flag"
    "net"
	"go-web/model"
	"go-web/controller"
	//"go-web/util/token"
	"github.com/gin-gonic/gin"
	//"gorm.io/driver/postgres"
	//"gorm.io/gorm"
    "net/http"
    "github.com/gin-contrib/cors"

	"google.golang.org/grpc"
	pb "go-web/util/grpc/helloworld/helloworld"

)

var (
	grpc_server_port = flag.Int("grpc_server_port", 50051, "The grpc server port")
	grpc_client_port = flag.Int("grpc_client_port", 50052, "The groc client port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName() + ", This is Gin"}, nil
}

/*
func SetupDB() *gorm.DB {
}
*/
func SetupRestRouter(mh *model.ModelHandler) *gin.Engine {
	r := gin.Default()

    //r.Use(CORSMiddleware())
    r.Use(cors.New(cors.Config{
     AllowOrigins:     []string{"*"},
     AllowMethods:     []string{"POST", "GET","DELETE","PUT"},
     AllowHeaders:     []string{"Origin","Authorization"},
     ExposeHeaders:    []string{"Content-Length"},
     AllowCredentials: true,
    }))

    protected := r.Group("/admin")
    protected.Use(controller.JwtAuthMiddleware());
    protected.GET("/user", mh.FetchUser);

	r.POST("/user", mh.CreateUser)
    r.GET("/user", mh.FetchUser)
    r.DELETE("/user", mh.DeleteUser)
    r.PUT("/user", mh.UpdateUser)

	r.POST("/login", mh.Login)

	r.POST("/post", mh.CreatePost)
	r.GET("/post", mh.FetchPost)

    r.POST("/file", mh.SaveFileHandler)
	r.POST("/upload", mh.UploadFiles)

    //Lets use HTML
    r.Static("/assets","./assets")
    //Lets Start to Write new Templates Through html/template module
    r.LoadHTMLGlob("templates/**/*")
    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "main", gin.H{
            "title": "Main website",
        })

   })
   return r
}

func main() {

/*
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Seoul"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect PG Database")
	}
	fmt.Println("PG db Connection Successful")
    modelHandler := model.ModelHandler{Db:db}
*/
    err, modelHandler := model.ModelInit("env.json")
	if err != nil {
		panic("Failed to Init PG Database")
	}

    /* Before Creation Router of gin, start grpc Server  in go routine */

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpc_server_port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

    go func() {
        if err := s.Serve(lis); err != nil {
            log.Fatalf("could not start grpc server: %v", err)
        }
    }()

	/* Gin Basic Usage ping/pong */
    r := SetupRestRouter(&modelHandler)

	r.Run(":8000") // listen and serve on 0.0.0.0:8000 (for windows "localhost:8080")
    /* After Web Framework, Start grpc Server*/
    //Check this to implement both server on this https://github.com/gin-gonic/gin/issues/2594


}
