package client

//go:generate sh -c "rm -rf mocks && mkdir -- mocks"
//go:generate minimock -i PasswordManager -o ./mocks/ -s "_minimock.go"
