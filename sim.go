package main

import "fmt"
import "bufio"
import "os"
import "time"
import "github.com/sirupsen/logrus" //for added flavor for the loppuhuipennus

//dating sim text adventure in old leisure suit larry spirit

func main() {

	fmt.Println("\n-= DeittiSim 2002 =- (enter jatkaaksesi)")
	fmt.Scanln()
	fmt.Println("Tervetuloa, sinä kaunis uros tai lihaksikas nainen!")
	fmt.Scanln()
	fmt.Println("Tämä on deittisimulaattori, jossa taitojasi naisten viettelijänä tai viettelittäjärenä koetellaan.")
	fmt.Scanln()
	fmt.Println("Pääset pelin haarakohdissa eteenpäin kirjoittamalla käskyn annetuista vaihtehdoista.")
	fmt.Scanln()
	reader := bufio.NewReader(os.Stdin) //standard input

	objGood := 0 //pelin eteneminen
	objBad := 0

	fmt.Println("********************")
	fmt.Scanln()
	fmt.Println("On perjantai-ilta ja saavut baariin vain yksi asia mielessä:")
	fmt.Scanln()
	fmt.Println("arvonlisäverot.")
	fmt.Scanln()
	fmt.Println("Siksi kaipaat seuraa ensi yöksi.")
	fmt.Scanln()
	fmt.Println("Baaritiskillä istuu kaunein nainen jonka olet koskaan nähnyt, pakarat hulmuten ja tukka kiiltäen.")
	fmt.Scanln()
	fmt.Println("Voit lähestyä naista joko kehulla, tai tökeröllä iskurepliikillä.")
	fmt.Println("Vaihtoehdot: 'kehu' 'repliikki'")

	for objGood == 0 && objBad == 0 {
		byteOption, _, _ := reader.ReadLine()
		option := string(byteOption)
		if option == "kehu" {
			objGood++
			fmt.Println("\n********************")
			fmt.Println("\nKehut naista kauniimmaksi kuin koivua kuusimetsässä.")
			fmt.Scanln()
		} else if option == "repliikki" {
			objBad++
			fmt.Println("\n********************")
			fmt.Println("\nKerrot naiselle että onnistuit hukkaamaan puhelinnumerosi, ja kysyt saisitko hänen.")
			fmt.Scanln()
		} else {
			fmt.Println("Se ei ollut vaihtoehto, päärynä. Yritä uudelleen.")
		}
	}

	for objGood == 1 && objBad == 0 {
		fmt.Println("Nainen hymyilee sinulle, mutta älä vielä nuolaise!")
		fmt.Scanln()
		fmt.Println("Voit jatkaa joko leveilemällä pankkitililläsi, tai kokeilla huonoa iskurepliikkiä.")
		fmt.Println("Vaihtoehdot: 'leveile' 'repliikki'")
		byteOption, _, _ := reader.ReadLine()
		option := string(byteOption)
		if option == "leveile" {
			objGood++
			fmt.Println("\n********************")
			fmt.Println("\nKerrot naiselle kuinka pankkitililläsi on enemmän nollia kuin hänellä on ollut huonoja miehiä elämässään.")
			fmt.Scanln()
		} else if option == "repliikki" {
			objBad++
			fmt.Println("\n********************")
			fmt.Println("\nKysyt naiselta, että sattuiko se kun hän putosi taivaasta, sillä hän on enkelimäisen kaunis.")
			fmt.Scanln()
		} else {
			fmt.Println("Se ei ollut vaihtoehto, päärynä. Yritä uudelleen.")
		}
	}

	for objGood == 0 && objBad == 1 {
		fmt.Println("Nainen on kaukana vaikuttuneesta. Nyt varovasti!")
		fmt.Scanln()
		fmt.Println("Voit joko tarjota naiselle drinkin, tai kertoa hänelle valtavasta miehisyydestästi.")
		fmt.Println("Vaihtoehdot: 'drinkki' 'ole mies'")
		byteOption, _, _ := reader.ReadLine()
		option := string(byteOption)
		if option == "drinkki" {
			objGood++
			fmt.Println("\n********************")
			fmt.Println("\nTarjoat naiselle juoman, Seksiä Rannalla.")
			fmt.Scanln()
		} else if option == "ole mies" {
			objBad++
			fmt.Println("\n********************")
			fmt.Println("\nKerrot naiselle valtavasta halkojastasi, jolla teit viime vuonna vaikutuksen naapurin Irmeliin.")
			fmt.Scanln()
		} else {
			fmt.Println("Se ei ollut vaihtoehto, päärynä. Yritä uudelleen.")
		}
	}

	for objGood == 1 && objBad == 1 {
		fmt.Println("Tämä ilta voi vielä johtaa johonkin. Tsemppiä!")
		fmt.Scanln()
		fmt.Println("Voit joko tehdä taikatempun tai pyytää naista kahville, kunnon keskustelun ääreen.")
		fmt.Println("Vaihtoehdot: 'taikatemppu' 'kahvi'")
		byteOption, _, _ := reader.ReadLine()
		option := string(byteOption)
		if option == "kahvi" {
			objGood++
			fmt.Println("\n********************")
			fmt.Println("\nKysyt jos naista kiinnostaisi oikeasti tutustua, ja ehdotat kahvia viereissä myöhäisillan kuppilassa")
			fmt.Scanln()
		} else if option == "taikatemppu" {
			objBad++
			fmt.Println("\n********************")
			fmt.Println("\nVedät teatraalisesti jäniksen sepaluksestasi.")
			fmt.Scanln()
			fmt.Println("Nainen on yhtäaikaa vaikuttuneen ja häiriintyneen näköinen.")
			fmt.Scanln()
		} else {
			fmt.Println("Se ei ollut vaihtoehto, päärynä. Yritä uudelleen.")
		}
	}

	for objBad == 2 {
		fmt.Println("********************")
		fmt.Println("\nNo nyt ei mennyt ihan nappiin. Parempi onni ensi kerralla!")
		os.Exit(3)
	}

	for objGood == 2 {
		timertime := 1 //viestiviiveen kertasäätäminen

		fmt.Println("Nainen on vaikuttunut, ja suostuu lähtemään matkaasi. Voitit!")
		time.Sleep(time.Duration(timertime)*time.Second)
		fmt.Println("\n      ,(())),")
		time.Sleep(time.Duration(timertime)*time.Second)
		fmt.Println("     '(('''))'")
		time.Sleep(time.Duration(timertime)*time.Second)
		fmt.Println("     '(|*_*|)'")
		time.Sleep(time.Duration(timertime)*time.Second)
		fmt.Println("       : = :")
		time.Sleep(time.Duration(timertime)*time.Second)
		fmt.Println("       _) (_")
		time.Sleep(time.Duration(timertime)*time.Second)
		logrus.Error("CENSORED")
		logrus.Error("CENSORED")
		logrus.Error("( ͡° ͜ʖ ͡°)")
		os.Exit(3)
	}
}
