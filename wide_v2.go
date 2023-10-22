package main

import (
	"fmt"

	survey "github.com/AlecAivazis/survey/v2"
)

var TransFuelQS = []*survey.Question{
	{
		Name: "Transmission",
		Prompt: &survey.Select{
			Message: "Какую коробку передач вы предпочитаете?",
			Options: []string{"Автомат", "Механика"},
		},
	},
	{
		Name: "Fuel",
		Prompt: &survey.Select{
			Message: "Тип топлива?",
			Options: []string{"Бензин", "Газ"},
		},
	},
}

type TrFl struct {
	Transmission string
	Fuel         string
}

func DDD() {
	prompCarType := survey.Select{
		Message: "Легковая или грузовая машина?",
		Options: []string{"Легковая", "Грузовая"},
		Default: "Легковая",
	}
	typeAns := ""
	err := survey.AskOne(&prompCarType, &typeAns)
	if err != nil {
		fmt.Println(err)
	}
	switch typeAns {
	case "Легковая":
		passPrice := ""
		survey.AskOne(&survey.Select{
			Message: "Какой у вас бюджет?",
			Options: []string{"<500", "500-1500", "1500-3000", ">3000"},
		}, &passPrice)
		switch passPrice {
		case "<500":
			fmt.Println("Машины, которые подходят вам: Лада")

		case "500-1500":
			TrFlAns := TrFl{}
			survey.Ask(TransFuelQS, &TrFlAns)
			switch {
			case TrFlAns.Transmission == "Автомат":
				switch TrFlAns.Fuel {
				case "Бензин":
					seatAns := ""
					survey.AskOne(&survey.Select{
						Message: "Количество мест?",
						Options: []string{"2", "4"},
					}, &seatAns)
					switch seatAns {
					case "2":
						fmt.Println("Машины, которые подходят вам: Хендай")
					case "4":
						roadAns := ""
						survey.AskOne(&survey.Select{
							Message: "Назначение?",
							Options: []string{"Город", "Бездорожье"},
						}, &roadAns)
						switch roadAns {
						case "Город":
							fmt.Println("Машины, которые подходят вам: Хендай, Лада")
						case "Бездорожье":
							fmt.Println("Машины, которые подходят вам: Хендай")
						}
					}
				case "Газ":
					fmt.Println("Машины, которые подходят вам: Хендай")
				}
			case TrFlAns.Transmission == "Механика":
				switch TrFlAns.Fuel {
				case "Бензин":
					seatAns := ""
					survey.AskOne(&survey.Select{
						Message: "Количество мест?",
						Options: []string{"2", "4"},
					}, &seatAns)
					switch seatAns {
					case "2":
						fmt.Println("Машины, которые подходят вам: Смарт")
					case "4":
						roadAns := ""
						survey.AskOne(&survey.Select{
							Message: "Назначение?",
							Options: []string{"Город", "Бездорожье"},
						}, &roadAns)
						switch roadAns {
						case "Город":
							fmt.Println("Машины, которые подходят вам: Лада, Хонда")
						case "Бездорожье":
							fmt.Println("Машины, которые подходят вам: Ситроен")
						}
					}
				case "Газ":
					fmt.Println("Машины, которые подходят вам: Хендай")
				}
			}

		case "1500-3000":
			TransAns := ""
			survey.AskOne(&survey.Select{
				Message: "Какую коробку передач вы предпочитаете?",
				Options: []string{"Автомат", "Механика"},
			}, &TransAns)
			switch TransAns {
			case "Автомат":
				FuelAns := ""
				survey.AskOne(&survey.Select{
					Message: "Тип топлива?",
					Options: []string{"Бензин", "Дизель"},
				}, &FuelAns)

				switch FuelAns {
				case "Бензин":
					seatAns := ""
					survey.AskOne(&survey.Select{
						Message: "Количество мест?",
						Options: []string{"2", "4"},
					}, &seatAns)
					switch seatAns {
					case "2":
						fmt.Println("Машины, которые подходят вам: Смарт")
					case "4":
						roadAns := ""
						survey.AskOne(&survey.Select{
							Message: "Назначение?",
							Options: []string{"Город", "Бездорожье"},
						}, &roadAns)
						switch roadAns {
						case "Город":
							fmt.Println("Машины, которые подходят вам: Ситроен")
						case "Бездорожье":
							fmt.Println("Машины, которые подходят вам: Хендай, Хонда")
						}
					}
				case "Дизель":
					seatAns := ""
					survey.AskOne(&survey.Select{
						Message: "Количество мест?",
						Options: []string{"2", "4"},
					}, &seatAns)
					switch seatAns {
					case "2":
						fmt.Println("Машины, которые подходят вам: Ситроен")
					case "4":
						roadAns := ""
						survey.AskOne(&survey.Select{
							Message: "Назначение?",
							Options: []string{"Город", "Бездорожье"},
						}, &roadAns)
						switch roadAns {
						case "Город":
							fmt.Println("Машины, которые подходят вам: Ситроен, Хонда")
						case "Бездорожье":
							fmt.Println("Машины, которые подходят вам: Хонда")
						}
					}
				}
			case "Механика":
				FuelAns := ""
				survey.AskOne(&survey.Select{
					Message: "Тип топлива?",
					Options: []string{"Бензин", "Электричество"},
				}, &FuelAns)

				switch FuelAns {
				case "Бензин":
					seatAns := ""
					survey.AskOne(&survey.Select{
						Message: "Количество мест?",
						Options: []string{"2", "4"},
					}, &seatAns)
					switch seatAns {
					case "2":
						fmt.Println("Машины, которые подходят вам: Смарт")
					case "4":
						roadAns := ""
						survey.AskOne(&survey.Select{
							Message: "Назначение?",
							Options: []string{"Город", "Бездорожье"},
						}, &roadAns)
						switch roadAns {
						case "Город":
							fmt.Println("Машины, которые подходят вам: Мерседес, Хендай")
						case "Бездорожье":
							fmt.Println("Машины, которые подходят вам: Ситроен, БМВ")
						}
					}
				case "Электричество":
					seatAns := ""
					survey.AskOne(&survey.Select{
						Message: "Количество мест?",
						Options: []string{"2", "4"},
					}, &seatAns)
					switch seatAns {
					case "2":
						fmt.Println("Машины, которые подходят вам: Смарт")
					case "4":
						fmt.Println("Машины, которые подходят вам: Тесла")
					}
				}
			}
		case ">3000":
			FuelAns := ""
			survey.AskOne(&survey.Select{
				Message: "Тип топлива?",
				Options: []string{"Бензин", "Электричество"},
			}, &FuelAns)
			switch FuelAns {
			case "Бензин":
				seatAns := ""
				survey.AskOne(&survey.Select{
					Message: "Количество мест?",
					Options: []string{"2", "4"},
				}, &seatAns)
				switch seatAns {
				case "2":
					fmt.Println("Машины, которые подходят вам: БМВ")
				case "4":
					roadAns := ""
					survey.AskOne(&survey.Select{
						Message: "Назначение?",
						Options: []string{"Город", "Бездорожье"},
					}, &roadAns)
					switch roadAns {
					case "Город":
						fmt.Println("Машины, которые подходят вам: Мерседес")
					case "Бездорожье":
						fmt.Println("Машины, которые подходят вам: Мерседес, БМВ")
					}
				}
			case "Электричество":
				fmt.Println("Машины, которые подходят вам: Тесла")
			}
		}
	case "Грузовая":
		cargoPrice := ""
		survey.AskOne(&survey.Select{
			Message: "Какой у вас бюджет?",
			Options: []string{"1000", ">5000"},
		}, &cargoPrice)
		switch cargoPrice {
		case "1000":
			fmt.Println("Машины, которые подходят вам: Лада, Мерседес")
		case ">5000":
			fmt.Println("Машины, которые подходят вам: БМВ, Вольво")
		}
	}
}
