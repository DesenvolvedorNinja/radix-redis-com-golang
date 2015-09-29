package main

import (
	"fmt"
	"github.com/mediocregopher/radix.v2/redis"
	"os"
)

func main() {

	client, err := redis.Dial("tcp", "pub-redis-13919.us-east-1-2.5.ec2.garantiadata.com:13919")
	if err != nil {
		fmt.Printf("Erro na conexão com o Redis\n")
		os.Exit(1)
	}

	err = client.Cmd("AUTH", "salamandra").Err
	if err != nil {
		fmt.Printf("Erro ao autenticar\n" + err.Error() + "\n")
		client.Close()
		os.Exit(1)
	}

	err = client.Cmd("SET", "admin_sistema", "Desenvolvedor Ninja").Err
	if err != nil {
		fmt.Printf("Erro ao definir chave\n" + err.Error() + "\n")
		client.Close()
		os.Exit(1)
	}

	chave, err := client.Cmd("GET", "admin_sistema").Str()
	if err != nil {
		fmt.Printf("Erro ao pesquisar chave\n")
		client.Close()
		os.Exit(1)
	}

	fmt.Printf(chave + "\n")

	err = client.Cmd("EXPIRE", "admin_sistema", "600").Err
	if err != nil {
		fmt.Printf("Erro ao definir tempo de expiração\n")
		client.Close()
		os.Exit(1)
	}

	tempo, err := client.Cmd("TTL", "admin_sistema").Int()
	if err != nil {
		fmt.Printf("Erro ao obter tempo de expiração\n")
		client.Close()
		os.Exit(1)
	}

	fmt.Printf("%d\n", tempo)

	err = client.Cmd("PERSIST", "admin_sistema").Err
	if err != nil {
		fmt.Printf("Erro ao retirar tempo de expiração\n")
		client.Close()
		os.Exit(1)
	}

	err = client.Cmd("DEL", "admin_sistema").Err
	if err != nil {
		fmt.Printf("Erro ao apagar chave\n")
		client.Close()
		os.Exit(1)
	}

	client.Close()
}
