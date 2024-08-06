package repository

//go:generate sh -c "rm -rf mocks && mkdir -- mocks"
//go:generate minimock -i UserRepository -o ./mocks/ -s "_minimock.go"
