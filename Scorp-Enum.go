package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func subdomain_finding() {
	fmt.Println("  ")

	fmt.Println("\033[0m")
	fmt.Println("")
	var domain string
	fmt.Println("\033[31m")
	fmt.Print("Enter the domain----------> ")
	fmt.Print("\033[0m")
	fmt.Scan(&domain)
	fmt.Println()
	fmt.Println("Please wait,it might be take a few minutes...")

	conn, err := http.Get("https://" + domain)
	if err != nil {
		conn_2, errr := http.Get("http://" + domain)
		if errr != nil {
			fmt.Println(errr)
		} else {

			if conn_2.StatusCode == 200 {
				fmt.Println("\033[36m")
				fmt.Println("We can access to this site on port 80")
				fmt.Println("\033[0m")
				var subdomain_wordlist string
				fmt.Println("\033[31m")
				fmt.Print("Enter the subdomain wordlists path----------> ")
				fmt.Print("\033[0m")
				fmt.Scan(&subdomain_wordlist)

				okunan, errs := os.Open(subdomain_wordlist)
				if errs != nil {
					fmt.Println(errs)
				}
				fileScanner := bufio.NewScanner(okunan)
				fileScanner.Split(bufio.ScanLines)

				for fileScanner.Scan() {
					url_1 := "https://" + fileScanner.Text() + "." + domain
					response, errs := http.Get(url_1)
					if errs != nil {
						url_2 := "http://" + domain + "." + fileScanner.Text()
						response_2, errso := http.Get(url_2)
						if errso != nil {
							continue
						} else {
							if response_2.StatusCode == 200 {
								fmt.Println("\033[33m")
								fmt.Println("[+] Found--->" + url_2)
								fmt.Println("\033[0m")

							} else {
								continue
							}

						}

					} else {
						if response.StatusCode == 200 {
							fmt.Println("\033[33m")
							fmt.Println("[+] Found--->" + url_1)
							fmt.Println("\033[0m")
						} else {
							continue
						}
					}

				}

			}

		}

	} else {
		if conn.StatusCode == 200 {
			fmt.Println("\033[36m")
			fmt.Println("We can access to this site on port 443")
			fmt.Println("\033[0m")
			var subdomain_wordlist string
			fmt.Println("\033[31m")
			fmt.Print("Enter the subdomain wordlist path----------> ")
			fmt.Print("\033[0m")
			fmt.Scan(&subdomain_wordlist)

			okunan, errs := os.Open(subdomain_wordlist)
			if errs != nil {
				fmt.Println(errs)
			}
			fileScannerSD := bufio.NewScanner(okunan)
			fileScannerSD.Split(bufio.ScanLines)

			for fileScannerSD.Scan() {
				url_1 := "https://" + domain + "." + fileScannerSD.Text()
				response, errs := http.Get(url_1)
				if errs != nil {
					url_2 := "http://" + domain + "." + fileScannerSD.Text()
					response_2, errso := http.Get(url_2)
					if errso != nil {
						continue
					} else {
						if response_2.StatusCode == 200 {
							fmt.Println("\033[33m")
							fmt.Println("[+] Found--->" + url_2)
							fmt.Println("\033[0m")

						} else {
							continue
						}
					}

				} else {
					if response.StatusCode == 200 {
						fmt.Println("\033[33m")
						fmt.Println("[+] Found--->" + url_1)
						fmt.Println("\033[0m")

					} else {
						continue
					}
				}

			}

		}
	}
}

func http_https_checking() {
	var file_path string
	fmt.Println("")
	fmt.Println("\033[33m")
	fmt.Print("Enter the wordlist path------------>" + "\033[0m")
	fmt.Scan(&file_path)

	readed_file, err := os.Open(file_path)
	if err != nil {
		fmt.Println("We got an error when opening the file!")

	}
	fileScannerHC := bufio.NewScanner(readed_file)
	fileScannerHC.Split(bufio.ScanLines)

	for fileScannerHC.Scan() {
		res, err := http.Get("https://" + fileScannerHC.Text())
		if err != nil {
			res_2, errr := http.Get("http://" + fileScannerHC.Text())
			if errr != nil {
				fmt.Println("Please check your domain!")
				break

			} else {

				if res_2.StatusCode == 200 {
					fmt.Println("\033[31m")
					fmt.Println("Site looking http---------------->" + "http://" + fileScannerHC.Text())
					fmt.Println("\033[0m")
				}
			}

		} else {
			if res.StatusCode == 200 {
				fmt.Println("\033[33m")
				fmt.Println("Site looking https--------------->" + "https://" + fileScannerHC.Text())
				fmt.Println("\033[0m")

			}
		}
	}
	readed_file.Close()
}

func main() {
	fmt.Println("\033[32m")
	fmt.Println("-----------------WELCOME TO MY SCRIPT-----------------")
	fmt.Println("\033[0m")
	fmt.Println("1-Discovery Subdomains\n2-Http/Https Checking")
	var option_number int
	fmt.Println("\033[0m")
	fmt.Println("")
	fmt.Print("What is your option:")
	fmt.Scan(&option_number)
	switch option_number {

	case 1:
		subdomain_finding()
	case 2:
		http_https_checking()
	default:
		fmt.Println("!!!Wrong number!!!")
		break
	}

}
