package _func

import "github.com/DenisKDO/TSIS1/pkg/structs"

// making team
func PrepareTeam() []structs.Player {
	var players []structs.Player

	var player structs.Player

	player.UniformNumber = "1"
	player.Country = "Japan"
	player.Position = "OP"
	player.FirstName = "Yuji"
	player.SecondName = "Nishida"
	player.Height = "186cm"
	player.Weight = "80kg"
	players = append(players, player)

	player.UniformNumber = "12"
	player.Country = "Japan"
	player.Position = "OH"
	player.FirstName = "Ran"
	player.SecondName = "Takahashi"
	player.Height = "188cm"
	player.Weight = "72kg"
	players = append(players, player)

	player.UniformNumber = "14"
	player.Country = "Japan"
	player.Position = "OH"
	player.FirstName = "Yuki"
	player.SecondName = "Ishikawa"
	player.Height = "191cm"
	player.Weight = "84kg"
	players = append(players, player)

	player.UniformNumber = "4"
	player.Country = "Japan"
	player.Position = "OP"
	player.FirstName = "Kento"
	player.SecondName = "Miyaura"
	player.Height = "189cm"
	player.Weight = "78kg"
	players = append(players, player)

	player.UniformNumber = "5"
	player.Country = "Japan"
	player.Position = "OH"
	player.FirstName = "Tatsunori"
	player.SecondName = "Otsuka"
	player.Height = "194cm"
	player.Weight = "80kg"
	players = append(players, player)

	player.UniformNumber = "11"
	player.Country = "Japan"
	player.Position = "OH"
	player.FirstName = "Shoma"
	player.SecondName = "Tomita"
	player.Height = "190cm"
	player.Weight = "75kg"
	players = append(players, player)

	player.UniformNumber = "30"
	player.Country = "Japan"
	player.Position = "OH"
	player.FirstName = "Masato"
	player.SecondName = "Kai"
	player.Height = "184cm"
	player.Weight = "72kg"
	players = append(players, player)

	player.UniformNumber = "2"
	player.Country = "Japan"
	player.Position = "MB"
	player.FirstName = "Onodera"
	player.SecondName = "Taishi"
	player.Height = "201cm"
	player.Weight = "98kg"
	players = append(players, player)

	player.UniformNumber = "6"
	player.Country = "Japan"
	player.Position = "MB"
	player.FirstName = "Akihiro"
	player.SecondName = "Yamauchi"
	player.Height = "204cm"
	player.Weight = "80kg"
	players = append(players, player)

	player.UniformNumber = "10"
	player.Country = "Japan"
	player.Position = "MB"
	player.FirstName = "Kentaro"
	player.SecondName = "Takahashi"
	player.Height = "201cm"
	player.Weight = "103kg"
	players = append(players, player)

	player.UniformNumber = "12"
	player.Country = "Japan"
	player.Position = "S"
	player.FirstName = "Masahiro"
	player.SecondName = "Sekita"
	player.Height = "175cm"
	player.Weight = "72kg"
	players = append(players, player)

	player.UniformNumber = "29"
	player.Country = "Japan"
	player.Position = "S"
	player.FirstName = "Ryu"
	player.SecondName = "Yamamoto"
	player.Height = "180cm"
	player.Weight = "65kg"
	players = append(players, player)

	player.UniformNumber = "13"
	player.Country = "Japan"
	player.Position = "L"
	player.FirstName = "Tomohiro"
	player.SecondName = "Ogawa"
	player.Height = "176cm"
	player.Weight = "66kg"
	players = append(players, player)

	player.UniformNumber = "20"
	player.Country = "Japan"
	player.Position = "L"
	player.FirstName = "Tomohiro"
	player.SecondName = "Yamamoto"
	player.Height = "171cm"
	player.Weight = "69"
	players = append(players, player)
	return players
}
