module example.com/my/parser

replace example.com/my/dataLoader => /home/pam/csvParser/src/dataLoader

replace example.com/my/dataParser => /home/pam/csvParser/src/dataParser

replace example.com/my/formatConverter => /home/pam/csvParser/src/formatConverter

replace example.com/my/dataOutputer => /home/pam/csvParser/src/dataOutputer

replace example.com/my/utils => /home/pam/csvParser/src/utils

go 1.13

require (
	example.com/my/dataLoader v0.0.0
	example.com/my/dataOutputer v0.0.0
	example.com/my/dataParser v0.0.0
	example.com/my/formatConverter v0.0.0
	example.com/my/utils v0.0.0
)
