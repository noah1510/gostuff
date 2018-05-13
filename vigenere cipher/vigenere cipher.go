package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

loop:
	for {
		println("\n\n------------------------------\n")
		fmt.Print("0 - close program\n1 - Encrypt a message\n2 - Decrypt a message\nYour Choice: ")

		ins, err := readString()
		if err == nil {
			choice, err := strconv.Atoi(ins)
			if err == nil {
				switch choice {

				//close the program
				case 0:
					println("\n------------------------------\n")
					println("Closing program")
					println("\n------------------------------\n")

					break loop

				//Encrypt a message
				case 1:
					var key string
					var message string

					for {

						println("\n------------------------------\n")
						println("Enter your key:")
						ins, err := readString()
						if err == nil {
							key = ins
							break
						} else {
							println("Some error with your input" + err.Error())
						}
					}

					for {
						println("\n------------------------------\n")
						println("Enter your message:")
						ins, err := readString()
						if err == nil {
							message = ins
							break
						} else {
							println("Some error with your input" + err.Error() + "\n")
						}
					}

					println("\n------------------------------\n")

					messagee, err := encrypt(&message, &key)

					if err == nil {
						println("Your message encrypted is: " + *messagee)
					} else {
						println("Something went wrong:" + err.Error())
					}

				//Decrypt a message
				case 2:
					var key string
					var message string

					for {

						println("\n------------------------------\n")
						println("Enter your key:")
						ins, err := readString()
						if err == nil {
							key = ins
							break
						} else {
							println("Some error with your input" + err.Error())
						}
					}

					for {
						println("\n------------------------------\n")
						println("Enter your message:")
						ins, err := readString()
						if err == nil {
							message = ins
							break
						} else {
							println("Some error with your input" + err.Error() + "\n")
						}
					}

					println("\n------------------------------\n")

					messaged, err := decrypt(&message, &key)

					if err == nil {
						println("Your message decrypted is: " + *messaged)
					} else {
						println("Something went wrong:" + err.Error())
					}

				//None of the above
				default:
					println("\n------------------------------\n")
					println("not a valid command!")
				}
			} else {
				println("input is not a number" + err.Error())
			}
		} else {
			println("Some error with your input" + err.Error())
		}
	}

}

func readString() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err == nil {
		text = strings.TrimSuffix(text, "\r\n")
		return text, nil
	}

	return "", err

}

func encrypt(input *string, k *string) (*string, error) {
	key := []byte(*k)
	output := make([]byte, len(*input))
	for i := range *input {
		output[i] = byte(uint(((*input)[i] + key[i%len(key)])) % 256)
	}
	sOutput := string(output)
	return &sOutput, nil

}

func decrypt(input *string, k *string) (*string, error) {
	key := []byte(*k)
	output := make([]byte, len(*input))
	for i := range *input {
		output[i] = byte(uint(((*input)[i] - key[i%len(key)])) % 256)
	}
	sOutput := string(output)
	return &sOutput, nil

}

func println(text string) (int, error) {
	return fmt.Println(text)
}
