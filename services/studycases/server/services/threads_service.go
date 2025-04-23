package services

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yujisoyama/go_microservices/pkg/logger"
)

type ThreadsService interface {
	TestThreads(test string) (string, int, error)
}

type threadsService struct {
	log *logger.Logger
}

func NewThreadsService(log *logger.Logger) ThreadsService {
	return &threadsService{
		log: log,
	}
}

func counter(log *logger.Logger, t string) string {
	for i := 0; i < 5; i++ {
		log.Infof("%s: %d", t, i)
	}
	return t
}

// Thread 1
func withoutGoRoutine(log *logger.Logger) string {
	t := counter(log, "Normal")
	log.Info("Hello 1")
	log.Info("Hello 2")
	// will return "Normal"
	return t
}

// Thread 1
func withGoRoutine(log *logger.Logger) string {
	t := counter(log, "Normal")
	// it creates a new thread and it executes after the "Hello" logs, because the "Hello" logs are fast to execute.
	// The value of "t" is going to be "Normal" because the go routine executes after the return statement.
	// Thread 2
	go func() {
		t = counter(log, "Goroutine")
	}()

	log.Info("Hello 1")
	log.Info("Hello 2")
	// will return "Normal" because the Thread 2 executes after the return of the Thread 1.
	return t
}

// Thread 1
func withGoRoutineAndDelay(log *logger.Logger) string {
	t := counter(log, "Normal")

	// Thread 2
	go func() {
		t = counter(log, "Goroutine")
	}()

	log.Info("Hello 1")
	log.Info("Hello 2")
	time.Sleep(1 * time.Second)
	// will return "Goroutine" because the Thread 2 finishes in the 1 second delay before the return of Thread 1.
	return t
}

// Thread 1
func simpleChannelTest(log *logger.Logger) string {
	result := "Channel test"

	// Create a channel to create a comunication between Thread 1 and Thread 2
	hello := make(chan string)

	// Thread 2
	go func() {
		hello <- "Hello"
	}()

	log.Info(result)
	result = <-hello
	log.Info(result)
	// will return "Hello", because the result waits the value of channel hello
	// it not necessary to add a delay
	return result
}

// Thread 1
func foreverChannelTest(log *logger.Logger) string {
	// Create a channel to create a comunication between Thread 1 and Thread 2
	forever := make(chan string)

	// Thread 2
	go func() {
		for {
		}
	}()

	log.Info("Waiting...")
	<-forever
	// will gonna wait forever because the channel waits a value to put in it.
	// não foi passado nenhum valor para o canal forever
	return "Forever test"
}

// Thread 1
func channelSelectTest(log *logger.Logger) string {
	hello := make(chan string)

	//Thread 2
	go func() {
		hello <- "Hello"
	}()

	select {
	case x := <-hello:
		log.Info(x)
	default:
		log.Info("Default")
	}

	// will log only "Default" and return "Return"
	// when occurs the select statement, the thread 2 is not finished yet, executing de default case
	return "Return"
}

// Thread 1
func readingValuesFromChannelInLoop(log *logger.Logger) string {
	queue := make(chan int)

	// Thread 2
	go func() {
		i := 0
		for {
			time.Sleep(1 * time.Second)
			queue <- i
			i++
		}
	}()

	for x := range queue {
		log.Info(x)
	}

	// não será retornado a string "Loop".
	// a Thread 2 fica parada até que seja feita a leitura do valor dele no for da Thread 1.
	// quando ocorre a leitura de queue (ao logar o valor de x), o valor de queue é esvaiado e a Thread 2 executa a próxima iteração, preenchendo o valor em queue novamente
	return "Loop"
}

func worker(log *logger.Logger, id int, msg chan int) {
	for res := range msg {
		log.Infof("Worker %d: %d", id, res)
		time.Sleep(1 * time.Second)
	}
}

// Thread 1
func channelWithWorker(log *logger.Logger) string {
	msg := make(chan int)
	
	// Thread 2
	go worker(log, 1, msg)

	for i := 0; i < 10; i++ {
		msg <- i
	}

	// é criado um canal de int e uma função worker para receber este canal criado para realizar a leitura do seu valor
	// o worker é chamado por uma go routine (criando uma thread nova) que ficará lendo o valor do canal msg sempre que algum valor é enviado para ele
	// ao finalizar as 10 iterações, o canal não recebe mais nenhum valor, retornando a string "Channel with worker"
	return "Channel with worker"
}

// Thread 1
func channelWithMultipleWorkers(log *logger.Logger) string {
	msg := make(chan int)
	
	// Thread 2
	go worker(log, 1, msg)

	// Thread 3
	go worker(log, 2, msg)

	// Thread 4
	go worker(log, 3, msg)

	// Thread 5
	go worker(log, 4, msg)

	for i := 0; i < 10; i++ {
		msg <- i
	}

	// é criado um canal de int e uma função worker para receber este canal criado para realizar a leitura do seu valor
	// são chamados vários workers através go routines (criando várias threads) que ficarão lendo o valor do canal msg sempre que algum valor é enviado para ele
	// desse modo, enquanto um worker está parado (por conta do time.Sleep), um outro worker já poderá ir lendo o valor que foi enviado para o canal através do loop
	// portanto, as 10 iterações são finalizadas bem mais rapidamente, pois existem vários workers que estão "esvaziando" o canal, permitindo o envio de valores através dele
	return "Channel with worker"
}

func (ts *threadsService) TestThreads(test string) (string, int, error) {
	ts.log.Info("---------------- Threads Test ----------------")

	switch test {
	case "1":
		// flow without goroutine
		return withoutGoRoutine(ts.log), fiber.StatusOK, nil
	case "2":
		// flow with goroutine
		return withGoRoutine(ts.log), fiber.StatusOK, nil
	case "3":
		// flow with goroutine and delay
		return withGoRoutineAndDelay(ts.log), fiber.StatusOK, nil
	case "4":
		// flow to a simple channel test
		return simpleChannelTest(ts.log), fiber.StatusOK, nil
	case "5":
		// flow to a forever channel test
		return foreverChannelTest(ts.log), fiber.StatusOK, nil
	case "6":
		return channelSelectTest(ts.log), fiber.StatusOK, nil
	case "7":
		return readingValuesFromChannelInLoop(ts.log), fiber.StatusOK, nil
	case "8":
		return channelWithWorker(ts.log), fiber.StatusOK, nil
	case "9":
		return channelWithMultipleWorkers(ts.log), fiber.StatusOK, nil
	default:
		return "", fiber.StatusOK, nil
	}
}
