package container

import (
	"fmt"
	"io"
	"os"
	"product-master/internal/domain/productd"
	"product-master/internal/helper"
	"product-master/internal/infrastructure/postgre"
	"product-master/internal/usecase/productu"
	"product-master/internal/utils"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

type (
	Container struct {
		Apps       Apps
		Logger     Logger
		ProductUsc productu.ProductUsc
	}

	Logger struct {
		Log     zerolog.Logger
		Path    string `env:"log_path"`
		LogFile string `env:"log_file"`
	}

	Apps struct {
		Name           string `env:"apps_appName"`
		Host           string `env:"apps_host"`
		Version        string `env:"apps_version"`
		SwaggerAddress string `env:"apps_swagger_address"`
		HttpPort       int    `env:"apps_httpport"`
		SecretJwt      string `env:"apps_secretJwt"`
		CtxTimeout     int    `env:"apps_timeout"`
	}
)

func NewContainer(envName ...string) *Container {
	if envName == nil {
		envName = append(envName, ".env")
	}
	env := envName[0]

	err := godotenv.Load(fmt.Sprintf("%s/%s", helper.ProjectRootPath, env))
	if err != nil {
		panic(err)
	}

	var appsConf = Apps{
		Name:           utils.EnvString("apps_appName"),
		Host:           utils.EnvString("apps_host"),
		Version:        utils.EnvString("apps_version"),
		SwaggerAddress: utils.EnvString("apps_swagger_address"),
		HttpPort:       utils.EnvInt("apps_httpport"),
		SecretJwt:      utils.EnvString("apps_secretJwt"),
		CtxTimeout:     utils.EnvInt("apps_timeout"),
	}

	utils.InitJWT(utils.EnvString("apps_secretJwt"))

	postgre, err := postgre.Init()
	if err != nil {
		panic(err)
	}

	sqlTrx := utils.NewSQLTransaction(postgre)
	productRepo := productd.NewProductDomain(postgre, sqlTrx)

	productUsc := productu.NewProductUsecase(productRepo)

	cont := Container{
		Apps:       appsConf,
		ProductUsc: productUsc,
		Logger:     LoggerInit(),
	}

	return &cont
}

func LoggerInit() Logger {
	var loggerConf = Logger{
		Path:    utils.EnvString("log_path"),
		LogFile: utils.EnvString("log_file"),
	}
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	helper.Logger(helper.LoggerLevelInfo, "Succeed when read loggerConf", nil)

	var stdout io.Writer = os.Stdout
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	if loggerConf.LogFile == "ON" {
		path := fmt.Sprintf("%s%s", helper.ProjectRootPath, loggerConf.Path)
		file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0664)
		if err != nil {
			helper.Logger(helper.LoggerLevelError, "", fmt.Errorf("error when setting loggerConf : %s", err.Error()))
		}
		// Create a multi writer with both the console and file writers
		stdout = zerolog.MultiLevelWriter(os.Stdout, file)

	}

	loggerConf.Log = zerolog.New(stdout).With().Caller().Timestamp().Logger()
	helper.Logger(helper.LoggerLevelInfo, "Succeed read loggerConf", nil)
	return loggerConf
}
